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

	// Decode the JSON body of the request into a car structure
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate a new unique ID for the car, e.g., using a counter or a random number generator
	// For simplicity, you can use a counter that increments with each new car created.
	// Note: In a real-world application, you might want to use a more robust method to generate IDs.
	//newCar.ID = generateUniqueCarID()

	// Call the service function to create the car
	services.CreateCar(newCar)

	// Return the JSON response with the newly created car
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
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
