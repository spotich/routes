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
	Lat            float32
	Lon            float32
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
		{"улица 1905 года, дом 73, квартира 21", "9:00 - 13:00", 55.044722, 82.908705},
		{"улица Гоголя, дом 4, квартира 31", "14:00 - 18:00", 55.041667, 82.914941},
		{"улица Писарева, дом 36, кв. 69", "18:00 - 22:00", 55.048253, 82.922511},
		{"улица 1905 года, дом 73, квартира 21", "9:00 - 13:00", 55.044722, 82.908705},
		{"улица Гоголя, дом 4, квартира 31", "14:00 - 18:00", 55.041667, 82.914941},
		{"улица Писарева, дом 36, кв. 69", "18:00 - 22:00", 55.048253, 82.922511},
		{"улица 1905 года, дом 73, квартира 21", "9:00 - 13:00", 55.044722, 82.908705},
		{"улица Гоголя, дом 4, квартира 31", "14:00 - 18:00", 55.041667, 82.914941},
		{"улица Писарева, дом 36, кв. 69", "18:00 - 22:00", 55.048253, 82.922511},
	}
	data := map[string]interface{}{
		"Title":        "Logist Page",
		"Name":         "Юрий",
		"TodayDate":    "21 декабря 2023 года",
		"TomorrowDate": "22 декабря 2023 года",
		"Orders":       orders,
	}
	var templates = template.Must(template.ParseGlob("../" + templatesDir + "/*.html"))
	err := templates.ExecuteTemplate(w, "logist.html", data)
	if err != nil {
		panic(err)
	}
}
