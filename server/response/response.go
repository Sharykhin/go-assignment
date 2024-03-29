package response

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Sharykhin/go-assignment/logger"
)

type (
	errorResponse struct {
		Message string `json:"message"`
	}
)

// Success returns a success response with status code 200
func Success(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		log.Printf("failed to encode http response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Internal Server Error"))
		if err != nil {
			log.Printf("failed to write http response: %v", err)
		}
	}

}

// Fail returns error response and also logs an error
func Fail(w http.ResponseWriter, status int, err error) {
	logger.Log.Error(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errRes := errorResponse{
		Message: err.Error(),
	}

	err = json.NewEncoder(w).Encode(&errRes)
	if err != nil {
		log.Printf("failed to encode http response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Internal Server Error"))
		if err != nil {
			log.Printf("failed to write http response: %v", err)
		}
	}
}
