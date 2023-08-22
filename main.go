package main

import (
	"encoding/json"
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
	godotenv.Load(".env.development")

	db.InitDB()

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
		r.Get("/me", auth.GetCurrentUser)
		r.Get("/me/grocery", handler.GetAllGroceriesFromUser)
		r.Post("/me/grocery", handler.AddGroceryForUser)
		r.Patch("/me/grocery/{id}", handler.UpdateStatusOfGrocery)
		r.Delete("/me/grocery/{id}", handler.DeleteGroceryFromUser)
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
