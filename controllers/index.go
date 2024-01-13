package controllers

import (
	"net/http"

	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindIndices(c *gin.Context) {
	var filter model.Index
	c.Bind(&filter)
	var indices []model.Index
	dao.DB.Where(&filter).Find(&indices)

	c.JSON(http.StatusOK, gin.H{"data": indices})
}

func GetIndexById(c *gin.Context) {
	id := c.Param("id")
	var index model.Index
	dao.DB.Where(&model.Index{ID: id}).Take(&index)

	c.JSON(http.StatusOK, gin.H{"data": index})
}

func CreateIndex(c *gin.Context) {
	// Validate input
	var input model.CreateIndexInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index := model.Index{
		ID:             uuid.NewString(),
		Name:           input.Name,
		Description:    input.Description,
		Valid:          input.Valid,
		SqlQuery:       input.SqlQuery,
		Scheduled:      input.Scheduled,
		CronExpression: input.CronExpression,
		SyncType:       input.SyncType,
		DataSourceId:   input.DataSourceId,
		DocumentField:  input.DocumentField,
	}
	dao.DB.Create(&index)

	c.JSON(http.StatusOK, gin.H{"data": index})
}
