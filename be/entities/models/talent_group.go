package models

import "time"

type TalentGroup struct {
	ID        uint64
	TalentID  uint64
	GroupID   uint64
	DebutDate time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
