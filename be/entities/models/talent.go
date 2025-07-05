package models

import "time"

type Talent struct {
	ID              uint64
	AccountID       string
	Email           string
	Password        string
	StageName       string
	TalentPhotoURL  string
	RelationStatus  string
	DateOfBirth     string
	HideDateOfBirth bool
	IsVerified      bool
	VerifiedTime    *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
