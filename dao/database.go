package dao

import (
	"eslasticsearchdatacollector/dao/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/es-data-collector"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&model.Datasource{})
	if err != nil {
		return
	}

	DB = database
}
