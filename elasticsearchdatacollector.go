package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

var myMap = make(map[string]interface{})
var mySlice = make([]map[string]interface{}, 0)

func main1() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/es-data-collector")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id,first_name,last_name,subscription_date,expense_amount FROM customers where id = 1")

	if err != nil {
		log.Fatal(err)
		return
	}

	//var firstname, lastname string
	var count int

	arr, err := rows.ColumnTypes()
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i].ScanType())
	}

	fmt.Println(rows.ColumnTypes())

	colNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	cols := make([]interface{}, len(colNames))
	colPtrs := make([]interface{}, len(colNames))
	for i := 0; i < len(colNames); i++ {
		colPtrs[i] = &cols[i]
	}

	for rows.Next() {
		err = rows.Scan(colPtrs...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("cols: ", cols)

		for i, col := range cols {
			myMap[colNames[i]] = col
		}
		mySlice = append(mySlice, myMap)

		// Do something with the map
		for key, val := range myMap {
			fmt.Println("Key:", key, "Value:", val, "Value Type:", reflect.TypeOf(val))
		}
		fmt.Println("myMap: ", myMap)
		fmt.Println("mySlice: ", mySlice)
	}

	fmt.Println(fmt.Sprintf("%s", mySlice[0]["id"]))
	fmt.Println(fmt.Sprintf("%s", mySlice[0]["first_name"]))
	fmt.Println(fmt.Sprintf("%s", mySlice[0]["subscription_date"]))
	fmt.Println(fmt.Sprintf("%s", mySlice[0]["expense_amount"]))
	fmt.Println("Total record count : ", count)

	defer db.Close()
}
