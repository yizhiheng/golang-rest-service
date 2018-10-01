package main

import (
	"log"
	"net/http"

	"github.com/yizhiheng/golang-rest-service/resource"

	"github.com/gorilla/handlers"
)

func main() {
	router := resource.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
