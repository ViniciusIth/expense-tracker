package models

import "time"

type GroupMember struct {
	GroupID  int       `json:"group_id"`
	UserID   int       `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
}
