package handlers

import (
	"log"
	"net/http"

	"github.com/dashinja/working/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProduct()

	err := lp.ToJSON(rw)

	if err != nil {
		log.Fatal(err)
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
