# Go CRUD API

This is a simple CRUD (Create, Read, Update, Delete) API for managing a collection of movies, built using Go and the Gorilla Mux router. This project demonstrates basic API operations and can be expanded or integrated with a database for more advanced use.

## Features

- **Get all movies**: Retrieve a list of all movies in the collection.
- **Get a single movie**: Retrieve a specific movie by its ID.
- **Create a new movie**: Add a new movie to the collection.
- **Update a movie**: Modify the details of an existing movie.
- **Delete a movie**: Remove a movie from the collection.

## Endpoints

- `GET /api/movies`: Get all movies.
- `GET /api/movies/{id}`: Get a specific movie by ID.
- `POST /api/movies`: Create a new movie.
- `PUT /api/movies/{id}`: Update an existing movie by ID.
- `DELETE /api/movies/{id}`: Delete a movie by ID.

## Getting Started

### Prerequisites

- Go 1.16 or later installed on your machine.

### Installation

1. Clone the repository:

```sh
   git clone https://github.com/aydenjahola/go-crud-api.git
   cd go-crud-api
```

### Running the API

To run the API server locally, execute:

```sh
go run main.go
```

This will start the server on `http://localhost:8000`. You can use tools like `curl`, `Postman`, or your web browser to interact with the API.

### Example Requests

#### Get all movies

```sh
curl -X GET http://localhost:8000/api/movies
```

#### Get a specific movie

```sh
curl -X GET http://localhost:8000/api/movies/1
```

#### Create a new movie

```sh
curl -X POST http://localhost:8000/api/movies -H "Content-Type: application/json" -d '{"isbn":"123456","title":"New Movie","director":{"firstname":"Jane","lastname":"Doe"}}'
```

#### Update a movie

```sh
curl -X PUT http://localhost:8000/api/movies/1 -H "Content-Type: application/json" -d '{"isbn":"654321","title":"Updated Movie","director":{"firstname":"Jane","lastname":"Doe"}}'
```

#### Delete a movie

```sh
curl -X DELETE http://localhost:8000/api/movies/1
```

### Project Structure

- `main.go`: The entry point of the application. It sets up the routes and starts the server.
- `Movie` and `Director` structs: Represent the movie and its director.
- CRUD operations: Implemented in separate functions for readability and organization.

### Future Enhancements

- **Database Integration**: Replace the hardcoded data with a persistent database.
- **Error Handling**: Improve error handling and validation for API inputs.
- **Authentication**: Implement user authentication for secure API access.
