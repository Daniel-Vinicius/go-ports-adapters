package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Daniel-Vinicius/go-ports-adapters/adapters/dtos"
	"github.com/Daniel-Vinicius/go-ports-adapters/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	router.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	router.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT", "OPTIONS")

	router.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			write.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(write).Encode(product)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Content-Type", "application/json")
		var productDto dtos.Product
		err := json.NewDecoder(request.Body).Decode(&productDto)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			write.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			write.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(write).Encode(product)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			write.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			write.WriteHeader(http.StatusNotFound)
			return
		}

		productEnabled, err := service.Enable(product)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			write.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(write).Encode(productEnabled)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		write.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			write.WriteHeader(http.StatusNotFound)
			return
		}

		productDisabled, err := service.Disable(product)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			write.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(write).Encode(productDisabled)
		if err != nil {
			write.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
