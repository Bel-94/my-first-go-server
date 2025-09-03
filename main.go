// Declares the main package for an executable program
package main

// Import necessary packages from the standard library
import (
	"encoding/json" // For encoding data to JSON
	"fmt"           // For formatting text (like Sprintf)
	"log"           // For logging errors
	"net/http"      // For all HTTP server functionality
	"strings"       // For manipulating strings
)

// Define a struct (a typed collection of fields) to structure our JSON response
type Message struct {
	Text string `json:"message"` // The `json:"message"` tag tells Go to output the field as "message"
}

// Handler function for the root route "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET, if not, send an error
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set the Content-Type header so clients know we're sending JSON
	w.Header().Set("Content-Type", "application/json")

	// Create an instance of our Message struct
	response := Message{Text: "Hello World! Welcome to my first Go server."}

	// Encode the struct into JSON and write it to the response
	json.NewEncoder(w).Encode(response)
}

// Handler function for the "/greet/" route
func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the path after "/greet/"
	// e.g., for "/greet/Alice", path will be "/Alice"
	path := r.URL.Path
	name := strings.TrimPrefix(path, "/greet/")

	// Basic check to see if a name was provided
	if name == "" {
		name = "Anonymous User"
	}

	w.Header().Set("Content-Type", "application/json")
	// Use fmt.Sprintf to create a formatted string
	response := Message{Text: fmt.Sprintf("Hello, %s! Greetings from my Go API.", name)}
	json.NewEncoder(w).Encode(response)
}

// The main function is the entry point of the executable
func main() {
	// Register our handler functions to specific URL patterns
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/greet/", greetHandler)

	// Print a message to the console so we know the server is starting
	log.Println("Starting server on :8080...")

	// Start the HTTP server on port 8080.
	// log.Fatal is used to log an error if the server fails to start (e.g., if the port is already in use).
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}