package models

import "time"

type Post struct {
	CreatedAt time.Time
	Title     string
	Content   string
	ID        int
}
