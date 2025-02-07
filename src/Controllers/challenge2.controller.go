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

// Gère le jeu "Devinez le nombre"
func Challenge2Page(w http.ResponseWriter, r *http.Request) {
	if temp.Temp == nil {
		http.Error(w, "Erreur : Template non chargé", http.StatusInternalServerError)
		return
	}

	 // Génère un nombre aléatoire entre 1 et 99
	message := ""
	if r.Method == http.MethodPost {
		r.ParseForm()
		guess, err := strconv.Atoi(r.FormValue("guess"))

		if err != nil {
			message = "Veuillez entrer un nombre valide."
		} else {
			 // Compare la réponse du joueur avec le nombre à deviner
			mutex.Lock()
			defer mutex.Unlock()
			  // Retourne des indices "plus grand" ou "plus petit"
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
