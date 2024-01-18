package service

import (
	"eslasticsearchdatacollector/dao"
	"eslasticsearchdatacollector/dao/model"
	"time"
)

func FindIndices(filter *model.Index) []model.Index {
	var indices []model.Index
	dao.DB.Where(&filter).Find(&indices)

	return indices
}

func GetIndexById(id string) model.Index {
	var index model.Index
	dao.DB.Where(&model.Index{ID: id}).Take(&index)

	return index
}

func UpdateIndexLastSyncDate(id string, last_sync_date time.Time) {
	var index = GetIndexById(id)
	index.LastExecutionTime = &last_sync_date
	dao.DB.Save(index)
}
