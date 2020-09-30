package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("Error: secret not provided")
	}
	fmt.Println(secret)
	//TODO: Handle request & test secret
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	route := os.Getenv("ROUTE")
	port := os.Getenv("PORT")

	fmt.Println(route)
	fmt.Println(port)

	http.HandleFunc(fmt.Sprintf("/%s", route), handleRequest)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
