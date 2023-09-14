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

// IsValidCar verifica si un automóvil tiene todos los campos requeridos y devuelve una lista de campos faltantes
func IsValidCar(car models.Car) []string {
	var missingFields []string

	if car.Make == "" {
		missingFields = append(missingFields, "make")
	}
	if car.Model == "" {
		missingFields = append(missingFields, "model")
	}
	if car.Package == "" {
		missingFields = append(missingFields, "package")
	}
	if car.Color == "" {
		missingFields = append(missingFields, "color")
	}
	if car.Category == "" {
		missingFields = append(missingFields, "category")
	}
	if car.Year == 0 {
		missingFields = append(missingFields, "year")
	}
	if car.Mileage < 0 {
		missingFields = append(missingFields, "mileage")
	}
	if car.Price <= 0 {
		missingFields = append(missingFields, "price")
	}

	return missingFields
}

// IsCarDuplicate verifica si ya existe un automóvil con los mismos detalles (marca, modelo y año)
func IsCarDuplicate(newCar models.Car) bool {
	for _, car := range models.Cars {
		if car.Make == newCar.Make && car.Model == newCar.Model && car.Year == newCar.Year {
			return true
		}
	}
	return false
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

// UpdateCar actualiza un automóvil por su ID con los nuevos datos proporcionados.
func UpdateCar(id int, updatedCar models.Car) bool {
	// Buscar el automóvil en la lista de automóviles por su ID
	for i, car := range models.Cars {
		if car.ID == id {
			// Mantein the ID (Does not allow to the user change the id from json)
			updatedCar.ID = models.Cars[i].ID
			models.Cars[i] = updatedCar
			return true
		}
	}
	// Si no se encuentra el automóvil, devolver falso
	return false
}
