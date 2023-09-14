package main

import (
	"fmt"
	"log"
	"net/http"
	"wwtchallenge/routers"
)

func StartServer() {
	r := routers.SetupRoutes()
	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}

func main() {
	StartServer()
}
