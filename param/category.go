package param

import "todoapp/entity"

type CreateCategoryRequest struct {
	Title string `json:"title"`
	Color string `json:"color"`
	UserID uint `json:"-"`
}

type CreateCategoryResponse struct {
	Category entity.Category `json:"category"`
}
