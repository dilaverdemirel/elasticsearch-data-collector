package model

import (
	"database/sql/driver"
	"time"
)

type SyncLogStatus string

const (
	SyncLogStatusStarted  SyncLogStatus = "STARTED"
	SyncLogStatusCompeted SyncLogStatus = "COMPLETED"
	SyncLogStatusFailed   SyncLogStatus = "FAILED"
)

func (ct *SyncLogStatus) Scan(value interface{}) error {
	*ct = SyncLogStatus(value.([]byte))
	return nil
}

func (ct SyncLogStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type SyncLog struct {
	ID                string        `json:"ID" gorm:"primary_key"`
	IndexId           string        `json:"IndexId" gorm:"size:191"`
	DocumentCount     int32         `json:"DocumentCount"`
	StartDate         time.Time     `json:"StartDate"`
	EndDate           *time.Time    `json:"EndDate"`
	ExecutionDuration int32         `json:"ExecutionDuration"`
	Status            SyncLogStatus `json:"Status" gorm:"size:20"`
	StatusMessage     string        `json:"StatusMessage"`
	CreatedAt         time.Time     `json:"CreatedAt" gorm:"autoCreateTime"`
	UpdatedAt         time.Time     `json:"UpdatedAt" gorm:"autoUpdateTime"`
}

type SyncDailyStatusStats struct {
	Day         time.Time     `json:"Day"`
	Status      SyncLogStatus `json:"Status"`
	RecordCount int32         `json:"RecordCount"`
}

type SyncDailyRecordStats struct {
	Day         time.Time `json:"Day"`
	RecordCount int32     `json:"RecordCount"`
}
