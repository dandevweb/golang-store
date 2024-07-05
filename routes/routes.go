package routes

import (
	"net/http"
	"store/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/store", controllers.Store)
	http.HandleFunc("/delete", controllers.Delete)
}
