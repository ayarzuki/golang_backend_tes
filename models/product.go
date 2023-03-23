package models

import "time"

type Product struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Price       int64     `json:"price"`
	Description string    `gorm:"size:255" json:"description"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
