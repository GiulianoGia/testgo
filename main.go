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
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
			next.ServeHTTP(w, r)
		})
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Application is running!")
	})

	router.Get("/groceries", apiHandler.AllGroceries)
	router.Get("/groceries/{name}", apiHandler.FindAllGroceriesByName)
	router.Post("/groceries", apiHandler.AddNewGrocery)
	router.Put("/groceries", apiHandler.UpadteGroceryById)
	router.Delete("/groceries/{id}", apiHandler.DeleteGrocery)

	router.Group(func(r chi.Router) {
		r.Get("/login", apiHandler.LoginUser)
		r.Post("/check", apiHandler.CheckAuthentication)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuth)
		r.Get("/me", apiHandler.GetCurrentUser)
		r.Get("/me/grocery", apiHandler.GetAllGroceriesFromUser)
		r.Post("/me/grocery", apiHandler.AddGroceryForUser)
		r.Patch("/me/grocery/{id}", apiHandler.UpdateStatusOfGrocery)
		r.Delete("/me/grocery/{id}", apiHandler.DeleteGroceryFromUser)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.AdminMiddleware)
		r.Get("/users", apiHandler.GetAllUsers)
		r.Get("/users/{name}", apiHandler.GetSingleUser)
		r.Post("/users", apiHandler.CreateNewUser)
		r.Delete("/users/{name}", apiHandler.DeleteUser)
	})

	log.Fatal(http.ListenAndServe(":8083", router))
}
