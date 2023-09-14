package services

import (
	"testing"
	"wwtchallenge/models"
)

func TestIsValidCar(t *testing.T) {
	// Test a valid car
	validCar := models.Car{
		Make:     "Toyota",
		Model:    "Camry",
		Package:  "Standard",
		Color:    "Blue",
		Year:     2020,
		Category: "Sedan",
		Mileage:  15000,
		Price:    2500000,
	}

	missingFields := IsValidCar(validCar)
	if len(missingFields) > 0 {
		t.Errorf("Expected no missing fields for a valid car, but got: %v", missingFields)
	}

	// Test an invalid car with missing fields
	invalidCar := models.Car{}
	missingFields = IsValidCar(invalidCar)
	expectedMissingFields := []string{"make", "model", "package", "color", "category", "year", "price"}
	if !stringSlicesEqual(missingFields, expectedMissingFields) {
		t.Errorf("Expected missing fields %v for an invalid car, but got: %v", expectedMissingFields, missingFields)
	}
}

func TestIsCarDuplicate(t *testing.T) {
	// Initialize the list of cars with some existing cars
	models.Cars = []models.Car{
		{ID: 1, Make: "Toyota", Model: "Camry", Year: 2020},
		{ID: 2, Make: "Honda", Model: "Civic", Year: 2019},
	}

	// Test a car that is a duplicate
	duplicateCar := models.Car{Make: "Toyota", Model: "Camry", Year: 2020}
	if !IsCarDuplicate(duplicateCar) {
		t.Errorf("Expected car to be a duplicate, but it's not")
	}

	// Test a car that is not a duplicate
	uniqueCar := models.Car{Make: "Ford", Model: "Focus", Year: 2022}
	if IsCarDuplicate(uniqueCar) {
		t.Errorf("Expected car to be unique, but it's a duplicate")
	}
}

func stringSlicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
