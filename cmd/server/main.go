package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petradevsteam/sms/internal/account"
	"github.com/petradevsteam/sms/internal/api"
	"github.com/petradevsteam/sms/internal/storage"
)

func main() {
	mysql, err := storage.NewMySQL()
	if err != nil {
		log.Fatal(err)
	}

	defer mysql.DB().Close()

	l := log.Default()

	r := mux.NewRouter()

	account := api.NewAccountHandler(account.NewService(account.NewStorage(mysql.DB())), l)
	r.Handle("/account/{id}", api.AppHandler(account.GetByID)).Methods(http.MethodGet)
	r.HandleFunc("/account", account.Index).Methods(http.MethodGet)
	r.HandleFunc("/account/{id}", account.Delete).Methods(http.MethodDelete)

	http.ListenAndServe(":8090", r)
}
