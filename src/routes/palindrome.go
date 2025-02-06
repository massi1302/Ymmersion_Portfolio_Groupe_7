package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func PalindromeRoutes() {
	http.HandleFunc("/palindrome", controllers.PalindromePage)
}
