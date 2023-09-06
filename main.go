package main

import (
	"context"
	"encoding/json"
	"gotest/config"
	"gotest/db"
	"gotest/handler"
	"gotest/middleware"
	"gotest/service"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env.development")

	serverConfig := config.NewServerConfig(context.Background())
	ds := db.NewMariaDBDataStore(serverConfig.DatabaseConnectionDetails)
	service := service.NewServiceStruct(ds)
	apiHandler := handler.NewAPIHandler(service)

	router := chi.NewRouter()

	router.Use(middleware.LoggerMiddleware)
	router.Use(middleware.MethodMiddleware)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Application is running!")
	})

	router.Group(func(r chi.Router) {
		r.Get("/login", apiHandler.LoginUser)
		r.Post("/register", apiHandler.RegisterUser)
		r.Post("/check", apiHandler.CheckAuthentication)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Get("/me", apiHandler.GetCurrentUser)
		r.Get("/me/grocery", apiHandler.GetAllGroceriesFromUser)
		r.Post("/me/grocery", apiHandler.AddGroceryForUser)
		r.Patch("/me/grocery/{id}", apiHandler.UpdateStatusOfGrocery)
		r.Delete("/me/grocery/{id}", apiHandler.DeleteGroceryFromUser)
		r.Get("/groceries", apiHandler.AllGroceries)
		r.Get("/groceries/{name}", apiHandler.FindAllGroceriesByName)
		r.Post("/groceries", apiHandler.AddNewGrocery)
		r.Put("/groceries", apiHandler.UpadteGroceryById)
		r.Delete("/groceries/{id}", apiHandler.DeleteGrocery)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.AdminMiddleware)
		r.Get("/users", apiHandler.GetAllUsers)
		r.Get("/users/{name}", apiHandler.GetSingleUser)
		r.Get("/users/{role}", apiHandler.GetUserByRole)
		r.Post("/users", apiHandler.CreateNewUser)
		r.Delete("/users/{name}", apiHandler.DeleteUser)
	})

	log.Fatal(http.ListenAndServe(":8083", router))
}
