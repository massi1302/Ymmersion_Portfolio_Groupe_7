package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

// ProjetsRoutes doit être exportée (majuscule)
func ProjetsRoutes() {
	http.HandleFunc("/index", controllers.HomePage)
	http.HandleFunc("/projets", controllers.ProjetsPage)
	http.HandleFunc("/guess-number", controllers.GuessNumberGameHandler)
}

// Handler séparé pour le jeu
func GuessNumberHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet, http.MethodPost:
		controllers.GuessNumberGameHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
