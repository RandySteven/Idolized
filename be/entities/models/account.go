package models

import "time"

type Account struct {
	ID              uint64
	AccountID       string
	ParentAccountID string
	AccountRole     string
	IsVerified      bool
	VerifiedTime    *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
