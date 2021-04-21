package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/picolloo/go-market/api/handlers"
	"github.com/picolloo/go-market/core/product"
)

func main() {
	db, err := sql.Open("sqlite3", "data/product.db")
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco %s", err.Error())
	}
	defer db.Close()

	service := product.NewService(db)
	r := mux.NewRouter()

	handlers.MakeProductHandlers(r, service)

	http.Handle("/", r)

	server := &http.Server{
		Addr: ":4000",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
