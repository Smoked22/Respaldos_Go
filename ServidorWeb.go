package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Saludos, bienvenido a %s!", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// url para acceder al servidor de prueba: http://localhost:8080/url_prueba
