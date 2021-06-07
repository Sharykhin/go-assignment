package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Sharykhin/go-assignment/server/response"
	"github.com/Sharykhin/go-assignment/switcher"
)

type (
	// HealthCheckResponse is a http response for health check route
	CalculateResponse struct {
		Result string `json:"result"`
	}
)

var (
	logicSwitcher = switcher.NewLogical()
)

// CalculateHandler checks whether server works or nor
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseBool(r.FormValue("a"))
	if err != nil {
		response.Fial(w, http.StatusBadRequest, err)
		return
	}

	b, err := strconv.ParseBool(r.FormValue("b"))
	if err != nil {
		response.Fial(w, http.StatusBadRequest, err)
		return
	}

	c, err := strconv.ParseBool(r.FormValue("c"))
	if err != nil {
		response.Fial(w, http.StatusBadRequest, err)
		return
	}

	d, err := strconv.ParseFloat(r.FormValue("d"), 64)
	if err != nil {
		response.Fial(w, http.StatusBadRequest, err)
		return
	}

	e, err := strconv.Atoi(r.FormValue("e"))
	if err != nil {
		response.Fial(w, http.StatusBadRequest, err)
		return
	}

	f, err := strconv.Atoi(r.FormValue("f"))
	if err != nil {
		response.Fial(w, http.StatusBadRequest, err)
		return
	}

	mode := switcher.Mode(r.FormValue("mode"))
	if mode == "" {
		mode = switcher.Base
	}

	k, err := logicSwitcher.Calculate(a, b, c, d, e, f, mode)
	if err != nil {
		if errors.Is(err, switcher.ErrUnexpectedInput) {
			response.Fial(w, http.StatusBadRequest, errors.New("unexpected input for a, b, c"))
			return
		}
		if errors.Is(err, switcher.ErrUnexpectedMode) {
			response.Fial(w, http.StatusBadRequest, errors.New("mode must be either of base, custom1 or custom2"))
			return
		}
		response.Fial(w, http.StatusInternalServerError, err)
		return
	}

	res := CalculateResponse{
		Result: fmt.Sprintf("%.2f", k),
	}

	response.Success(w, &res)
}