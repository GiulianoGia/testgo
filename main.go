package main

import (
	"gotest/db"
	"gotest/handler"
	auth "gotest/handler/auth"
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

	router.Get("/groceries", handler.AllGroceries)
	router.Get("/groceries/{name}", handler.FindAllGroceriesByName)
	router.Post("/groceries", handler.AddNewGrocery)
	router.Put("/groceries", handler.UpadteGroceryById)
	router.Delete("/groceries/{id}", handler.DeleteGrocery)

	router.Group(func(r chi.Router) {
		r.Get("/login", auth.LoginUser)
		r.Post("/check", auth.CheckAuthentication)
	})

	router.Group(func(r chi.Router) {
		r.Post("/me/grocery", handler.AddGroceryForUser)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth)
		r.Get("/users", handler.GetAllUsers)
		r.Get("/users/{name}", handler.GetSingleUser)
		r.Post("/users", handler.CreateNewUser)
		r.Delete("/users/{name}", handler.DeleteUser)
	})

	log.Fatal(http.ListenAndServe(":8083", router))
}
