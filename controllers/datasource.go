package controllers

import (
	"net/http"

	"eslasticsearchdatacollector/appenv"
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

	dao.ClearDatabaseConnectionCacheByDatasourceId(datasource.ID)

	c.JSON(http.StatusNoContent, datasource)
}

func CreateDataSource(c *gin.Context) {
	// Validate input
	var input model.CreateDatasourceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pwd := appenv.Encrypt(input.DbPassword)

	// Create datasource
	datasource := model.Datasource{
		ID:               uuid.NewString(),
		Name:             input.Name,
		ConnectionString: input.ConnectionString,
		MaxPoolSize:      input.MaxPoolSize,
		MinIdle:          input.MinIdle,
		UserName:         input.UserName,
		DbPassword:       pwd,
		DriverName:       input.DriverName,
	}
	dao.DB.Create(&datasource)

	c.JSON(http.StatusCreated, datasource)
}

func UpdateDatasource(c *gin.Context) {
	id := c.Param("id")

	// Validate input
	var input model.CreateDatasourceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var datasource model.Datasource
	dao.DB.Where(&model.Datasource{ID: id}).Take(&datasource)

	pwd := input.DbPassword
	if datasource.DbPassword != input.DbPassword {
		pwd = appenv.Encrypt(input.DbPassword)
	}

	dao.DB.Model(&model.Datasource{ID: id}).Updates(&model.Datasource{
		ID:               id,
		Name:             input.Name,
		ConnectionString: input.ConnectionString,
		MaxPoolSize:      input.MaxPoolSize,
		MinIdle:          input.MinIdle,
		UserName:         input.UserName,
		DbPassword:       pwd,
		DriverName:       input.DriverName,
	})

	dao.DB.Where(&model.Datasource{ID: id}).Take(&datasource)

	dao.ClearDatabaseConnectionCacheByDatasourceId(datasource.ID)

	c.JSON(http.StatusOK, datasource)
}
