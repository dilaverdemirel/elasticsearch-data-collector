package dao

import (
	"eslasticsearchdatacollector/appenv"
	"eslasticsearchdatacollector/dao/model"
	"eslasticsearchdatacollector/gormlock"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/sijms/go-ora/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var datasource_map = make(map[string]sqlx.DB)

func ConnectDatabase() {

	var dbConnectionString = os.Getenv("ES_DATA_COLLECTOR_APP_DB_CONNECTION_STRING")
	var waitToConnect = os.Getenv("ES_DATA_COLLECTOR_APP_DB_WAIT_CONNECTION")
	if dbConnectionString == "" {
		dbConnectionString = "root:root@tcp(127.0.0.1:3306)/es-data-collector?parseTime=true"
	}

	if waitToConnect != "" {
		time.Sleep(60 * time.Second)
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

func ClearDatabaseConnectionCacheByDatasourceId(datasource_id string) {
	temp_datasource, ok := datasource_map[datasource_id]
	if ok {
		temp_datasource.Close()
		delete(datasource_map, datasource_id)
	}
}

func ConnectDatabaseWithDefinedDatasource(datasource_id string) sqlx.DB {
	var datasource model.Datasource
	temp_datasource, ok := datasource_map[datasource_id]
	if !ok {
		DB.Where(&model.Datasource{ID: datasource_id}).Take(&datasource)

		pwd := appenv.Decrypt(datasource.DbPassword)

		connectionString := datasource.ConnectionString

		connectionString = strings.ReplaceAll(connectionString, "#USER#", datasource.UserName)
		connectionString = strings.ReplaceAll(connectionString, "#PWD#", pwd)

		db, err := sqlx.Open(datasource.DriverName, connectionString)
		if datasource.DriverName == "mysql" {
			db, err = sqlx.Open(datasource.DriverName,
				datasource.UserName+":"+pwd+
					"@"+connectionString)
		}

		if datasource.DriverName == "oracle" {
			connectionString := "oracle://" + datasource.UserName + ":" + pwd + "@" + datasource.ConnectionString
			db, err = sqlx.Open(datasource.DriverName, connectionString)
		}

		// 	db, err := sqlx.Open("godror", `user="system" password="oracle" connectString="localhost:49161/xe"
		// poolSessionTimeout=42s configDir="/home/dilaverdemirel/Downloads/instantclient-basic-linux.x64-23.5.0.24.07/instantclient_23_5/"
		// heterogeneousPool=false standaloneConnection=false`)
		db.SetMaxIdleConns(int(datasource.MinIdle))
		db.SetMaxOpenConns(int(datasource.MaxPoolSize))

		if err != nil {
			log.Println(err)
		} else {
			datasource_map[datasource_id] = *db
			temp_datasource = *db
		}
		return temp_datasource
	}

	return temp_datasource
}
