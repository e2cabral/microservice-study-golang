package controllers

import (
	"github.com/gorilla/mux"
	"microservice-golang/src/data/entities"
	"microservice-golang/src/domain/models"
	"microservice-golang/src/domain/services"
	"microservice-golang/src/infra/helpers"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	response := helpers.Response{}

	service := services.ProductService{}

	err := helpers.FromJSON(r.Body, &product)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	err = service.Create(product)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.Ok(w, models.EntityToModel(&product))
}

func Find(w http.ResponseWriter, r *http.Request) {
	response := helpers.Response{}

	query := r.URL.Query()

	current, err := strconv.Atoi(query.Get("current"))
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	service := services.ProductService{}

	products, err := service.Find(current, limit)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	var modelProducts []models.Product

	for _, element := range products {
		modelProducts = append(modelProducts, models.EntityToModel(&element))
	}

	response.Ok(w, modelProducts)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var product entities.Product

	response := helpers.Response{}

	id := mux.Vars(r)["id"]

	service := services.ProductService{}

	err := helpers.FromJSON(r.Body, &product)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	err = service.Update(product, id)
	if err != nil {
		response.InternalServerError(w, err.Error())
	}

	response.Ok(w, models.EntityToModel(&product))
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	response := helpers.Response{}

	service := services.ProductService{}

	product, err := service.FindOne(id)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	if product.ID == 0 {
		response.Ok(w, nil)
		return
	}

	response.Ok(w, models.EntityToModel(product))
}
