package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	PetHandler *PetHandler
}

func setupResponse(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	setupResponse(w, r)
	if r.Method == "OPTIONS" {
		return
	}

	if strings.HasPrefix(r.URL.Path, "/pets") {
		h.PetHandler.ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func EncodeJSON(w http.ResponseWriter, v interface{}, logger *log.Logger) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		Error(w, err, http.StatusInternalServerError, logger)
		return
	}
}

func Error(w http.ResponseWriter, err error, code int, logger *log.Logger) {
	// Log error.
	logger.Printf("http error: %s (code=%d)", err, code)

	// Hide error from client if it is internal.
	if code == http.StatusInternalServerError {
	}

	// Write generic error response.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&errorResponse{Err: err.Error()})
}

type errorResponse struct {
	Err string `json:"err,omitempty"`
}
