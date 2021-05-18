package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dashinja/working/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":3000", sm)
}
