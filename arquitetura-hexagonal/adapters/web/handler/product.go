package handler

import (
	"github.com/codeedu/go-hexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"github.com/codeedu/go-hexagonal/adapters/dto"
)

func MakeProductHandlers(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT", "OPTIONS")
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		// Pegar os params do endpoint
		vars := mux.Vars(request)
		id := vars["id"]

		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(writer).Encode(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var productDTO dto.ProductDTO
		err := json.NewDecoder(request.Body).Decode(&productDTO)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}

		product, err := service.Create(productDTO.Name, productDTO.Price)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}

		err = json.NewEncoder(writer).Encode(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		result, err := service.Enable(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(writer).Encode(result)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(request)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			writer.WriteHeader(http.StatusNotFound)
			return
		}

		var productDTO dto.ProductDTO
		err = json.NewDecoder(request.Body).Decode(&productDTO)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}
		err = product.ChangePrice(productDTO.Price)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}

		result, err := service.Disable(product)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(writer).Encode(result)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
