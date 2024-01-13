module eslasticsearchdatacollector

go 1.16

require (
	github.com/elastic/go-elasticsearch/v7 v7.17.10
	github.com/gin-gonic/gin v1.9.1
	github.com/go-sql-driver/mysql v1.7.1
	github.com/google/uuid v1.5.0
	github.com/jmoiron/sqlx v1.3.5
	gorm.io/driver/mysql v1.5.2
	gorm.io/gorm v1.25.5
)

replace github.com/dilaverdemirel/eslasticsearchdatacollector => ../eslasticsearchdatacollector
