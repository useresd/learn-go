package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/petradevsteam/sms/internal/account"
	"github.com/petradevsteam/sms/internal/storage"
)

type AccountHandler struct {
	s *account.Service
	l *log.Logger
}

func NewAccountHandler(s *account.Service, l *log.Logger) *AccountHandler {
	return &AccountHandler{s: s, l: l}
}

func (h *AccountHandler) GetByID(rw http.ResponseWriter, r *http.Request) *httpError {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		h.l.Fatal(err)
	}

	account, err := h.s.GetByID(id)

	if err != nil {

		h.l.Println(err.Error())

		// Check if not found!
		if err, ok := err.(*storage.ErrorRecordNotFound); ok {
			return &httpError{Message: err.Error(), Code: http.StatusNotFound}
		}

		return NewServerError(err)
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(account)
	return nil
}

func (a *AccountHandler) Index(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Now handling")
	accounts, err := a.s.Get(r.Context(), 10, 0, "")

	if err != nil {
		a.handleError(rw, err, http.StatusNotFound)
		return
	}
	a.l.Println(accounts)

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(accounts)
}

func (a *AccountHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Validate id
	if params["id"] == "" {
		a.handleError(rw, errors.New("id is required"), 400)
		return
	}

	// Convert params id to int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		a.handleError(rw, errors.New("given id is not valid"), 400)
	}

	// Delete the account
	err = a.s.DeleteByID(r.Context(), id)

	if err != nil {
		a.handleError(rw, err, 500)
		return
	}

	a.toJson(rw, http.StatusOK, map[string]interface{}{
		"message": "Account Deleted!",
	})

}

func (a *AccountHandler) handleError(rw http.ResponseWriter, err error, code int) {
	a.l.Println(err)
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(map[string]string{
		"message": err.Error(),
	})
}

func (a *AccountHandler) toJson(rw http.ResponseWriter, code int, res map[string]interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	json.NewEncoder(rw).Encode(res)
}
