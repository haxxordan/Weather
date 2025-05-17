package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/weather/{location}", GetWeatherHandler).Methods("GET")

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
    // Handler logic will be implemented later
    w.WriteHeader(http.StatusNotImplemented)
}