package routers

import (
	"net/http"
	"wwtchallenge/controllers"
)

func SetupRoutes() *http.ServeMux {
	r := http.NewServeMux()

	// Ruta para obtener la lista de coches (GET) y crear un nuevo coche (POST)
	r.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controllers.GetCarListHandler(w, r)
		case http.MethodPost:
			controllers.CreateCarHandler(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return r
}
