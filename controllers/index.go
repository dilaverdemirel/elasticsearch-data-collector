package controllers

import (
	"net/http"

	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
	"eslasticsearchdatacollector/scheduler"

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
		Alias:          input.Name,
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

	//TODO eğer index scheduled ise scheduler üzerinden Job schedule edilmeli

	if index.Scheduled {
		scheduler.Add_new_job_to_scheduler_by_index_id(index.ID)
	}

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

	scheduler.Add_new_job_to_scheduler_by_index_id(input.IndexId)

	c.JSON(http.StatusOK, gin.H{"data": dao.DB.Where(&model.Index{ID: input.IndexId}).First})
}

func IndexUnscheduleDataSync(c *gin.Context) {
	id := c.Param("id")

	dao.DB.Save(&model.Index{
		ID:             id,
		CronExpression: "",
		DocumentField:  "",
	})

	scheduler.Delete_job_by_index_id(id)

	c.JSON(http.StatusOK, gin.H{"data": dao.DB.Where(&model.Index{ID: id}).First})
}
