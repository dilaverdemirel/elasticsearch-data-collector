package controllers

import (
	"fmt"
	"log"
	"net/http"

	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"

	"github.com/gin-gonic/gin"
)

func PreviewQueryMetaData(c *gin.Context) {
	var input model.QueryPreviewInput
	c.Bind(&input)

	db_connection := dao.ConnectDatabaseWithDefinedDatasource(input.DataSourceId)

	fmt.Println("query : ", input.Query)
	rows, err := db_connection.Queryx(input.Query)

	if err != nil {
		log.Println(err)
	}

	colNames, err := rows.Columns()
	if err != nil {
		log.Println(err)
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Println(err)
	}

	var meta_output model.QueryPreviewOutput
	meta_items := make([]model.FieldMetaData, len(colNames))
	for i := 0; i < len(colNames); i++ {
		var dataType = colTypes[i].DatabaseTypeName()
		switch colTypes[i].DatabaseTypeName() {
		case "INT", "INT4", "BIGINT", "DECIMAL", "FLOAT4", "FLOAT8", "NUMBER":
			dataType = "NUMERIC"
		case "VARCHAR", "NVARCHAR", "BOOL", "NCHAR", "CHAR", "LongVarChar", "LONGVARCHAR":
			dataType = "TEXT"
		}
		var meta model.FieldMetaData
		meta.DataType = dataType
		meta.FieldName = colNames[i]
		meta_items[i] = meta
	}
	meta_output.MetaDataList = meta_items

	ch := make(chan map[string]interface{})
	go dao.ScanRows(*rows, 100, ch)

	meta_data_items := make([]map[string]interface{}, 100)
	var row_index = 0
	for {
		l, ok := <-ch
		if !ok {
			break
		}
		row := make(map[string]interface{})
		for i := 0; i < len(colNames); i++ {
			col_name := colNames[i]
			row[col_name] = l[col_name]
		}
		meta_data_items[row_index] = row
		row_index++
	}

	rows.Close()
	meta_output.ExampleData = meta_data_items[:row_index]

	//----------------------------

	/*
		var output model.QueryPreviewOutput

		var meta1 model.FieldMetaData
		meta1.DataType = "CHAR"
		meta1.FieldName = "ID"

		var meta2 model.FieldMetaData
		meta2.DataType = "CHAR"
		meta2.FieldName = "name"

		items := make([]model.FieldMetaData, 2)
		items[0] = meta1
		items[1] = meta2
		output.MetaDataList = items

		dataItems := make([]map[string]interface{}, 2)
		row1 := make(map[string]interface{})
		row1["ID"] = "1"
		row1["name"] = "Dilaver Demirel"
		dataItems[0] = row1

		row2 := make(map[string]interface{})
		row2["ID"] = "2"
		row2["name"] = "Ediz Demirel"
		dataItems[1] = row2

		output.ExampleData = dataItems
	*/

	c.JSON(http.StatusOK, meta_output)
}
