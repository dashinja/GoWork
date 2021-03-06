package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request)  {
	h.l.Println("Hello World!")

	data, err := ioutil.ReadAll((req.Body))
	if err != nil {
		http.Error(res, "Real oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(res, "Hello %s", data)
}
