package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func ProjetsRoutes() {
	http.HandleFunc("/index", controllers.HomePage)
	http.HandleFunc("/projets", controllers.ProjetsPage)
}
