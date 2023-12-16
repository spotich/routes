package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/spotich/routes/pkg/dbmanager/main.go"
)

func main() {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
