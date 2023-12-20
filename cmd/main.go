package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/spotich/routes/pkg/dbmanager"
)

const templatesDir = "frontend/templates"

type order struct {
	Address        string
	DeliveryWindow string
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/logist", sayHelloLogist)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func sayHelloLogist(w http.ResponseWriter, r *http.Request) {
	orders := []order{
		{"г. Новосибирск, улица 1905 года, дом 73, квартира 21", "9:00 - 13:00"},
		{"г. Новосибирск, улица Гоголя, дом 4, квартира 31", "14:00 - 18:00"},
		{"г. Новосибирск, улица Сибирская, дом 20, кв. 69", "18:00 - 22:00"},
	}
	data := map[string]interface{}{
		"Title":  "Logist Page",
		"Orders": orders,
	}
	var templates = template.Must(template.ParseGlob("../" + templatesDir + "/*.html"))
	err := templates.ExecuteTemplate(w, "logist.html", data)
	if err != nil {
		panic(err)
	}
}
