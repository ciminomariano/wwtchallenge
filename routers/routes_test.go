package routers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"wwtchallenge/routers"
)

func TestRouteAvailability(t *testing.T) {
	r := routers.SetupRoutes()

	// Test for route: POST /cars
	jsonBody := []byte(`{"make": "TestMake", "model": "TestModel", "package": "TestPackage", "color": "TestColor", "year": 2022, "category": "TestCategory", "mileage": 0, "price": 30000}`)

	reqPostCar, err := http.NewRequest(http.MethodPost, "/cars", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	rrPostCar := httptest.NewRecorder()
	r.ServeHTTP(rrPostCar, reqPostCar)

	if rrPostCar.Code != http.StatusCreated {
		t.Errorf("Expected status code %d for route POST /cars, but got %d", http.StatusCreated, rrPostCar.Code)
	}

	// Test for route: GET /cars
	reqGetCars, err := http.NewRequest(http.MethodGet, "/cars", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	rrGetCars := httptest.NewRecorder()
	r.ServeHTTP(rrGetCars, reqGetCars)

	if rrGetCars.Code != http.StatusOK {
		t.Errorf("Expected status code %d for route GET /cars, but got %d", http.StatusOK, rrGetCars.Code)
	}

	// Test for route: PUT /cars/1
	reqPutCarByID, err := http.NewRequest(http.MethodPut, "/cars/1", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	rrPutCarByID := httptest.NewRecorder()
	r.ServeHTTP(rrPutCarByID, reqPutCarByID)

	if rrPutCarByID.Code != http.StatusOK {
		t.Errorf("Expected status code %d for route PUT /cars/1, but got %d", http.StatusOK, rrPutCarByID.Code)
	}

	// Test for route: GET /cars/1
	reqGetCarByID, err := http.NewRequest(http.MethodGet, "/cars/1", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	rrGetCarByID := httptest.NewRecorder()
	r.ServeHTTP(rrGetCarByID, reqGetCarByID)

	if rrGetCarByID.Code != http.StatusOK {
		t.Errorf("Expected status code %d for route GET /cars/1, but got %d", http.StatusOK, rrGetCarByID.Code)
	}

}
