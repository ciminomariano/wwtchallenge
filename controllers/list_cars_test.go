package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wwtchallenge/models"
)

func TestGetCarListHandler(t *testing.T) {
	// Clear the in-memory Cars list (if needed)
	models.Cars = []models.Car{
		{ID: 1, Make: "Toyota", Model: "Camry", Package: "Standard", Color: "Blue", Year: 2020, Category: "Sedan", Mileage: 15000, Price: 2500000},
		{ID: 2, Make: "Honda", Model: "Civic", Package: "EX", Color: "Red", Year: 2019, Category: "Sedan", Mileage: 12000, Price: 2200000},
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the GetCarListHandler function to handle the request
	GetCarListHandler(rr, req)

	// Verify the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Decode the response body into a slice of cars
	var carList []models.Car
	err = json.Unmarshal(rr.Body.Bytes(), &carList)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the number of cars returned (in this case, there should be 2)
	expectedCarCount := 2
	if len(carList) != expectedCarCount {
		t.Errorf("Expected %d cars in the list, but got %d", expectedCarCount, len(carList))
	}
}

func TestGetCarByIDHandler(t *testing.T) {
	// Create a new HTTP request to get the car details by ID (e.g., ID 1)
	req, err := http.NewRequest("GET", "/cars/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the GetCarByIDHandler function with the required parameters
	GetCarByIDHandler(rr, req, 1)

	// Verify the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Decode the response body into a car object
	var car models.Car
	err = json.Unmarshal(rr.Body.Bytes(), &car)
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the car details match the expected values
	expectedCar := models.Car{
		ID:       1,
		Make:     "Toyota",
		Model:    "Camry",
		Package:  "Standard",
		Color:    "Blue",
		Year:     2020,
		Category: "Sedan",
		Mileage:  15000,
		Price:    2500000,
	}

	if car != expectedCar {
		t.Errorf("Expected car details to match %+v, but got %+v", expectedCar, car)
	}
}

func TestGetCarByIDHandlerInvalidID(t *testing.T) {
	// Create a new HTTP request to get the details of a car with an invalid ID
	req, err := http.NewRequest("GET", "/cars/xx", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the GetCarByIDHandler function with the required parameters
	GetCarByIDHandler(rr, req, 0) // Provide an invalid ID (e.g., 0)

	// Verify the response status code
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, rr.Code)
	}
}
