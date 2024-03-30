package main

import (
	"eslasticsearchdatacollector/controllers"
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/elasticsearch"
	"eslasticsearchdatacollector/scheduler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	webApp := gin.Default()
	dao.ConnectDatabase()
	elasticsearch.ConnectElasticsearch()
	scheduler.InitializeSchedulerAndActivateJobs()

	webApp.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	webApp.GET("/datasources", controllers.FindDatasources)
	webApp.GET("/datasources/:id", controllers.GetDatasourceById)
	webApp.POST("/datasources", controllers.CreateDataSource)
	webApp.PUT("/datasources/:id", controllers.UpdateDatasource)
	webApp.DELETE("/datasources/:id", controllers.DeleteDatasourceById)

	webApp.GET("/indices", controllers.FindIndices)
	webApp.GET("/indices/:id", controllers.GetIndexById)
	webApp.POST("/indices", controllers.CreateIndex)
	webApp.DELETE("/indices/:id", controllers.DeleteIndexById)
	webApp.PUT("/indices/:id", controllers.UpdateIndex)
	webApp.PUT("/indices/:id/schedule-data-sync", controllers.IndexScheduleDataSync)
	webApp.DELETE("/indices/:id/unschedule-data-sync", controllers.IndexUnscheduleDataSync)
	webApp.GET("/indices/:id/sync-daily-status-stats", controllers.GetIndexSyncDailyStatusStats)
	webApp.GET("/indices/:id/sync-daily-record-stats", controllers.GetIndexSyncDailyRecordStats)

	webApp.POST("/query-meta-data/preview", controllers.PreviewQueryMetaData)

	webApp.GET("/sync-logs", controllers.FindSyncLogs)

	webApp.Run()
}
