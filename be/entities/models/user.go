package models

import "time"

type User struct {
	ID                uint64
	UserID            string
	AccountID         string
	Email             string
	Password          string
	FullName          string
	UserName          string
	DateOfBirth       string
	Status            string
	Gender            string
	ProfilePictureURL string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time
}
