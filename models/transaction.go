package models

import "time"

type Transaction struct {
	ID          uint                `json:"id" gorm:"primaryKey"`
	TotalAmount int                 `json:"total_amount"`
	Details     []TransactionDetail `json:"details"`
	CreatedAt   time.Time           `json:"created_at" gorm:"autoCreateTime" `
	UpdatedAt   time.Time           `json:"updated_at" gorm:"autoUpdateTime"`
}

type TransactionDetail struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	ProductID     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
	Quantity      int    `json:"quantity"`
	Price         int    `json:"price"`
	Subtotal      int    `json:"subtotal"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}
