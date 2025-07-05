package models

import "time"

type Group struct {
	ID           uint64
	AccountID    string
	Email        string
	Password     string
	Name         string
	Description  string
	PictureURL   string
	IsVerified   bool
	VerifiedTime *time.Time
	VerifiedBy   string
	DebutDate    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
