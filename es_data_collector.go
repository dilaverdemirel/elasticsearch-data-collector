package main

import (
	"eslasticsearchdatacollector/controllers"
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/elasticsearch"
	"eslasticsearchdatacollector/scheduler"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	webApp := gin.Default()
	dao.ConnectDatabase()
	elasticsearch.ConnectElasticsearch()
	scheduler.InitializeSchedulerAndActivateJobs()

	webApp.GET("/datasources", controllers.FindDatasources)
	webApp.GET("/datasources/:id", controllers.GetDatasourceById)
	webApp.POST("/datasources", controllers.CreateDataSource)

	webApp.GET("/indices", controllers.FindIndices)
	webApp.GET("/indices/:id", controllers.GetDatasourceById)
	webApp.POST("/indices", controllers.CreateIndex)
	webApp.DELETE("/indices", controllers.DeleteIndexById)
	webApp.PUT("/indices/:id/schedule-data-sync", controllers.IndexScheduleDataSync)
	webApp.DELETE("/indices/:id/unschedule-data-sync", controllers.IndexUnscheduleDataSync)

	webApp.POST("/query-meta-data/preview", controllers.PreviewQueryMetaData)

	webApp.Run()
}
