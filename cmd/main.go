package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    // Define a simple handler function
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    // Start the server on port 8080
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}