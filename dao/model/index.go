package model

import "time"

type Index struct {
	ID                string     `json:"ID" gorm:"primary_key"`
	Name              string     `json:"Name"`
	Alias             string     `json:"Alias"`
	Description       string     `json:"Description"`
	Valid             bool       `json:"Valid"`
	SqlQuery          string     `json:"SqlQuery"`
	Scheduled         bool       `json:"Scheduled"`
	CronExpression    string     `json:"CronExpression"`
	LastExecutionTime *time.Time `json:"LastExecutionTime"`
	SyncType          string     `json:"SyncType"`
	DataSourceId      string     `json:"DataSourceId"`
	DocumentField     string     `json:"DocumentField"`
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
