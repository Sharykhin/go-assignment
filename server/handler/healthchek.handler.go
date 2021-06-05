package handler

import (
	"net/http"

	"github.com/Sharykhin/go-assignment/server/response"
)

type (
	// HealthCheckResponse is a http response for health check route
	HealthCheckResponse struct {
		Status string `json:"status"`
	}
)

// HealthCheck checks whether server works or nor
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := HealthCheckResponse{
		Status: "Server is running",
	}

	response.Success(w, &res)
}