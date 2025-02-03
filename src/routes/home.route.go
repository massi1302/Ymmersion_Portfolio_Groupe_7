package routes

import (
	"Ymmersion/src/controllers"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", controllers.HomePage)
}