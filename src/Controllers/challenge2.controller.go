package controllers

import (
	temp "Ymmersion/assets/Templates"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	numberToGuess int
	mutex         sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
	numberToGuess = rand.Intn(99) + 1
}

func Challenge2Page(w http.ResponseWriter, r *http.Request) {
	if temp.Temp == nil {
		http.Error(w, "Erreur : Template non chargé", http.StatusInternalServerError)
		return
	}

	message := ""
	if r.Method == http.MethodPost {
		r.ParseForm()
		guess, err := strconv.Atoi(r.FormValue("guess"))

		if err != nil {
			message = "Veuillez entrer un nombre valide."
		} else {
			mutex.Lock()
			defer mutex.Unlock()
			if guess < numberToGuess {
				message = "C'est plus grand !"
			} else if guess > numberToGuess {
				message = "C'est plus petit !"
			} else {
				message = "Bravo, vous avez trouvé le bon nombre !"
				numberToGuess = rand.Intn(99) + 1 // Réinitialisation
			}
		}
	}

	temp.Temp.ExecuteTemplate(w, "challenge2", message)
}
