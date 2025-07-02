package models

import "time"

type Talent struct {
	ID             uint64
	UserID         uint64
	StageName      string
	TalentPhotoURL string
	RelationStatus string
	CreatedAt      time.Time
	CreatedBy      uint64
	UpdatedAt      time.Time
	UpdatedBy      uint64
	DeletedAt      *time.Time
	DeletedBy      *uint64
}
