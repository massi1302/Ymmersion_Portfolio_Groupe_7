package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func PasswordGeneratorRoutes() {
	http.HandleFunc("/password-generator", controllers.PasswordGeneratorPage)
}
