package model

import (
	"database/sql/driver"
	"time"
)

type IndexSyncType string

const (
	IndexSyncTypeIterative IndexSyncType = "ITERATIVE"
	IndexSyncTypeReloadAll IndexSyncType = "RELOAD_ALL"
)

func (ct *IndexSyncType) Scan(value interface{}) error {
	*ct = IndexSyncType(value.([]byte))
	return nil
}

func (ct IndexSyncType) Value() (driver.Value, error) {
	return string(ct), nil
}

type Index struct {
	ID                string        `json:"ID" gorm:"primary_key"`
	Name              string        `json:"Name" gorm:"size:100"`
	Alias             string        `json:"Alias" gorm:"size:100"`
	Description       string        `json:"Description" gorm:"size:150"`
	Valid             bool          `json:"Valid"`
	SqlQuery          string        `json:"SqlQuery"`
	Scheduled         bool          `json:"Scheduled"`
	CronExpression    string        `json:"CronExpression" gorm:"size:50"`
	LastExecutionTime *time.Time    `json:"LastExecutionTime"`
	SyncType          IndexSyncType `json:"SyncType" gorm:"size:20"`
	DataSourceId      string        `json:"DataSourceId" gorm:"size:191"`
	DocumentField     string        `json:"DocumentField" gorm:"size:50"`
	CreatedAt         time.Time     `json:"CreatedAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time     `json:"UpdatedAt" gorm:"autoUpdateTime"`
}

type CreateIndexInput struct {
	Name           string        `json:"Name" binding:"required"`
	Description    string        `json:"Description"`
	Valid          bool          `json:"Valid"`
	SqlQuery       string        `json:"SqlQuery" binding:"required"`
	Scheduled      bool          `json:"Scheduled"`
	CronExpression string        `json:"CronExpression"`
	SyncType       IndexSyncType `json:"SyncType" binding:"required"`
	DataSourceId   string        `json:"DataSourceId" binding:"required"`
	DocumentField  string        `json:"DocumentField"`
}

type ScheduleIndexInput struct {
	CronExpression  string        `json:"CronExpression" binding:"required"`
	DocumentIdField string        `json:"DocumentIdField" binding:"required"`
	IndexId         string        `json:"IndexId" binding:"required"`
	SyncType        IndexSyncType `json:"SyncType" binding:"required"`
}
