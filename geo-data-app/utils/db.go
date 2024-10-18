
package utils

import (
	"fmt"
	"database/sql"
	
	_ "github.com/lib/pq" 
)

var DB *sql.DB


func ConnectDatabase() {
	connStr := "user=postgres dbname=geospatial_app sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
	} else {
		fmt.Println("Database connection established.")
	}
}
