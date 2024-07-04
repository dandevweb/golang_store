package controllers

import (
	"html/template"
	"log"
	"net/http"
	"store/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.GetAll()
	temp.ExecuteTemplate(w, "Index", products)
}

func Create(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Create", nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Erro na conversão do preço")
		}

		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("Erro na conversão da quantidade")
		}

		models.Create(name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}
