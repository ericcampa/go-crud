package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"postgres-demo/pkg/api"
	"postgres-demo/pkg/db"
)

func main() {
	pgdb, err := db.NewDb()
	if err != nil {
		panic(err)
	}

	router := api.NewApi(pgdb)

	log.Print("Starting Server")
	port := os.Getenv("PORT")
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("Router error: %v\n", err)
	}

}
