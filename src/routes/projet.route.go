package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func projetsRoutes() {
	http.HandleFunc("/index", controllers.HomePage)
	http.HandleFunc("/projets", controllers.ProjetsPage)
	http.HandleFunc("/guess-number", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Render the initial game page
			controllers.RenderTemplate(w, "guess-number", nil)
		} else if r.Method == http.MethodPost {
			// Handle form submission
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			// Create a minimal Fiber context-like structure
			fiberCtx := struct {
				FormValue func(string) string
			}{
				FormValue: func(key string) string {
					return r.FormValue(key)
				},
			}

			// Render the result using the existing game logic
			controllers.GuessNumberGameHandler(w, r, fiberCtx)
		}
	})
}
