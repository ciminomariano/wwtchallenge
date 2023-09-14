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
