# Go Movies CRUD API

A beginner-friendly REST API backend for managing movies, built with Go and Gorilla Mux. This project demonstrates basic CRUD (Create, Read, Update, Delete) operations on movie data stored in memory, using JSON for communication.

## Project Overview

This API allows you to:
- List all movies
- Get details of a single movie
- Add a new movie
- Update an existing movie
- Delete a movie

All data is stored in memory (no database), making it ideal for learning Go web development and RESTful API basics.

## Features

- Simple in-memory data storage using Go structs and slices
- JSON-based API requests and responses
- CRUD operations for movie resources
- Uses Gorilla Mux for routing

## Tech Stack

- **Language:** Go
- **Router:** Gorilla Mux
- **Data Format:** JSON

## API Endpoints

| Method | Endpoint           | Description                |
|--------|--------------------|----------------------------|
| GET    | `/movies`          | List all movies            |
| GET    | `/movies/{id}`     | Get a movie by ID          |
| POST   | `/movies`          | Add a new movie            |
| PUT    | `/movies/{id}`     | Update a movie by ID       |
| DELETE | `/movies/{id}`     | Delete a movie by ID       |

### Movie JSON Structure

```json
{
  "id": "1",
  "isbn": "12345",
  "title": "Inception",
  "director": {
    "firstname": "Christopher",
    "lastname": "Nolan"
  }
}
```

## How to Run

1. **Install Go** (if not already installed):  
   [Download Go](https://golang.org/dl/)

2. **Install Gorilla Mux:**
   ```
   go get -u github.com/gorilla/mux
   ```

3. **Run the server:**
   ```
   go run main.go
   ```

4. **Server will start at:**  
   `http://localhost:8000`

## Example Requests

### List all movies
```bash
curl http://localhost:8000/movies
```

### Get a movie by ID
```bash
curl http://localhost:8000/movies/1
```

### Add a new movie
```bash
curl -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{"isbn":"12345","title":"Inception","director":{"firstname":"Christopher","lastname":"Nolan"}}'
```

### Update a movie
```bash
curl -X PUT http://localhost:8000/movies/1 \
  -H "Content-Type: application/json" \
  -d '{"isbn":"54321","title":"Interstellar","director":{"firstname":"Christopher","lastname":"Nolan"}}'
```

### Delete a movie
```bash
curl -X DELETE http://localhost:8000/movies/1
```

## Testing with Postman

You can also test all API endpoints using [Postman](https://www.postman.com/):

- Import the endpoints manually or use the above URLs.
- For POST and PUT requests, set the request body type to **raw** and select **JSON**.
- Example POST body:
  ```json
  {
    "isbn": "12345",
    "title": "Inception",
    "director": {
      "firstname": "Christopher",
      "lastname": "Nolan"
    }
  }
  ```
- Send requests and view responses directly in Postman.

## Notes

- Data is stored in memory and resets when the server restarts.
- This project is for educational purposes and does not use persistent