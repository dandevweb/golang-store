package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "Camiseta", Description: "Azul, bem bonita", Price: 39.90, Quantity: 5},
		{"Tênis", "Confortável", 89, 3},
		{"Fone de Ouvido", "Muito bom", 59, 2},
		{"Produto Novo", "Descrição de teste", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
