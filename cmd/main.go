package main

import (
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
	http.Redirect(w, r, "/logist", http.StatusSeeOther)
}

func sayHelloLogist(w http.ResponseWriter, r *http.Request) {
	orders := []order{
		{"улица 1905 года, дом 73, квартира 21", "9:00 - 13:00", 55.044722, 82.908705},
		{"улица Гоголя, дом 4, квартира 31", "14:00 - 18:00", 55.041667, 82.914941},
		{"улица Писарева, дом 36, квартира. 69", "18:00 - 22:00", 55.048253, 82.922511},
		{"улица Железнодорожная, дом 8, квартира 213", "9:00 - 13:00", 55.043802, 82.897636},
		{"улица Дмитрия Шамшурина, дом 27, квартира 2", "14:00 - 18:00", 55.032634, 82.89956},
		{"улица Ленина, дом 21, квартира. 66", "18:00 - 22:00", 55.029209, 82.904255},
		{"улица Революции, дом 10, квартира 444", "9:00 - 13:00", 55.026749, 82.910197},
		{"улица Урицкого, дом 6, квартира 17", "14:00 - 18:00", 55.025108, 82.915526},
		{"улица Каменская, дом 1, квартира 1", "18:00 - 22:00", 55.024974, 82.928038},
	}
	data := map[string]interface{}{
		"Title":        "Кабинет логиста",
		"Name":         "Юрий",
		"TodayDate":    "23 декабря 2023 года",
		"TomorrowDate": "24 декабря 2023 года",
		"Orders":       orders,
	}
	var templates = template.Must(template.ParseGlob("../" + templatesDir + "/*.html"))
	err := templates.ExecuteTemplate(w, "logist.html", data)
	if err != nil {
		panic(err)
	}
}
