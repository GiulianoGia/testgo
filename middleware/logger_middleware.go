package middleware

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			green := color.New(color.BgGreen)
			green.Print(r.Method)
		case "POST":
			yellow := color.New(color.BgYellow)
			yellow.Print(r.Method)
		case "PUT":
			blue := color.New(color.BgBlue)
			blue.Print(r.Method)
		case "PATCH":
			magenta := color.New(color.BgMagenta)
			magenta.Print(r.Method)
		case "DELETE":
			red := color.New(color.BgRed)
			red.Print(r.Method)
		}
		fmt.Println(" ", r.URL)
		next.ServeHTTP(w, r)
		fmt.Println()
	})
}

func MethodMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		next.ServeHTTP(w, r)
	})
}
