package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
}
