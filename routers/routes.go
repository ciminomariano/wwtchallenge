package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wwtchallenge/controllers"
	"wwtchallenge/models"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

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

	// Route to get a car by its ID or update it by ID (GET and PUT)
	r.HandleFunc("/cars/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			controllers.GetCarByIDHandler(w, r, id)
		case http.MethodPut:
			// Decode the JSON body of the request into a car structure
			var updatedCar models.Car
			err := json.NewDecoder(r.Body).Decode(&updatedCar)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			controllers.UpdateCarHandler(w, r, id, updatedCar)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return r
}
