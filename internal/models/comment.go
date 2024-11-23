package models

import "time"

type Comment struct {
	Id        uint      `json:"id"`
	Content   string    `json:"content"`
	PostId    uint      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
