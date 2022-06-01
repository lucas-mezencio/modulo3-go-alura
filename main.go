package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	_ = http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{
			Nome:       "camiseta",
			Descricao:  "Azul escuro",
			Preco:      39,
			Quantidade: 10,
		},
		{
			"Tenis", "Comfortável", 200, 5,
		},
		{
			"Fone", "pure bass", 150, 10,
		},
	}

	_ = templates.ExecuteTemplate(w, "Index", produtos)
}
