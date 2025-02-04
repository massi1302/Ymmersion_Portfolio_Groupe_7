package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func projetsRoutes() {
	http.HandleFunc("/projets", controllers.ProjetsPage)
}
