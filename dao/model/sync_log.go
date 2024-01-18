package model

import "time"

type SyncLogStatus string

const (
	SyncLogStatusStarted  SyncLogStatus = "STARTED"
	SyncLogStatusCompeted SyncLogStatus = "COMPLETED"
	SyncLogStatusFailed   SyncLogStatus = "FAILED"
)

type SyncLog struct {
	ID                string        `json:"ID" gorm:"primary_key"`
	IndexId           string        `json:"IndexId"`
	DocumentCount     int32         `json:"DocumentCount"`
	StartDate         time.Time     `json:"StartDate"`
	EndDate           *time.Time    `json:"EndDate"`
	ExecutionDuration int32         `json:"ExecutionDuration"`
	Status            SyncLogStatus `json:"Status"` // STARTED, COMPLETED, FAILED
	StatusMessage     string        `json:"StatusMessage"`
}
