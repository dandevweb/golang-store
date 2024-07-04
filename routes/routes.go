package routes

import (
	"net/http"
	"store/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
