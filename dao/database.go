package dao

import (
	"eslasticsearchdatacollector/appenv"
	"eslasticsearchdatacollector/dao/model"
	"eslasticsearchdatacollector/gormlock"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var datasource_map = make(map[string]sqlx.DB)

func ConnectDatabase() {

	var dbConnectionString = os.Getenv("ES_DATA_COLLECTOR_APP_DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		dbConnectionString = "root:root@tcp(127.0.0.1:3306)/es-data-collector?parseTime=true"
	}

	database, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&model.Datasource{}, &model.Index{}, &model.SyncLog{}, &gormlock.CronJobLock{})
	if err != nil {
		return
	}

	DB = database
}

func ConnectDatabaseWithDefinedDatasource(datasource_id string) sqlx.DB {
	var datasource model.Datasource
	temp_datasource, ok := datasource_map[datasource_id]
	if !ok {
		DB.Where(&model.Datasource{ID: datasource_id}).Take(&datasource)

		pwd := appenv.Decrypt(datasource.DbPassword)

		db, err := sqlx.Open(datasource.DriverName,
			datasource.UserName+":"+pwd+
				"@"+datasource.ConnectionString)
		db.SetMaxIdleConns(int(datasource.MinIdle))
		db.SetMaxOpenConns(int(datasource.MaxPoolSize))

		if err != nil {
			log.Fatal(err)
		} else {
			datasource_map[datasource_id] = *db
			temp_datasource = *db
		}
		return temp_datasource
	}

	return temp_datasource
}
