package main

import (
	"encoding/json"
	"log"
	"net/http"
)


// Health Response
type HealthResponse struct {
	Status string `json:"status"`
}


// Login Request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


// Login Response
type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}


func healthHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	res := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}


func loginHandler(w http.ResponseWriter, r *http.Request) {

	//  Only POST allowed
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	//  Parse JSON
	var req LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//  Check credentials (hardcoded)
	if req.Username != "admin" || req.Password != "password123" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Send JSON response
	res := LoginResponse{
		Status:  "ok",
		Message: "login successful",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}


func main() {

	mux := http.NewServeMux()

	
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/login", loginHandler)

	log.Println("Server running at http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}