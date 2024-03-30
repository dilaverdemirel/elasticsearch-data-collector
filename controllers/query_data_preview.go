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

	var meta_output model.QueryPreviewOutput
	meta_items := make([]model.FieldMetaData, len(colNames))
	for i := 0; i < len(colNames); i++ {
		var meta model.FieldMetaData
		meta.DataType = colTypes[i].DatabaseTypeName()
		meta.FieldName = colNames[i]
		meta_items[i] = meta
	}
	meta_output.MetaDataList = meta_items

	row_list := dao.ScanRows(*rows, 100)
	rows.Close()
	//row_list.Remove(nil)

	meta_data_items := make([]map[string]interface{}, row_list.Len())
	var row_index = 0
	for l := row_list.Front(); l != nil; l = l.Next() {
		row := make(map[string]interface{})
		for i := 0; i < len(colNames); i++ {
			col_name := colNames[i]
			row[col_name] = l.Value.(map[string]interface{})[col_name]
		}
		meta_data_items[row_index] = row
		row_index++
	}
	meta_output.ExampleData = meta_data_items

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
