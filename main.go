package main

import (
  "github.com/gorilla/mux"
  "go-jwt/configs"
  "go-jwt/routes"
  "log"
  "net/http"
)

func main() {
  configs.ConnectDB()

  r := mux.NewRouter()
  router := r.PathPrefix("/api").Subrouter()

  routes.AuthRoutes(router)
  routes.UserRoutes(router)

  log.Println("Starting server on port 8080")
  err := http.ListenAndServe(":8080", router)
  if err != nil {
    return
  }
}
