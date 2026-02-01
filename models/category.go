package models

import (
	"time"
)

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Request DTOs
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}

type SearchCategoryRequest struct {
	Name  string `form:"name"`
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
}
