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

	c.JSON(http.StatusOK, index)
}

func DeleteIndexById(c *gin.Context) {
	id := c.Param("id")

	var existsIndex model.Index
	dao.DB.Where(&model.Index{ID: id}).Take(&existsIndex)

	if *existsIndex.Scheduled {
		scheduler.Delete_job_by_index_id(id)
	}

	scheduler.DeleteElasticsearchIndex(id)

	var index model.Index
	dao.DB.Where(&model.Index{ID: id}).Delete(&index)

	c.JSON(http.StatusNoContent, index)
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
		Valid:          &input.Valid,
		SqlQuery:       input.SqlQuery,
		Scheduled:      &input.Scheduled,
		CronExpression: input.CronExpression,
		SyncType:       input.SyncType,
		DataSourceId:   input.DataSourceId,
		DocumentField:  input.DocumentField,
	}
	dao.DB.Create(&index)

	//TODO eğer index scheduled ise scheduler üzerinden Job schedule edilmeli

	if *index.Scheduled {
		scheduler.Add_new_job_to_scheduler_by_index_id(index.ID)
	}

	c.JSON(http.StatusCreated, gin.H{"data": index})
}

func UpdateIndex(c *gin.Context) {
	id := c.Param("id")

	// Validate input
	var input model.UpdateIndexInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dao.DB.Model(&model.Index{ID: id}).Updates(&model.Index{
		ID:            id,
		Name:          input.Name,
		Description:   input.Description,
		Valid:         &input.Valid,
		SqlQuery:      input.SqlQuery,
		SyncType:      input.SyncType,
		DataSourceId:  input.DataSourceId,
		DocumentField: input.DocumentField,
	})

	var index model.Index
	dao.DB.Where(&model.Index{ID: id}).Take(&index)

	c.JSON(http.StatusOK, index)
}

func IndexScheduleDataSync(c *gin.Context) {
	// Validate input
	var input model.ScheduleIndexInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scheduled := true
	dao.DB.Model(&model.Index{ID: input.IndexId}).Updates(&model.Index{
		ID:             input.IndexId,
		CronExpression: input.CronExpression,
		DocumentField:  input.DocumentIdField,
		SyncType:       input.SyncType,
		Scheduled:      &scheduled,
	})

	scheduler.Add_new_job_to_scheduler_by_index_id(input.IndexId)

	var index model.Index
	dao.DB.Where(&model.Index{ID: input.IndexId}).Take(&index)
	c.JSON(http.StatusOK, index)
}

func IndexUnscheduleDataSync(c *gin.Context) {
	id := c.Param("id")

	scheduled := false
	dao.DB.Model(&model.Index{ID: id}).Updates(&model.Index{
		ID:             id,
		CronExpression: "",
		DocumentField:  "",
		Scheduled:      &scheduled,
	})

	scheduler.Delete_job_by_index_id(id)

	var index model.Index
	dao.DB.Where(&model.Index{ID: id}).Take(&index)
	c.JSON(http.StatusOK, index)
}

func GetIndexSyncDailyStatusStats(c *gin.Context) {
	var result []model.SyncDailyStatusStats
	dao.DB.Model(&model.SyncLog{}).Select("date(created_at) as Day, status as Status, count(id) RecordCount").Group("date(created_at), status").Order("date(created_at)").Limit(10).Find(&result)
	c.JSON(http.StatusOK, result)
}

func GetIndexSyncDailyRecordStats(c *gin.Context) {
	var result []model.SyncDailyRecordStats
	dao.DB.Model(&model.SyncLog{}).Select("date(created_at) as Day, sum(document_count) RecordCount").Group("date(created_at), status").Order("date(created_at)").Limit(10).Find(&result)
	c.JSON(http.StatusOK, result)
}

func StartSyncImmetiately(c *gin.Context) {
	id := c.Param("id")
	scheduler.One_time_schedule_by_index_id(id)

	c.Status(http.StatusOK)
}
