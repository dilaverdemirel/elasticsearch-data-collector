package scheduler

import (
	"context"
	"encoding/json"
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
	"eslasticsearchdatacollector/service"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/jmoiron/sqlx"
)

func Sync(index_id string) {
	var index = service.GetIndexById(index_id)
	var sync_log = model.SyncLog{
		IndexId:   index_id,
		StartDate: time.Now(),
		Status:    model.SyncLogStatusStarted,
	}
	service.CreateSyncLog(&sync_log)

	data_source := dao.ConnectDatabaseWithDefinedDatasource(index.DataSourceId)
	var last_sync_execution_date = time.Now()
	record_count, err := migrate_data_to_rlasticsearch(&data_source, &index)

	if err != nil {
		service.UpdateSyncLogAsFailed(sync_log.ID, err.Error())
	} else {
		service.UpdateIndexLastSyncDate(index_id, last_sync_execution_date)
		service.UpdateSyncLogAsCompleted(sync_log.ID, record_count)
	}
}

func migrate_data_to_rlasticsearch(data_source *sqlx.DB, index *model.Index) (int32, error) {
	// fetch all places from the db
	rows, err := data_source.Queryx(index.SqlQuery)
	if err != nil {
		return -1, err
	}
	colNames, err := rows.Columns()
	if err != nil {
		return -1, err
	}

	// iterate over each row
	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{})

	var idField string = index.DocumentField

	var record_count int32 = 0

	row_list := dao.ScanRowsWithoutRowLimit(*rows)

	for l := row_list.Front(); l != nil; l = l.Next() {
		row := make(map[string]interface{})
		var id_value string = ""
		for i := 0; i < len(colNames); i++ {
			col_name := colNames[i]
			value := l.Value.(map[string]interface{})[col_name]
			row[col_name] = value

			if idField == col_name {
				field_data_type := reflect.TypeOf(value).String()

				switch field_data_type {
				case "int", "int8", "int16", "int32", "int64":
					id_value = fmt.Sprintf("%s", value)
				case "float32", "float64":
					id_value = fmt.Sprintf("%s", value)
				case "string":
					id_value = value.(string)
				default:
					panic(col_name + " column data can't convert to expected value ")
				}
			}
		}

		jsonString, err := json.Marshal(row)
		if err != nil {
			return -1, err
		}

		indexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Index:      index.Name,
				Action:     index.DocumentField,
				DocumentID: id_value,
				Body:       strings.NewReader(string(jsonString)),
				OnFailure: func(ctx context.Context, bii esutil.BulkIndexerItem, biri esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						fmt.Println("-------------------------------------")
						log.Fatal(err)
					}
				},
			})
		record_count++
		if record_count%500 == 0 {
			fmt.Println("Record count : ", record_count)
		}
	}

	fmt.Println("Record count : ", record_count)
	indexer.Close(context.Background())
	return record_count, nil
}
