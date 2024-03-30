package scheduler

import (
	"context"
	"encoding/json"
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
	"eslasticsearchdatacollector/elasticsearch"
	"eslasticsearchdatacollector/service"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/google/uuid"
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
	record_count, err := migrate_data_to_elasticsearch(&data_source, &index)

	if err != nil {
		service.UpdateSyncLogAsFailed(sync_log.ID, err.Error())
	} else {
		service.UpdateIndexLastSyncDate(index_id, last_sync_execution_date)
		service.UpdateSyncLogAsCompleted(sync_log.ID, record_count)
	}
}

func migrate_data_to_elasticsearch(data_source *sqlx.DB, index *model.Index) (int32, error) {
	sql_meta_data := prepare_sql_query_meta_data(*index)
	rows, err := execute_query(data_source, sql_meta_data)
	if err != nil {
		return -1, err
	}
	colNames, err := rows.Columns()
	if err != nil {
		return -1, err
	}

	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{})
	var id_field string = index.DocumentField
	var record_count int32 = 0
	row_list := dao.ScanRowsWithoutRowLimit(*rows)
	rows.Close()

	var index_name = index.Name
	if index.SyncType == model.IndexSyncTypeReloadAll {
		index_name, err = prepare_index_for_reload_all(*index)
		if err != nil {
			return -1, err
		}
	}

	for l := row_list.Front(); l != nil; l = l.Next() {
		row := make(map[string]interface{})
		var id_value string = ""
		for i := 0; i < len(colNames); i++ {
			col_name := colNames[i]
			value := l.Value.(map[string]interface{})[col_name]
			row[col_name] = value
			if id_field == col_name {
				id_value = ConvertGenericTypeDataToString(value)
			}
		}

		jsonString, err := json.Marshal(row)
		if err != nil {
			return -1, err
		}

		indexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Index:      index_name,
				Action:     "index",
				DocumentID: id_value,
				Body:       strings.NewReader(string(jsonString)),
				OnFailure: func(ctx context.Context, bii esutil.BulkIndexerItem, biri esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						log.Fatal(err)
					}
				},
			})
		record_count++
		if record_count%500 == 0 {
			log.Println("Record count : ", record_count)
		}
	}

	log.Println("Record count : ", record_count)
	log.Println("Index stats, Added :", indexer.Stats().NumAdded,
		",Created : ", indexer.Stats().NumCreated,
		",Deleted : ", indexer.Stats().NumDeleted,
		",Failed : ", indexer.Stats().NumFailed,
		",Flushed : ", indexer.Stats().NumFlushed,
		",Indexed : ", indexer.Stats().NumIndexed,
		",Requests : ", indexer.Stats().NumRequests,
		",Updated : ", indexer.Stats().NumUpdated,
	)
	indexer.Close(context.Background())
	complete_reload_all_data_import(*index, index_name)
	return record_count, nil
}

func ConvertGenericTypeDataToString(value interface{}) string {
	field_data_type := reflect.TypeOf(value).String()
	switch field_data_type {
	case "int", "int8", "int16", "int32", "int64":
		temp := fmt.Sprintf("%v", value)
		return temp
	case "float32", "float64":
		temp := fmt.Sprintf("%v", value)
		return temp
	case "string":
		return value.(string)
	default:
		panic("Data can't convert to expected value!")
	}
}

func prepare_index_for_reload_all(index model.Index) (string, error) {
	var new_index_name = index.Name + strings.ReplaceAll(uuid.NewString(), "-", "")
	var index_create_response, err = elasticsearch.ES.Indices.Create(new_index_name)

	if err != nil {
		return "", err
	}

	if index_create_response.StatusCode != 200 {
		return "", fmt.Errorf("there is an error when create index %s, error is %s", index.Name, index_create_response.String())
	}

	return new_index_name, nil
}

type SQLMetaData struct {
	sql_query               string
	contains_sql_last_value bool
	sql_last_value_count    int
	sql_last_value          time.Time
}

func prepare_sql_query_meta_data(index model.Index) SQLMetaData {
	var sql_query = strings.ToLower(index.SqlQuery)
	var contains_sql_last_value = false
	var sql_last_value_count = 0
	var sql_last_value_pattern = ":#sql_last_value"
	if strings.Contains(sql_query, sql_last_value_pattern) {
		contains_sql_last_value = true
		sql_last_value_count = strings.Count(sql_query, sql_last_value_pattern)
		sql_query = strings.ReplaceAll(sql_query, sql_last_value_pattern, "?")
	}

	var sql_last_value = time.Now()
	if index.LastExecutionTime != nil {
		sql_last_value = *index.LastExecutionTime
	}

	return SQLMetaData{
		sql_query:               sql_query,
		contains_sql_last_value: contains_sql_last_value,
		sql_last_value_count:    sql_last_value_count,
		sql_last_value:          sql_last_value,
	}
}

func execute_query(data_source *sqlx.DB, sql_meta_data SQLMetaData) (*sqlx.Rows, error) {
	log.Printf("sql_last_value_count : %d, contains_sql_last_value : %t, sql_last_value : %s",
		sql_meta_data.sql_last_value_count,
		sql_meta_data.contains_sql_last_value,
		sql_meta_data.sql_last_value.String())

	if sql_meta_data.contains_sql_last_value {
		var params = make([]interface{}, sql_meta_data.sql_last_value_count)
		for i := 0; i < len(params); i++ {
			params[i] = sql_meta_data.sql_last_value
		}
		return data_source.Queryx(sql_meta_data.sql_query, params)
	}

	return data_source.Queryx(sql_meta_data.sql_query)
}

func complete_reload_all_data_import(index model.Index, alias string) error {
	if index.SyncType == model.IndexSyncTypeReloadAll {
		old_alias := index.Alias

		index_exists_response, err := elasticsearch.ES.Indices.Exists([]string{old_alias})

		if err != nil {
			return err
		}

		if index_exists_response.StatusCode == 200 {
			index_delete_response, err := elasticsearch.ES.Indices.Delete([]string{old_alias})
			if err != nil {
				return err
			}
			if index_delete_response.StatusCode == 200 {
				log.Printf("Index deleted successfully %s", old_alias)
			}
		}

		index_put_alias_response, err := elasticsearch.ES.Indices.PutAlias([]string{alias}, index.Name)
		if err != nil {
			return err
		}
		if index_put_alias_response.StatusCode == 200 {
			log.Printf("Alias was added to Index %s, alias %s", old_alias, index.Name)
		}

		service.UpdateIndexAlias(index.ID, alias)
	}

	return nil
}

func DeleteElasticsearchIndex(index_id string) {
	var index = service.GetIndexById(index_id)
	index_delete_response, err := elasticsearch.ES.Indices.Delete([]string{index.Alias})
	if err != nil {
		panic(err)
	}
	if index_delete_response.StatusCode == 200 {
		log.Printf("Index was deleted successfuly")
	}
}
