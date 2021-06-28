package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/picolloo/go-market/api/handlers"
	"github.com/picolloo/go-market/core/product"
)

func main() {
	db, err := sql.Open("postgres")
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err.Error())
	}
	defer db.Close()

	service := product.NewService(db)

	r := mux.NewRouter()
	handlers.MakeProductHandlers(r, service)
	http.Handle("/", r)

	server := &http.Server{
		Addr: ":9000",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
