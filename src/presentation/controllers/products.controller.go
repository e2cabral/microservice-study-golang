package controllers

import (
	"github.com/gorilla/mux"
	"microservice-golang/src/domain/entities"
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

	response.Ok(w, &product)
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

	response.Ok(w, products)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var product entities.Product

	response := helpers.Response{}

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	service := services.ProductService{}

	err = helpers.FromJSON(r.Body, &product)
	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	err = service.Update(product, id)
	if err != nil {
		response.InternalServerError(w, err.Error())
	}

	response.Ok(w, product)
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

	response.Ok(w, product)
}
