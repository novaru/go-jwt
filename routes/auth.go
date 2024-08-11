package routes

import (
  "github.com/gorilla/mux"
  "go-jwt/controllers"
)

func AuthRoutes(r *mux.Router) {
  router := r.PathPrefix("/auth").Subrouter()

  router.HandleFunc("/register", controllers.Register).Methods("POST")
  router.HandleFunc("/login", controllers.Login).Methods("POST")
}
