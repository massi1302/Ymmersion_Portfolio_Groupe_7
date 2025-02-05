package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func Challenge1Routes() {
	http.HandleFunc("/challenge1", controllers.Challenge1Page)
}