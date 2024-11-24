package models

import "time"

// create model for a post
type Post struct {
	Id        uint      `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// create model for detailed post
type DetailedPost struct {
	Id        uint       `json:"id"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Reactions []Reaction `json:"reactions"`
	Comments  []Comment  `json:"comments"`
}
