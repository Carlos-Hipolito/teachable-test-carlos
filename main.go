package main

import (
	"fmt"
	"net/http"
	"teachable-test/src/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./src/config/.env")
	if err != nil {
		fmt.Println("error trying to load environment variables: ", err)
		return
	}
	router := mux.NewRouter()
	routes.StartRoutes(router)

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("error trying to start server: ", err)
		return
	}
}
