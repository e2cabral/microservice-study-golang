package controllers

import (
	"microservice-golang/src/domain/entities"
	"microservice-golang/src/domain/services"
	"microservice-golang/src/infra/helpers"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var product entities.Product

	service := services.ProductService{}

	err := helpers.FromJSON(r.Body, &product)
	if err != nil {
		helpers.InternalServerError(w, err.Error())
	}

	err = service.Create(product)
	if err != nil {
		helpers.InternalServerError(w, err.Error())
	}

	helpers.Ok(w, &product)
}
