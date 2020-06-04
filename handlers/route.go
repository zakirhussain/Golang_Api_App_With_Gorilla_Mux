package handlers

import (
	c "../controllers"
	middle "../middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func New() http.Handler {
	route := mux.NewRouter()
	// Auth middleware check
	authRoute := route.PathPrefix("/auth").Subrouter()
	authRoute.Use(middle.SimpleMiddleware)
	//authRoute.Use(middle.AuthMiddleware)
	authRoute.HandleFunc("/users", c.GetAllUsers)
	authRoute.HandleFunc("/user/{id}", c.GetUser)
	/*	route.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello World")
	})*/
	return route
}
