
package controllers

import (
	"encoding/json"
	"net/http"
	"geo-data-app/models"
	"geo-data-app/services"
)

func SaveShape(w http.ResponseWriter, r *http.Request) {
	var shape models.GeoShape
	_ = json.NewDecoder(r.Body).Decode(&shape)

	err := services.SaveShape(shape)
	if err != nil {
		http.Error(w, "Shape saving failed", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetShapes(w http.ResponseWriter, r *http.Request) {
	shapes, err := services.GetAllShapes()
	if err != nil {
		http.Error(w, "Error retrieving shapes", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shapes)
}
