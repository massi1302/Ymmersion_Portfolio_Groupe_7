package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func homeRoutes() {
	http.HandleFunc("/", controllers.HomePage)
}

func ProjetsRoutes() {
	http.HandleFunc("/index", controllers.HomePage)
	http.HandleFunc("/projets", controllers.ProjetsPage)
}

func Challenge1Routes() {
	http.HandleFunc("/challenge1", controllers.Challenge1Page)
}
func Challenge2Routes() {
	http.HandleFunc("/challenge2", controllers.Challenge2Page)
}

func PasswordGeneratorRoutes() {
	http.HandleFunc("/password-generator", controllers.PasswordGeneratorPage)
}

func tableauRoutes() {
	http.HandleFunc("/tableau", controllers.TableauPage)
}

func PalindromeRoutes() {
	http.HandleFunc("/palindrome", controllers.PalindromePage)
}

