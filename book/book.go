package book

import (
	"time"
)

type Book struct {
	Title        string
	Author       string
	Description  string
	Completed    bool
	Created_at   time.Time
	Completed_at *time.Time
}

func NewBook(title, author, description string, completed bool, created_at time.Time) *Book {
	return &Book{
		Title:       title,
		Author:      author,
		Description: description,
		Completed:   completed,
		Created_at:  created_at,
	}
}
