package models

import "time"

type TodoModel struct {
	Id          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Sample Todo"`
	Description string    `json:"description" example:"This is a sample todo item"`
	Completed   bool      `json:"completed" example:"false"`
	CreatedAt   time.Time `json:"created_at" example:"2023-05-23T08:00:00Z"`
}
