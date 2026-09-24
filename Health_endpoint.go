package main

import (
	"encoding/json"
	"log"
	"net/http"
)


type HealthResponse struct {
	Status string `json:"status"`
}


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func healthHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow GET
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	// response making
	res := HealthResponse{
		Status: "ok",
	}

	// header + status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// JSON send
	json.NewEncoder(w).Encode(res)
}


func loginHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest

	// JSON body read
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	
	log.Println("Username:", req.Username)
	log.Println("Password:", req.Password)

	// simple response
	w.Write([]byte("Login successful"))
}


func main() {

	mux := http.NewServeMux()

	
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/login", loginHandler)

	log.Println("Server running at http://localhost:8080")

	// server start
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}