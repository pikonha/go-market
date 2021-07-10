package product_ports

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/picolloo/go-market/common/http/handlers"
	product_domain "github.com/picolloo/go-market/product/domain"
	"github.com/picolloo/go-market/product/usecases"
)

func MakeProductHandlers(r *mux.Router, service product_usecase.Service) {
	r.Handle("/products", getProductsHandler(service)).Methods("GET")
	r.Handle("/products/{id}", getProductHandler(service)).Methods("GET")
	r.Handle("/products", storeProductsHandler(service)).Methods("POST")
	r.Handle("/products/{id}", updateProductHandler(service)).Methods("PUT")
	r.Handle("/products/{id}", deleteProductsHandler(service)).Methods("DELETE")
}

func getProductsHandler(service product_usecase.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		products, err := service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(handlers.FormatJSONerror(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(handlers.FormatJSONerror("Error trying to marshal json response"))
		}
	})
}

func storeProductsHandler(service product_usecase.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var p product_domain.Product

		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = service.Store(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		json.NewEncoder(w)
	})
}

func deleteProductsHandler(service product_usecase.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)

		err := service.Delete(int(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(handlers.FormatJSONerror(err.Error()))
			return
		}

		var product product_domain.Product
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(handlers.FormatJSONerror(err.Error()))
		}
	})
}

func updateProductHandler(service product_usecase.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var product *product_domain.Product
		json.NewDecoder(r.Body).Decode(&product)

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		product.ID = int(id)

		err := service.Update(product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(handlers.FormatJSONerror(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(handlers.FormatJSONerror(err.Error()))
		}
	})
}

func getProductHandler(service product_usecase.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id, _ := strconv.ParseInt(vars["id"], 10, 64)
		product, err := service.Get(int(id))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(handlers.FormatJSONerror(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(handlers.FormatJSONerror(err.Error()))
			return
		}
	})
}
