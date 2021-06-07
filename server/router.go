package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Sharykhin/go-assignment/server/handler"
)

func router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/v1/calculate", handler.CalculateHandler)

	return r
}
