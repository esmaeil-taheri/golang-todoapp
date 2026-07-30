package entity

type Task struct {
	ID uint
	Title string
	DueDate string
	CategoryID uint
	IsDone bool
	UserID uint
}