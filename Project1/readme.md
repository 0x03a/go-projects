# Basic Web Server Documentation

This project demonstrates a simple web server built with Go. The server serves static files and handles basic HTTP requests.

## Project Structure

- `go-server/`
  - `main.go`: Main Go source file for the server.
  - `static/`: Directory containing static HTML files.
    - `index.html`: Homepage for the static website.
    - `form.html`: HTML form for submitting data.

## How the Server Works

### Static File Serving

The server uses Go's `http.FileServer` to serve files from the `static` directory. When users visit the root URL (`/`), they receive the `index.html` file.

### Form Handling

- The form in [`static/form.html`](go-server/static/form.html) sends a POST request to `/form` with `name` and `address` fields.
- The [`formHandler`](go-server/main.go) function in [`main.go`](go-server/main.go) processes this request:
  - Parses the form data.
  - Responds with the submitted name and address.

### Hello Endpoint

- Visiting `/hello` triggers the [`helloHandler`](go-server/main.go) function.
- Only GET requests are accepted.
- Responds with a simple "hello!" message.

## Running the Server

1. Make sure you have Go installed.
2. Run the server from the `go-server` directory:
   ```sh
   go run main.go
   ```
3. Open your browser and visit:
   - `http://localhost:4000/` for the static homepage.
   - `http://localhost:4000/form` for the form.
   - `http://localhost:4000/hello` for the hello endpoint.

## Summary

- Static files are served from the `static` folder.
- `/form` handles form submissions.
- `/hello` responds with a greeting.

Refer to [`main.go`](go-server/main.go) for the server implementation and [`static/`](go-server/static/) for HTML
