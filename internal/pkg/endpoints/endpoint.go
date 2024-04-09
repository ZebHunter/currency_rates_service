package endpoints

import (
	"encoding/json"
	"fmt"
	"hw1/internal/pkg/model"
	"net/http"
)

type ServiceImpl interface {
	GetInfo() (model.InfoRes, error)
	GetCurrency(date string, name string) (model.CurrencyRes, error)
}

type Endpoint struct {
	serv ServiceImpl
}

func NewEndpoint(serv ServiceImpl) *Endpoint {
	return &Endpoint{serv}
}

func (e *Endpoint) CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")
	name := r.URL.Query().Get("currency")
	currency, err := e.serv.GetCurrency(date, name)
	fmt.Println(currency)
	jsonValue, err := json.Marshal(currency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonValue)
}

func (e *Endpoint) InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	response, _ := e.serv.GetInfo()
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
}
