package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manuelDk8s/go-api/db"
	"github.com/manuelDk8s/go-api/models"
	"github.com/manuelDk8s/go-api/routes"
)

func main() {

	// connect db
	db.DbConnection()

	//migrations
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}
