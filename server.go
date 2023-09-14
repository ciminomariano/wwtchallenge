package main

import (
	"fmt"
	"log"
	"net/http"
	"wwtchallenge/routers"
	// Importa las rutas desde el mismo directorio
)

func main() {

	r := routers.SetupRoutes() // Llama a la función SetupRoutes desde el paquete routes
	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
