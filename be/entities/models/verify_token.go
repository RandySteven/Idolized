package models

import "time"

type VerifyToken struct {
	ID         uint64
	UserID     uint64
	Token      string
	ExpiryTime time.Time
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
