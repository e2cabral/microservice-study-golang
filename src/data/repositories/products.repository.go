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

func (p ProductRepository) Find(current int, limit int, products *[]entities.Product) {
	p.handler.
		Limit(limit).
		Offset((current - 1) * limit).
		Find(&products)
}

func (p ProductRepository) Update(product *entities.Product, id int) {
	p.handler.
		Model(&product).
		Where("id = ?", id).
		Updates(
			map[string]interface{}{
				"price":       product.Price,
				"name":        product.Name,
				"description": product.Description,
			},
		)
}
