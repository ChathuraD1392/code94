package models

import "time"

type ReactionType string

// Constants for predefined reaction types.
const (
	Like ReactionType = "like"
)

// create model for reaction
type Reaction struct {
	Id        uint         `json:"id"`
	Type      ReactionType `json:"type"`
	PostId    uint         `json:"-"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
