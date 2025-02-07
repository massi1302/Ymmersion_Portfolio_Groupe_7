package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", controllers.HomePage) // Configure la route principale vers la page d'accueil
}

func ProjetsRoutes() {
	http.HandleFunc("/index", controllers.HomePage)// Configure les routes pour la section projets
	http.HandleFunc("/projets", controllers.ProjetsPage)// Inclut la page d'accueil et la page des projets
}

func Challenge1Routes() {
	http.HandleFunc("/challenge1", controllers.Challenge1Page) // Configure la route pour le premier challenge
}
func Challenge2Routes() {
	http.HandleFunc("/challenge2", controllers.Challenge2Page) // Configure la route pour le jeu "Devinez le nombre"
}

func PasswordGeneratorRoutes() {
	http.HandleFunc("/password-generator", controllers.PasswordGeneratorPage)// Configure la route pour le générateur de mot de passe
}

func tableauRoutes() {
	http.HandleFunc("/tableau", controllers.TableauPage) // Configure la route pour la page du tableau
}

func PalindromeRoutes() {
	http.HandleFunc("/palindrome", controllers.PalindromePage)// Configure la route pour le vérificateur de palindromes
}

func PasswordValidatorRoutes() {
	http.HandleFunc("/password-validator", controllers.PasswordValidatorPage)// Configure la route pour le validateur de mot de passe
}

func TextAnalyzerRoutes() {
	http.HandleFunc("/text-analyzer", controllers.TextAnalyzerPage)  // Configure la route pour l'analyseur de texte
}
