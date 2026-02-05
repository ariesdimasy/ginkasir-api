package models

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" `
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  *uint     `json:"category_id"`
	Category    Category  `gorm:"foreignkey:CategoryID"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime" `
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// Error implements [error].
func (p *Product) Error() string {
	panic("unimplemented")
}

// Request DTOs
type CreateProductRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=10,max=200"`
	Price       int    `json:"price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=10,max=200"`
	Price       int    `json:"price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
}

type SearchProductRequest struct {
	Name  string `form:"name"`
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
}
