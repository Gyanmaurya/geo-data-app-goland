package routes

import (
	"geo-data-app/controllers"
	"github.com/gorilla/mux"
)


func SetupRoutes(router *mux.Router) {
	
	router.HandleFunc("/api/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/upload", controllers.UploadFile).Methods("POST")
	router.HandleFunc("/api/save-shape", controllers.SaveShape).Methods("POST")
	router.HandleFunc("/api/shapes", controllers.GetShapes).Methods("GET")
}
