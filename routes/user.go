package routes

import (
  "github.com/gorilla/mux"
  "go-jwt/controllers"
  "go-jwt/middleware"
)

func UserRoutes(r *mux.Router) {
  router := r.PathPrefix("/user").Subrouter()
  router.Use(middleware.Auth)

  router.HandleFunc("/me", controllers.Me).Methods("GET")
}
