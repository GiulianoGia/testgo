package main

import (
	"gotest/db"
	"gotest/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db.InitDB()

	router := chi.NewRouter()

	router.Use(middleware.LoggerMiddleware)

	router.Get("/groceries", AllGroceries)
	router.Get("/groceries/{name}", FindAllGroceriesByName)
	router.Post("/groceries", AddNewGrocery)
	router.Put("/groceries", UpadteGroceryById)
	router.Delete("/groceries/{id}", DeleteGrocery)

	log.Fatal(http.ListenAndServe(":8083", router))
}
