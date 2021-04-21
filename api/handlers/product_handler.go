package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/picolloo/go-market/core/product"
)

func MakeProductHandlers(r *mux.Router, service product.UseCase) {
	r.Handle("/v1/products", getProductsHandler(service)).Methods("GET")
	r.Handle("/v1/products/{id}", getProductHandler(service)).Methods("GET")
	r.Handle("/v1/products", storeProductsHandler(service)).Methods("POST")
	r.Handle("/v1/products/{id}", updateProductHandler(service)).Methods("PUT")
	r.Handle("/v1/products/{id}", deleteProductsHandler(service)).Methods("DELETE")
}

func getProductsHandler(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		products, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONerror(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONerror("Error trying to marshal json response"))
		}
	})
}

func storeProductsHandler(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var p product.Product

		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		product, err := service.Store(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		json.NewEncoder(w).Encode(product)
	})
}

func deleteProductsHandler(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application;json")

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)

		product, err := service.Delete(int(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONerror(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONerror(err.Error()))
		}
	})
}

func updateProductHandler(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var product *product.Product
		json.NewDecoder(r.Body).Decode(&product)

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		product.ID = int(id)

		product, err := service.Update(product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONerror(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONerror(err.Error()))
		}
	})
}

func getProductHandler(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		product, err := service.Get(int(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONerror(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONerror(err.Error()))
			return
		}
	})
}
