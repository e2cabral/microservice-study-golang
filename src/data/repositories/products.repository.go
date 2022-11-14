package repositories

import (
	"gorm.io/gorm"
	"microservice-golang/src/domain/entities"
	"microservice-golang/src/infra/database"
)

type ProductRepository struct {
	handler *gorm.DB
}

func NewProductRepository() (*ProductRepository, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	return &ProductRepository{handler: db}, nil
}

func (p ProductRepository) Create(product *entities.Product) {
	p.handler.Create(&product)
}
