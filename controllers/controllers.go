package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"wwtchallenge/models"
	"wwtchallenge/services"
)

func CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body of the request into a car structure
	var newCar models.Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the car has all required fields
	missingFields := services.IsValidCar(newCar)
	if len(missingFields) > 0 {
		errorMessage := "Missing required fields: " + strings.Join(missingFields, ", ")
		errorResponse := map[string]string{"error": errorMessage}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Check if the car is a duplicate
	if services.IsCarDuplicate(newCar) {
		http.Error(w, "Car already exists", http.StatusConflict)
		return
	}

	// Call the service function to create the car and obtain the car with the assigned ID
	createdCar := services.CreateCar(newCar)

	// Return the JSON response with the newly created car
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

func UpdateCarHandler(w http.ResponseWriter, r *http.Request, id int, updatedCar models.Car) {

	// Check if the car has all required fields
	missingFields := services.IsValidCar(updatedCar)
	if len(missingFields) > 0 {
		errorMessage := "Missing required fields: " + strings.Join(missingFields, ", ")
		errorResponse := map[string]string{"error": errorMessage}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Llamar a la función de servicio para actualizar el coche
	updated := services.UpdateCar(id, updatedCar)
	if !updated {
		// Coche no encontrado, devolver una respuesta 404
		errorResponse := map[string]string{"error": "Car not found"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	// Devolver una respuesta exitosa con un mensaje
	response := map[string]string{"message": "Car updated successfully"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
