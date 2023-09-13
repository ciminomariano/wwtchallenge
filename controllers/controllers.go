package controllers

import (
	"encoding/json"
	"net/http"
	// Importa el paquete de servicios
)

// Estructura para representar un coche (puedes personalizarla según tus necesidades)
type Car struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Color string `json:"color"`
}

// Lista de coches simulada (para la demostración)
var cars = []Car{
	{ID: 1, Brand: "Toyota", Model: "Camry", Year: 2020, Color: "Blue"},
	{ID: 2, Brand: "Honda", Model: "Civic", Year: 2019, Color: "Red"},
}

// CreateCarHandler maneja la creación de un nuevo coche
func CreateCarHandler(w http.ResponseWriter, r *http.Request) {
	var newCar Car

	// Decodifica el cuerpo JSON de la solicitud en una estructura de coche
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simula el proceso de asignación de un nuevo ID (esto se puede hacer mejor)
	newCar.ID = len(cars) + 1

	// Agrega el nuevo coche a la lista (simulada)
	cars = append(cars, newCar)

	// Puedes llamar a una función de servicio para realizar operaciones adicionales si es necesario
	// services.CreateCar(newCar)

	// Devuelve la respuesta JSON con el coche recién creado
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}
