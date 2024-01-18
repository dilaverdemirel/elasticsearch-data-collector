package service

import (
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
	"time"

	"github.com/google/uuid"
)

func CreateSyncLog(sync_log *model.SyncLog) {
	sync_log.ID = uuid.NewString()
	dao.DB.Create(&sync_log)
}

func GetSyncLogById(id string) model.SyncLog {
	var sync_log model.SyncLog
	dao.DB.Where(&model.SyncLog{ID: id}).Take(&sync_log)

	return sync_log
}

func UpdateSyncLogAsCompleted(id string, record_count int32) {
	var t = time.Now()

	var sync_log = GetSyncLogById(id)
	sync_log.EndDate = &t
	sync_log.Status = model.SyncLogStatusCompeted
	sync_log.DocumentCount = record_count

	dao.DB.Save(sync_log)
}
func UpdateSyncLogAsFailed(id string, error_message string) {
	var t = time.Now()

	var sync_log = GetSyncLogById(id)
	sync_log.EndDate = &t
	sync_log.Status = model.SyncLogStatusFailed
	sync_log.StatusMessage = error_message

	dao.DB.Save(sync_log)
}
