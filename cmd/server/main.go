package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/krutik0204/crud-api-go/pkg/api"
	"github.com/krutik0204/crud-api-go/pkg/db"
)

func main() {
	log.Print("server has started")
	//start the db
	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error: %v", err)
		panic("error starting the database")

	}
	//get the router of the API by passing the db
	router := api.StartAPI(pgdb)
	//get the port from the environment variable
	port := os.Getenv("PORT")
	//pass the router and start listening with the server
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
		return
	}
}
