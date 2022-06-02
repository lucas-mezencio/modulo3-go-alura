package routes

import (
	"net/http"
	"web-crud/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
