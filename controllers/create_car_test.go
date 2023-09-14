package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wwtchallenge/models"
)

func TestCreateCarHandlerValidJSON(t *testing.T) {
	// Create a new car object for the test
	newCar := models.Car{
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
	jsonCar, err := json.Marshal(newCar)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the JSON data
	req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Call the CreateCarHandler function to handle the request
	CreateCarHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, rr.Code)
	}

	// Decode the response body into a car object
	var createdCar models.Car
	err = json.Unmarshal(rr.Body.Bytes(), &createdCar)
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the car was created correctly
	if createdCar.ID == 0 {
		t.Errorf("Expected non-zero car ID, but got 0")
	}

	// Clean up: Remove the created car from the list
	for i, car := range models.Cars {
		if car.ID == createdCar.ID {
			// Remove the car from the list
			models.Cars = append(models.Cars[:i], models.Cars[i+1:]...)
			break
		}
	}
}

func TestCreateCarHandlerInvalidJSON(t *testing.T) {
	// Create an invalid JSON
	invalidJSON := []byte(`{"make": "TestMake", "model": "TestModel", "year": "2022"}`)

	// Create a new HTTP request with the invalid JSON data
	req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(invalidJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Call the CreateCarHandler function to handle the request
	CreateCarHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestCreateDuplicateCarHandler(t *testing.T) {
	// Create a new car object for the test
	newCar := models.Car{
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
	jsonCar, err := json.Marshal(newCar)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the JSON data
	req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Call the CreateCarHandler function to handle the request
	CreateCarHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, rr.Code)
	}

	// Try to create the same car again
	req, err = http.NewRequest("POST", "/cars", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")

	// Call the CreateCarHandler function to handle the request
	CreateCarHandler(rr, req)

	// Check the response status code for a duplicate car
	if rr.Code != http.StatusConflict {
		t.Errorf("Expected status code %d for a duplicate car, but got %d", http.StatusConflict, rr.Code)
	}
}

func TestCreateCarHandlerMissingFields(t *testing.T) {
	// Create a JSON object with missing fields
	newCarWithIncompleteFields := models.Car{
		Make:     "TestMake",
		Year:     2022,
		Category: "TestCategory",
		Mileage:  0,
		Price:    1,
	}

	// Encode the car object to JSON
	jsonCar, err := json.Marshal(newCarWithIncompleteFields)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP request with the JSON data
	req, err := http.NewRequest("POST", "/cars", bytes.NewBuffer(jsonCar))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Call the CreateCarHandler function to handle the request
	CreateCarHandler(rr, req)

	// Check the response status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
	}

	// Decode the response body into a JSON object
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Verify that the response contains an error message
	expectedErrorMessage := "Missing required fields: model, package, color"
	if response["error"] != expectedErrorMessage {
		t.Errorf("Expected error message: %s, but got: %s", expectedErrorMessage, response["error"])
	}
}
