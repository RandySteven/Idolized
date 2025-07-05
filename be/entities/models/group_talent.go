package models

import "time"

type GroupTalent struct {
	ID           uint64
	GroupID      string
	TalentID     string
	IsVerified   bool
	VerifiedTime *time.Time
	MemberStatus string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
