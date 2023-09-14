Car Microservice - Coding Challenge

This is a basic microservice implementation for getting and updating information about cars. It consists of a RESTful OpenAPI contract and a matching HTTP Go implementation.
Endpoints
1. GET /cars/{id}

    Retrieves an existing car by its ID.
    Example: /cars/1

2. GET /cars

    Retrieves a list of all cars.
    Example: /cars

3. POST /cars

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

4. PUT /cars/{id}

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

Implementation Details

    The implementation is done in Golang using standard libraries, including net/http.
    It is assumed to be an internal API with no authentication or authorization requirements.
    For persistence, in-memory storage is used, and car data is not persisted between application runs.
    Logging is implemented for observability.
    Automated testing of the endpoints is included.

Sample Dataset

A sample dataset for reference:

go

var Cars = []Car{
    {ID: 1, Make: "Toyota", Model: "Camry", Package: "Standard", Color: "Blue", Year: 2020, Category: "Sedan", Mileage: 15000, Price: 2500000},
    {ID: 2, Make: "Honda", Model: "Civic", Package: "EX", Color: "Red", Year: 2019, Category: "Sedan", Mileage: 12000, Price: 2200000},
    // Add more cars as needed
}

Running the Application

To run the application, execute the following command:

shell

go run main.go

Testing

To run the automated tests, use the following command:

shell

go test

TODOs

    Add more test cases for edge scenarios.
    Implement more advanced logging and metrics.
    Add input validation and error handling.
    Implement API versioning if required in the future.
    Consider performance improvements for larger datasets.

Feel free to discuss the design and code during the interview.