package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func Challenge2Routes() {
	http.HandleFunc("/challenge2", controllers.Challenge2Page)
}
