package model

import "time"

type Index struct {
	ID                string     `json:"ID" gorm:"primary_key"`
	Name              string     `json:"Name" gorm:"size:100"`
	Alias             string     `json:"Alias" gorm:"size:100"`
	Description       string     `json:"Description" gorm:"size:150"`
	Valid             bool       `json:"Valid"`
	SqlQuery          string     `json:"SqlQuery"`
	Scheduled         bool       `json:"Scheduled"`
	CronExpression    string     `json:"CronExpression" gorm:"size:50"`
	LastExecutionTime *time.Time `json:"LastExecutionTime"`
	SyncType          string     `json:"SyncType" gorm:"size:20"`
	DataSourceId      string     `json:"DataSourceId" gorm:"size:191"`
	DocumentField     string     `json:"DocumentField" gorm:"size:50"`
}

type CreateIndexInput struct {
	Name           string `json:"Name" binding:"required"`
	Description    string `json:"Description"`
	Valid          bool   `json:"Valid" binding:"required"`
	SqlQuery       string `json:"SqlQuery" binding:"required"`
	Scheduled      bool   `json:"Scheduled" binding:"required"`
	CronExpression string `json:"CronExpression"`
	SyncType       string `json:"SyncType" binding:"required"`
	DataSourceId   string `json:"DataSourceId" binding:"required"`
	DocumentField  string `json:"DocumentField"`
}

type ScheduleIndexInput struct {
	CronExpression  string `json:"CronExpression" binding:"required"`
	DocumentIdField string `json:"DocumentIdField" binding:"required"`
	IndexId         string `json:"IndexId" binding:"required"`
	SyncType        string `json:"SyncType" binding:"required"`
}
