package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wwtchallenge/models"
)

func TestUpdateCarHandlerValidJSON(t *testing.T) {
	updatedCar := models.Car{
		Make:     "TestMake",
		Model:    "TestModel",
		Package:  "TestPackage",
		Color:    "TestColor",
		Year:     2022,
		Category: "TestCategory",
		Mileage:  0,
		Price:    30000,
	}

	// Encode the car object to JSON
	jsonCar, err := json.Marshal(updatedCar)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the JSON data
	req, err := http.NewRequest("PUT", "/cars/1", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Call the UpdateCarHandler function with the required parameters
	UpdateCarHandler(rr, req, 1, updatedCar) // Provide the car ID (1) and updatedCar

	// Verify the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

}

func TestUpdateCarHandlerInvalidJSON(t *testing.T) {
	// Create an invalid JSON request body
	invalidJSON := []byte(`{"make": "TestMake", "model": "TestModel", "year": "2022"}`)

	// Create a new HTTP request with the invalid JSON data
	req, err := http.NewRequest("PUT", "/cars/1", bytes.NewBuffer(invalidJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Call the UpdateCarHandler function with the required parameters
	UpdateCarHandler(rr, req, 1, models.Car{}) // Provide the car ID (1) and an empty updatedCar

	// Verify the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestUpdateCarHandlerNonExistentID(t *testing.T) {
	updatedCar := models.Car{
		Make:     "TestMake",
		Model:    "TestModel",
		Package:  "TestPackage",
		Color:    "TestColor",
		Year:     2022,
		Category: "TestCategory",
		Mileage:  0,
		Price:    30000,
	}

	// Encode the car object to JSON
	jsonCar, err := json.Marshal(updatedCar)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the JSON data and a non-existent ID
	req, err := http.NewRequest("PUT", "/cars/999", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Call the UpdateCarHandler function with the required parameters
	UpdateCarHandler(rr, req, 999, updatedCar) // Provide a non-existent car ID (999) and updatedCar

	// Verify the response status code
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, rr.Code)
	}
}

func TestUpdateCarHandlerMismatchedIDs(t *testing.T) {
	updatedCar := models.Car{
		Make:     "TestMake",
		Model:    "TestModel",
		Package:  "TestPackage",
		Color:    "TestColor",
		Year:     2022,
		Category: "TestCategory",
		Mileage:  0,
		Price:    30000,
	}

	// Encode the car object to JSON
	jsonCar, err := json.Marshal(updatedCar)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the JSON data and a different car ID in the URL
	req, err := http.NewRequest("PUT", "/cars/2", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Call the UpdateCarHandler function to handle the request
	UpdateCarHandler(rr, req, 2, updatedCar) // Provide a different car ID (2) in the URL

	// Verify the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}
}

func TestGetCarByIDHandlerNonExistentID(t *testing.T) {
	// Create a new HTTP request to get the details of a non-existent car (e.g., ID 999)
	req, err := http.NewRequest("GET", "/cars/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the GetCarByIDHandler function with the required parameters
	GetCarByIDHandler(rr, req, 999)

	// Verify the response status code
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, rr.Code)
	}

	// Verify that the response body contains an error message
	expectedErrorMessage := "Car not found\n"
	if body := rr.Body.String(); body != expectedErrorMessage {
		t.Errorf("Expected error message: %s, but got: %s", expectedErrorMessage, body)
	}
}
