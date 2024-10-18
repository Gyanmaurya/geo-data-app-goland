package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"geo-data-app/routes"
	"geo-data-app/utils"
)

func main() {
	
	utils.ConnectDatabase()
	fmt.Println("Database connected successfully!")

	router := mux.NewRouter()

	routes.SetupRoutes(router)

	port := ":8080"

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	fmt.Printf("Server is running on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, corsHandler))
}
