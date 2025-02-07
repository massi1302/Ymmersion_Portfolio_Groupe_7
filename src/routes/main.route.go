package routes

import (
	"fmt"
	"net/http"
)


func InitServe() {
  // Initialise toutes les routes de l'application
	homeRoutes()
	ProjetsRoutes()
	tableauRoutes()
	Challenge2Routes()
	Challenge1Routes()
	PasswordGeneratorRoutes()
	PalindromeRoutes()
	PasswordValidatorRoutes()
	TextAnalyzerRoutes()
	// Affiche un message de confirmation dans la console
	fmt.Println("Le serveur est opérationel : http://localhost:8080")
	// Configure le serveur sur le port 8080  
	http.ListenAndServe("localhost:8080", nil)
}
