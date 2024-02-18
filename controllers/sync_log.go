package controllers

import (
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindSyncLogs(c *gin.Context) {
	sizeStr := c.DefaultQuery("size", "10")
	pageStr := c.DefaultQuery("page", "0")

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	var filter model.SyncLog
	c.Bind(&filter)
	var syncLogs []model.SyncLog
	dao.DB.Debug().Offset(page * size).Limit(size).Where(&filter).Order("created_at DESC").Find(&syncLogs)

	c.JSON(http.StatusOK, syncLogs)
}
