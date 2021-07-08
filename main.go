package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	// "github.com/picolloo/go-market/api/handlers"
	"github.com/picolloo/go-market/product/infra"
	"github.com/picolloo/go-market/product/use_cases"
	"github.com/picolloo/go-market/product/domain"
)

func main() {
	dbUrl := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "docker", "docker", "go-market",
	)

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %s", err.Error())
	}
	defer db.Close()

  postgresProductRepo := product_infra.NewPostgresProductRepository(db)
	service := product_use_case.NewService(postgresProductRepo)

  product := product_domain.Product{
    Name: "toalha",
    Price: 20.5,
    Type: product_domain.Food,
  }
  service.Store(&product)

	// r := mux.NewRouter()
	// handlers.MakeProductHandlers(r, service)
	// http.Handle("/", r)

	// server := &http.Server{
	// 	Addr: ":9000",
	// }

	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatal(err)
	// }
}
