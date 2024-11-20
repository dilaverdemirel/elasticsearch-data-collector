package dao

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

func ScanRowsWithoutRowLimit(rows sqlx.Rows, ch chan map[string]interface{}) {
	ScanRows(rows, -1, ch)
}

func ScanRows(rows sqlx.Rows, read_row_count int, ch chan map[string]interface{}) {
	colNames, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Println(err)
	}

	var recordCount int = 0

	for rows.Next() {
		results := make(map[string]interface{})
		errx := rows.MapScan(results)
		if errx != nil {
			log.Println(errx)
		}
		document := make(map[string]interface{})
		for i := 0; i < len(colNames); i++ {
			if results[colNames[i]] == nil {
				continue
			}
			//println(reflect.TypeOf(results[colNames[i]]).Kind())
			var interfaceType = reflect.TypeOf(results[colNames[i]]).Kind()
			//println(colTypes[i].DatabaseTypeName())
			switch colTypes[i].DatabaseTypeName() {
			case "VARCHAR", "TEXT", "NCHAR", "CHAR", "LongVarChar", "LONGVARCHAR":
				if interfaceType == reflect.Slice {
					document[colNames[i]] = string(results[colNames[i]].([]uint8))
				} else {
					document[colNames[i]] = results[colNames[i]].(string)
				}
			case "DECIMAL", "FLOAT4", "FLOAT8", "NUMBER":
				if interfaceType == reflect.Float32 || interfaceType == reflect.Float64 {
					document[colNames[i]] = results[colNames[i]].(float64)
				} else {
					x := fmt.Sprintf("%s", results[colNames[i]])
					theFloat, _ := strconv.ParseFloat(strings.TrimSpace(x), 64)
					document[colNames[i]] = theFloat
				}

			case "INT", "INT4":
				if interfaceType == reflect.Int64 {
					document[colNames[i]] = results[colNames[i]].(int64)
				} else {
					x := fmt.Sprintf("%s", results[colNames[i]])
					theInt, _ := strconv.ParseInt(strings.TrimSpace(x), 32, 64)
					document[colNames[i]] = theInt
				}
			case "DATE", "DATETIME":
				if interfaceType == reflect.Struct {
					document[colNames[i]] = results[colNames[i]].(time.Time)
				} else {
					x := fmt.Sprintf("%s", results[colNames[i]])
					theTime, err := time.Parse("2006-01-02 03:04:05", x)
					if err != nil {
						fmt.Println("Could not parse time:", err)
					}
					//fmt.Println(colNames[i], x, theTime, reflect.TypeOf(theTime))
					document[colNames[i]] = theTime
				}
			default:
				fmt.Println(colNames[i], "-")
			}
		}

		recordCount++
		ch <- document

		if read_row_count != -1 && recordCount >= read_row_count {
			break
		}
	}
	close(ch)
	fmt.Println("Record count : ", recordCount)
}
