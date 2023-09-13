package routers

import (
	"net/http"
	"wwtchallenge/controllers"
)

func SetupRoutes() *http.ServeMux {
	r := http.NewServeMux()

	// Ruta para crear un nuevo coche (POST)
	r.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controllers.CreateCarHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Configura otras rutas aquí

	return r
}
