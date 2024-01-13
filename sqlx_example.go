package main

import (
	"context"
	"encoding/json"
	"eslasticsearchdatacollector/controllers"
	"eslasticsearchdatacollector/dao"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	webApp := gin.Default()
	dao.ConnectDatabase()

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	info, err := es.Info()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cluseter info : ", info)

	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/es-data-collector")

	if err != nil {
		log.Fatal(err)
	}

	// fetch all places from the db
	rows, err := db.Queryx("SELECT * FROM customers")
	if err != nil {
		log.Fatal(err)
	}
	colNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}

	// iterate over each row
	indexer, _ := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{})

	var idField string = "id"

	var recordCount int = 0

	for rows.Next() {
		results := make(map[string]interface{})
		errx := rows.MapScan(results)
		if errx != nil {
			log.Fatal(errx)
		}
		//fmt.Println(results)
		document := make(map[string]interface{})
		var idValue string = ""
		for i := 0; i < len(colNames); i++ {
			//fmt.Println(colNames[i], results[colNames[i]], colTypes[i].DatabaseTypeName())
			/*
				VARCHAR
				DATETIME
				DECIMAL
				INT
			*/
			switch colTypes[i].DatabaseTypeName() {
			case "VARCHAR":
				//fmt.Println(colNames[i], string(results[colNames[i]].([]uint8)))
				document[colNames[i]] = string(results[colNames[i]].([]uint8))
				if idField == colNames[i] {
					idValue = string(results[colNames[i]].([]uint8))
				}
			case "DECIMAL":
				x := fmt.Sprintf("%s", results[colNames[i]])
				theFloat, _ := strconv.ParseFloat(strings.TrimSpace(x), 64)
				//fmt.Println(colNames[i], x, theFloat, reflect.TypeOf(theFloat))
				document[colNames[i]] = theFloat
				if idField == colNames[i] {
					idValue = string(results[colNames[i]].([]uint8))
				}
			case "INT":
				x := fmt.Sprintf("%s", results[colNames[i]])
				theInt, _ := strconv.ParseInt(strings.TrimSpace(x), 32, 64)
				//fmt.Println(colNames[i], x, theInt, reflect.TypeOf(theInt))
				document[colNames[i]] = theInt
				if idField == colNames[i] {
					idValue = string(results[colNames[i]].([]uint8))
				}
			case "DATETIME":
				x := fmt.Sprintf("%s", results[colNames[i]])
				theTime, err := time.Parse("2006-01-02 03:04:05", x)
				if err != nil {
					fmt.Println("Could not parse time:", err)
				}
				//fmt.Println(colNames[i], x, theTime, reflect.TypeOf(theTime))
				document[colNames[i]] = theTime
			default:
				fmt.Println(colNames[i], "-")
			}

		}

		recordCount++
		if recordCount%500 == 0 {
			fmt.Println("Record count : ", recordCount)
		}

		jsonString, err := json.Marshal(document)
		if err != nil {
			log.Fatal(err)
		}
		indexer.Add(
			context.Background(),
			esutil.BulkIndexerItem{
				Index:      "sqlx",
				Action:     "index",
				DocumentID: idValue,
				Body:       strings.NewReader(string(jsonString)),
				OnFailure: func(ctx context.Context, bii esutil.BulkIndexerItem, biri esutil.BulkIndexerResponseItem, err error) {
					if err != nil {
						fmt.Println("-------------------------------------")
						log.Fatal(err)
					}
				},
			})

		// es.Index(
		// 	"sqlx",
		// 	bytes.NewReader(jsonString),
		// )
	}
	fmt.Println("Record count : ", recordCount)

	indexer.Close(context.Background())

	// check the error from rows

	webApp.GET("/datasources", controllers.FindDatasources)       // new
	webApp.GET("/datasources/:id", controllers.GetDatasourceById) // new
	webApp.POST("/datasources", controllers.CreateDataSource)     // new

	webApp.GET("/indices", controllers.FindIndices)           // new
	webApp.GET("/indices/:id", controllers.GetDatasourceById) // new
	webApp.POST("/indices", controllers.CreateIndex)          // new

	webApp.Run()
}
