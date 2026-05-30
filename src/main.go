package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type appStatus struct {
	Name      string `json:"name"`
	Exercise  string `json:"exercise"`
	Pipeline  string `json:"pipeline"`
	State     string `json:"state"`
	Timestamp string `json:"timestamp"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := http.NewServeMux()
	router.HandleFunc("/", statusHandler)
	router.HandleFunc("/health", healthHandler)

	log.Printf("Aplikacja testowa działa na porcie %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(appStatus{
		Name:      "PAwChO - zadanie drugie",
		Exercise:  "GitHub Actions oraz obraz Docker",
		Pipeline:  "GHCR publication after CVE scan with DockerHub registry cache",
		State:     "OK",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK\n"))
}
