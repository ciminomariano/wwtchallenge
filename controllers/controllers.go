package controllers

import (
	"encoding/json"
	"net/http"
	"wwtchallenge/models"
	"wwtchallenge/services"
)

// CreateCarHandler handles the creation of a new car
func CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	var newCar models.Car

	// Decodifica el cuerpo JSON de la solicitud en una estructura de coche
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Llama a la función de servicio para crear el coche y obtener el coche con el ID asignado
	createdCar := services.CreateCar(newCar)

	// Devuelve la respuesta JSON con el coche recién creado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdCar)
}

// GetCarListHandler handles the GET request to retrieve the list of cars
func GetCarListHandler(w http.ResponseWriter, r *http.Request) {
	// Call the service function to get the list of cars
	carList := services.GetCarList()

	// Return the list of cars in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(carList)
}

// GetCarByIDHandler handles the GET request to retrieve a car by its ID

// GetCarByIDHandler handles the GET request to retrieve a car by its ID
func GetCarByIDHandler(w http.ResponseWriter, r *http.Request, id int) {
	// Convert the ID to the appropriate data type (e.g., int) and use it to retrieve the car
	car := services.GetCarByID(id)

	if car == nil {
		// Car not found, return a 404 response
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	// Return the car in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}
