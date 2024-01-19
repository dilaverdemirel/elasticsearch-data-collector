package model

type Datasource struct {
	ID               string `json:"ID" gorm:"primary_key"`
	Name             string `json:"Name" gorm:"size:50"`
	ConnectionString string `json:"ConnectionString" gorm:"size:500"`
	MaxPoolSize      uint   `json:"MaxPoolSize"`
	MinIdle          uint   `json:"MinIdle"`
	UserName         string `json:"UserName" gorm:"size:50"`
	DbPassword       string `json:"DbPassword" gorm:"size:50"`
	DriverName       string `json:"DriverName" gorm:"size:50"`
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
