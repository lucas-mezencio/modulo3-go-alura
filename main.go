package main

import (
	"net/http"
	"web-crud/routes"
)

func main() {

	routes.CarregaRotas()
	_ = http.ListenAndServe(":8000", nil)
}
