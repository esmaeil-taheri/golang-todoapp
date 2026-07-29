package param

import "todoapp/entity"

type CreateTaskRequest struct {
	Title string `json:"title"`
	DueDate string `json:"due_date"`
	CategoryID uint `json:"category_id"`
	// UserID is filled from the authenticated session, never from client input,
	// so it's excluded from json (`-`).
	UserID uint `json:"-"`
}

type CreateTaskResponse struct {
	Task entity.Task `json:"task"`
}

type ListTaskRequest struct {
	UserID uint `json:"-"`
}

type ListTaskResponse struct {
	Tasks []entity.Task `json:"tasks"`
}
