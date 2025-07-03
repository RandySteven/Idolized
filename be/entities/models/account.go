package models

import "time"

type Account struct {
	ID           uint64
	Email        string
	Password     string
	IsVerified   bool
	VerifiedTime time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
