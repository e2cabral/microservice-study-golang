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

func (p ProductService) Find(current int, limit int) ([]entities.Product, error) {
	var products []entities.Product

	repository, err := repositories.NewProductRepository()
	if err != nil {
		return nil, err
	}

	repository.Find(current, limit, &products)

	return products, nil
}

func (p ProductService) Update(product entities.Product, id int) error {
	repository, err := repositories.NewProductRepository()
	if err != nil {
		return err
	}

	repository.Update(&product, id)

	return nil
}
