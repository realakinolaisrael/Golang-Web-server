package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
	"Web-server/internal/config"
    "Web-server/internal/handler"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
    })
}

func authenticationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Example: Check for a token in the header
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func errorHandlingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
                log.Printf("panic: %v", err)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

func main() {
    // Load the configuration
    cfg, err := config.LoadConfig("configs/config.yaml")
    if err != nil {
        log.Fatalf("could not load config: %s\n", err)
    }

    // Create a new router
    r := mux.NewRouter()

    // Define routes and associate them with handler functions
    r.HandleFunc("/", handler.HomeHandler).Methods(http.MethodGet)
    r.HandleFunc("/about", handler.AboutHandler).Methods(http.MethodGet)
    r.NotFoundHandler = http.HandlerFunc(handler.NotFoundHandler)

    // Apply middleware
    r.Use(loggingMiddleware)
    r.Use(authenticationMiddleware)
    r.Use(errorHandlingMiddleware)

    // Start the server
    log.Printf("Starting server on %s\n", cfg.Server.Port)
    if err := http.ListenAndServe(cfg.Server.Port, r); err != nil {
        log.Fatalf("could not start server: %s\n", err)
    }
}
