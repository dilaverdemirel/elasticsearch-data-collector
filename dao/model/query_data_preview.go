package model

type QueryPreviewInput struct {
	DataSourceId string `json:"DataSourceId" binding:"required"`
	Query        string `json:"Query" binding:"required"`
}

type FieldMetaData struct {
	FieldName string
	DataType  string
}

type QueryPreviewOutput struct {
	ExampleData  []map[string]interface{}
	MetaDataList []FieldMetaData
}
