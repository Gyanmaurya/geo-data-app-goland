
package controllers

import (
	"net/http"
	"geo-data-app/services"
	"io/ioutil"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	geoJSON, err := services.ProcessFile(data)
	if err != nil {
		http.Error(w, "File processing failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(geoJSON)
}
