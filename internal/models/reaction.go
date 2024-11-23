package models

import "time"

type ReactionType string

const (
	Like ReactionType = "like"
)

type Reaction struct {
	Id        uint         `json:"id"`
	Type      ReactionType `json:"type"`
	PostId    uint         `json:"-"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
