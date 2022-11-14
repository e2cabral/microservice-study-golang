package services

import (
	"microservice-golang/src/data/repositories"
	"microservice-golang/src/domain/entities"
)

type ProductService struct{}

func (p ProductService) Create(product entities.Product) error {
	repository, err := repositories.NewProductRepository()
	if err != nil {
		return err
	}

	repository.Create(&product)

	return nil
}
