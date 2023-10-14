package routes

import (
	"encoding/json"
	"net/http"

	"github.com/manuelDk8s/go-api/db"
	"github.com/manuelDk8s/go-api/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	db.DB.Find(&user)
	json.NewEncoder(w).Encode(&user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	error := createdUser.Error
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete user"))
}
