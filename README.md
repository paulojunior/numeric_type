# Numeric Type API
This repository contains a simple API for categorizing and managing numeric values. It includes a programmatic module that prints numbers replacing multiples of 3 with "Type 1," multiples of 5 with "Type 2," and common multiples of both with "Type 3," all accomplished with a single if statement. Additionally, the API offers RESTful endpoints for dynamically storing, retrieving, and querying numeric values, allowing clients to seamlessly interact with and manage collections of numbers. The API leverages SOLID principles for robust and extensible design and adheres to REST best practices for simplicity and usability.

## Endpoints
### 1. Get Number
#### Endpoint
```
GET /number/{number}
```

##### Description
Get number information.

##### Parameters
- number (path, string, required)
  
##### Responses
- 200 OK: Successfully retrieved number information.
- 400 Bad Request: Number invalid.
- 404 Not Found: Number not found.
- 500 Internal Server Error: Internal server error.
  
### 2. Get All
#### Endpoint
```
GET /numbers/
```

##### Description
Retrieve information for all numbers
  
##### Responses
- 200 OK: Successfully retrieved all numbers information (array of objects).
- 500 Internal Server Error: Internal server error.
  
### 3. Save a number
#### Endpoint
```
POST /number
```

#### Description
Save a number to the collection.

#### Parameters
- number (body, string, required): Number to save.
  
Request Body
```
{
  "number": "string"
}
```

#### Responses
- 201 Created: number inserted successfully.
- 400 Bad Request: number not found or invalid.
- 409 Conflict: The number already exists in the collection.
- 500 Internal Server Error: Internal server error.
  
#### Dependencies
Echo: A fast and minimalist Go web framework.

#### Getting Started
Clone this repository. Install dependencies using go get. Build and run the application.
