package api

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Error   error
	Message string
	Code    int
}

func NewServerError(err error) *httpError {
	return &httpError{Message: "Server Error", Code: http.StatusInternalServerError}
}

type AppHandler func(http.ResponseWriter, *http.Request) *httpError

func (h AppHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if err := h(rw, r); err != nil {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(err.Code)
		json.NewEncoder(rw).Encode(map[string]string{
			"message": err.Message,
		})
	}
}
