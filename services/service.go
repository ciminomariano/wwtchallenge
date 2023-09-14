package services

import (
	"wwtchallenge/models"
)

var lastCarID int // Track the last assigned car ID

func init() {
	// Initialize the last car ID based on the existing cars
	lastCarID = 0
	for _, car := range models.Cars {
		if car.ID > lastCarID {
			lastCarID = car.ID
		}
	}
}

// CreateCar creates a new car and adds it to the list of cars
func CreateCar(car models.Car) models.Car {

	// Generate a new unique ID for the car, e.g., using a counter or a random number generator
	//newCar.ID = generateUniqueCarID()
	// Increment the last car ID and assign it to the new car
	lastCarID++
	car.ID = lastCarID

	// Add the new car to the list
	models.Cars = append(models.Cars, car)

	return car
}

// GetCarList returns the list of cars
func GetCarList() []models.Car {
	// Return the list of cars stored in the models package
	return models.Cars
}

// GetCarByID obtiene un coche por su ID
func GetCarByID(carID int) *models.Car {
	for _, car := range models.Cars {
		if car.ID == carID {
			return &car
		}
	}
	return nil
}
