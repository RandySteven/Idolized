package models

import "time"

type Account struct {
	ID         uint64
	Email      string
	Password   string
	UserID     uint64
	RoleID     uint64
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
