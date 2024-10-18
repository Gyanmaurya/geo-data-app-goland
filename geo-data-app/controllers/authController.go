
package controllers

import (
	"encoding/json"
	"net/http"
	"geo-data-app/services"
	"geo-data-app/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	err := services.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	_ = json.NewDecoder(r.Body).Decode(&credentials)

	token, err := services.AuthenticateUser(credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
