package entity

import "time"

type Task struct {
	ID uint
	Title string
	DueDate time.Time
	CategoryId uint
	IsDone bool
	UserID uint
}