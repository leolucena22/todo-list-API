package model

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // "pending" ou "completed"
	CreatedAt   time.Time `json:"created_at"`
}
