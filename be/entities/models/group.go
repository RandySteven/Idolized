package models

import "time"

type Group struct {
	ID          uint64
	Name        string
	Description string
	PictureURL  string
	DebutDate   time.Time
	CreatedAt   time.Time
	CreatedBy   uint64
	UpdatedAt   time.Time
	UpdatedBy   uint64
	DeletedAt   *time.Time
	DeletedBy   uint64
}
