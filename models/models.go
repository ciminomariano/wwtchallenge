// models/car.go

package models

// Car representa la estructura de datos de un coche
type Car struct {
	ID       int    `json:"id"`
	Make     string `json:"make"`
	Model    string `json:"model"`
	Package  string `json:"package"`
	Color    string `json:"color"`
	Year     int    `json:"year"`
	Category string `json:"category"`
	Mileage  int    `json:"mileage"`
	Price    int    `json:"price"`
}

var Cars = []Car{
	// Aquí puedes inicializar tu lista de coches
	// Por ejemplo:
	{ID: 1, Make: "Toyota", Model: "Camry", Package: "Standard", Color: "Blue", Year: 2020, Category: "Sedan", Mileage: 15000, Price: 2500000},
	{ID: 2, Make: "Honda", Model: "Civic", Package: "EX", Color: "Red", Year: 2019, Category: "Sedan", Mileage: 12000, Price: 2200000},
}
