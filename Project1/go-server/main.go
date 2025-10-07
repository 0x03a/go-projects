package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	// if r.URL.Path != "/form" {
	// 	http.Error(w, "404 not found", http.StatusNotFound)
	// 	return
	// }
	// if r.Method != "GET" {
	// 	http.Error(w, "Method is not supported", http.StatusNotFound)
	// 	return
	// }
	// parse the form data

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name") // we are getting the value of the name field from the form
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name =%s\n", name)
	fmt.Fprintf(w, "Address =%s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) { // here request is what the client sends to the server and response is what the server sends back to the client
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// when we write /hello in the browser, it by default sends a GET request to the server
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // we are telling go lang that we want to serve files from the static folder
	http.Handle("/", fileServer)                        // we are telling go lang that we want to handle all requests to the root URL ("/") using the file server we just created (means in a way that we are serving the static files when the user accesses the root URL)
	http.HandleFunc("/form", formHandler)               // we are telling go lang that we want to handle all requests to the /form URL using the formHandler function
	http.HandleFunc("/hello", helloHandler)             // we are telling go lang that we want to handle all requests to the /hello URL using the helloHandler function

	fmt.Println("Starting server on port 4000...")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}
