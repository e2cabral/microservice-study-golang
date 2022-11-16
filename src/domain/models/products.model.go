package models

import (
	"gorm.io/gorm"
	"microservice-golang/src/data/entities"
	"time"
)

type Product struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ID          uint           `json:"id"`
	Price       float64        `json:"price"`
	CreatedAt   time.Time      `json:"createdAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

func EntityToModel(product *entities.Product) Product {
	return Product{
		Name:        product.Name,
		Description: product.Description,
		ID:          product.ID,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		DeletedAt:   product.DeletedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}
