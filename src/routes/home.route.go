package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", controllers.HomePage)
}
