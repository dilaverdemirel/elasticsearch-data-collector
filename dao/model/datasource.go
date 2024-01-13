package model

type Datasource struct {
	ID               string `json:"ID" gorm:"primary_key"`
	Name             string `json:"Name"`
	ConnectionString string `json:"ConnectionString"`
	MaxPoolSize      uint   `json:"MaxPoolSize"`
	MinIdle          uint   `json:"MinIdle"`
	UserName         string `json:"UserName"`
	DbPassword       string `json:"DbPassword"`
	DriverName       string `json:"DriverName"`
}

type CreateDatasourceInput struct {
	Name             string `json:"Name" binding:"required"`
	ConnectionString string `json:"ConnectionString" binding:"required"`
	MaxPoolSize      uint   `json:"MaxPoolSize" binding:"required"`
	MinIdle          uint   `json:"MinIdle" binding:"required"`
	UserName         string `json:"UserName" binding:"required"`
	DbPassword       string `json:"DbPassword" binding:"required"`
	DriverName       string `json:"DriverName" binding:"required"`
}
