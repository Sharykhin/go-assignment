package handler

import (
	"net/http"

	"github.com/Sharykhin/go-assignment/server/response"
)

type (
	// HealthCheckResponse is a http response for health check route
	CalculateResponse struct {
		Result float64 `json:"result"`
	}
)

// CalculateHandler checks whether server works or nor
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	res := CalculateResponse{
		Result: 0.1 + 0.2,
	}

	response.Success(w, &res)
}