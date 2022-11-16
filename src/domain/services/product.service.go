package services

import (
	"errors"
	"microservice-golang/src/data/entities"
	"microservice-golang/src/data/repositories"
	"microservice-golang/src/infra/helpers"
	"strconv"
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

func (p ProductService) Update(product entities.Product, id string) error {
	if helpers.IsEmpty(id) {
		return errors.New("null parameter: You should provide an Id")
	}

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	repository, err := repositories.NewProductRepository()
	if err != nil {
		return err
	}

	repository.Update(&product, convertedId)

	return nil
}

func (p ProductService) FindOne(id string) (*entities.Product, error) {
	if helpers.IsEmpty(id) {
		return nil, errors.New("null parameter: You should provide an Id")
	}
	var product entities.Product

	convertedId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	repository, err := repositories.NewProductRepository()
	if err != nil {
		return nil, err
	}

	repository.FindOne(convertedId, &product)

	return &product, nil
}
