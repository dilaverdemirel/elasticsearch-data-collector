package controllers

import (
	"net/http"

	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
	"eslasticsearchdatacollector/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindDatasources(c *gin.Context) {
	var filter model.Datasource
	c.Bind(&filter)
	c.JSON(http.StatusOK, gin.H{"data": service.FindDatasources(&filter)})
}

func GetDatasourceById(c *gin.Context) {
	id := c.Param("id")
	var datasource model.Datasource
	dao.DB.Where(&model.Datasource{ID: id}).Take(&datasource)

	c.JSON(http.StatusOK, datasource)
}

func DeleteDatasourceById(c *gin.Context) {
	id := c.Param("id")

	var datasource model.Datasource
	dao.DB.Where(&model.Datasource{ID: id}).Delete(&datasource)

	c.JSON(http.StatusNoContent, datasource)
}

func CreateDataSource(c *gin.Context) {
	// Validate input
	var input model.CreateDatasourceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create datasource
	datasource := model.Datasource{
		ID:               uuid.NewString(),
		Name:             input.Name,
		ConnectionString: input.ConnectionString,
		MaxPoolSize:      input.MaxPoolSize,
		MinIdle:          input.MinIdle,
		UserName:         input.UserName,
		DbPassword:       input.DbPassword,
		DriverName:       input.DriverName,
	}
	dao.DB.Create(&datasource)

	c.JSON(http.StatusOK, gin.H{"data": datasource})
}
