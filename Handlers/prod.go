package Handlers

import (
	"encoding/json"
	"log"
	"net/http"

	data "example.com/bharath/rest/Data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "unable to marshall json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
