package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    fmt.Printf("Server listening at port 8080\n")

    // Create a home page handler
    http.HandleFunc("/", homePageHandler)
    // Create a hello handler
    http.HandleFunc("/hello", helloHandler)
    // Create an auth handler
    http.HandleFunc("/auth", authHandler)

    // Start the server, if any error, console log it
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

// Handler for the hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Page!")
}

// Handler for the home route
func homePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Home Page!")
}

// Handler for the auth route
func authHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    // Authentication logic here
    fmt.Fprintf(w, "Authenticated!")
}
