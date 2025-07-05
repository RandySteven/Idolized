package models

import "time"

type Record struct {
	ID              uint64
	RecordURL       string
	ThumbnailURL    string
	Title           string
	Description     string
	Duration        int32
	RecordStartDate time.Time
	RecordEndDate   time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
