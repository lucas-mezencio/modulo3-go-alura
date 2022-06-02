package controllers

import (
	"html/template"
	"net/http"
	"web-crud/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	_ = templates.ExecuteTemplate(w, "Index", produtos)
}
