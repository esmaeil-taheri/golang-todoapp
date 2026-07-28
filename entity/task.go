package entity

import "time"

type Task struct {
	ID int
	Title string
	DueDate time.Time
	CategoryId int
	IsDone bool
	UserID int
}