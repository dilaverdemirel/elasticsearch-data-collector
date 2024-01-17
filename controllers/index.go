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

func DeleteIndexById(c *gin.Context) {
	id := c.Param("id")
	var index model.Index
	dao.DB.Where(&model.Index{ID: id}).Delete(&index)

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

	c.JSON(http.StatusCreated, gin.H{"data": index})
}

func IndexScheduleDataSync(c *gin.Context) {
	// Validate input
	var input model.ScheduleIndexInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dao.DB.Save(&model.Index{
		ID:             input.IndexId,
		CronExpression: input.CronExpression,
		DocumentField:  input.DocumentIdField,
		SyncType:       input.SyncType,
	})

	c.JSON(http.StatusOK, gin.H{"data": dao.DB.Where(&model.Index{ID: input.IndexId}).First})
}

func IndexUnscheduleDataSync(c *gin.Context) {
	id := c.Param("id")

	dao.DB.Save(&model.Index{
		ID:             id,
		CronExpression: "",
		DocumentField:  "",
	})

	c.JSON(http.StatusOK, gin.H{"data": dao.DB.Where(&model.Index{ID: id}).First})
}
