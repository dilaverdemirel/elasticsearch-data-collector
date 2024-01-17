package dao

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

func ScanRows(rows sqlx.Rows, read_row_count int) list.List {
	documents := list.New()

	colNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}

	var recordCount int = 0

	for rows.Next() {
		results := make(map[string]interface{})
		errx := rows.MapScan(results)
		if errx != nil {
			log.Fatal(errx)
		}
		document := make(map[string]interface{})
		for i := 0; i < len(colNames); i++ {
			switch colTypes[i].DatabaseTypeName() {
			case "VARCHAR":
				document[colNames[i]] = string(results[colNames[i]].([]uint8))
			case "DECIMAL":
				x := fmt.Sprintf("%s", results[colNames[i]])
				theFloat, _ := strconv.ParseFloat(strings.TrimSpace(x), 64)
				document[colNames[i]] = theFloat
			case "INT":
				x := fmt.Sprintf("%s", results[colNames[i]])
				theInt, _ := strconv.ParseInt(strings.TrimSpace(x), 32, 64)
				document[colNames[i]] = theInt
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
		documents.PushBack(document)

		if read_row_count != -1 && recordCount >= read_row_count {
			break
		}
	}
	fmt.Println("Record count : ", recordCount)

	return *documents
}
