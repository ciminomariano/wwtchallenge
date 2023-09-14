# Car Microservice - Coding Challenge

This project is a Go-based microservice that provides a basic set of functionalities for managing car data. It offers RESTful endpoints for retrieving and updating car information. Below, you will find details about the project's functionality, prerequisites, installation, and usage.
# Functionality

The Car Microservice exposes the following endpoints:

    GET /cars/{id}
        Retrieves an existing car by its ID.
        Example: /cars/1

    GET /cars
        Retrieves a list of all cars.
        Example: /cars

    POST /cars
        Creates a new car.
        Example: /cars

    Request Body:

    json

{
    "make": "Toyota",
    "model": "Camry",
    "package": "Standard",
    "color": "Blue",
    "year": 2020,
    "category": "Sedan",
    "mileage": 15000,
    "price": 2500000
}

PUT /cars/{id}

    Updates an existing car by its ID.
    Example: /cars/1

Request Body (example):

json

    {
        "make": "Toyota",
        "model": "Camry",
        "package": "Deluxe",
        "color": "Red",
        "year": 2020,
        "category": "Sedan",
        "mileage": 18000,
        "price": 2800000
    }

# Prerequisites

Before running the project, make sure you have the following prerequisites installed:

    Go (Golang)

# Installation

    Clone the repository:

    shell

git clone https://github.com/ciminomariano/wwtchallenge.git

Navigate to the project directory:

shell

cd your-repo

Run the application:

shell

    go run main.go

# Usage

The Car Microservice is now running locally. You can interact with it using a tool like Postman or by making HTTP requests.

    To retrieve a car by ID, use a GET request to /cars/{id} (e.g., /cars/1).
    To retrieve the list of all cars, use a GET request to /cars (e.g., /cars).
    To create a new car, use a POST request to /cars with the car details in the request body.
    To update an existing car, use a PUT request to /cars/{id} with the updated car details in the request body.

Make sure to replace {id} with the desired car's ID when using the endpoints.
# Testing

To run the automated tests, use the following command:

shell

go test ./... from the project directorory same in which you have the server.go file

# TODOs

    Add more test cases for edge scenarios.
    Implement more advanced logging and metrics.    
    Implement API versioning if required in the future.
    Consider performance improvements for larger datasets.

Feel free to discuss the design and code during the interview.
Sample Dataset

A sample dataset for reference:

go

var Cars = []Car{
    {ID: 1, Make: "Toyota", Model: "Camry", Package: "Standard", Color: "Blue", Year: 2020, Category: "Sedan", Mileage: 15000, Price: 2500000},
    {ID: 2, Make: "Honda", Model: "Civic", Package: "EX", Color: "Red", Year: 2019, Category: "Sedan", Mileage: 12000, Price: 2200000},
    // Add more cars as needed
}
