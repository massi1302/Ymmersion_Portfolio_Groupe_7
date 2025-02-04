package epreuve

import (
	"math/rand"
	"time"
)

type GameResult struct {
	TargetNumber int
	GuessNumber  int
	Attempts     int
	Message      string
	IsCorrect    bool
}

func PlayGame(guess int) GameResult {
	// Initialisation du seed une seule fois
	rand.Seed(time.Now().UnixNano())

	// Nombre à deviner (généré à chaque appel)
	targetNumber := rand.Intn(99) + 1

	// Compteur de tentatives
	attempts := 6

	// Résultat du jeu
	result := GameResult{
		TargetNumber: targetNumber,
		GuessNumber:  guess,
		Attempts:     attempts,
	}

	// Logique de comparaison
	if guess < targetNumber {
		result.Message = "Trop bas ! Le nombre est plus grand."
		result.IsCorrect = false
	} else if guess > targetNumber {
		result.Message = "Trop haut ! Le nombre est plus petit."
		result.IsCorrect = false
	} else {
		result.Message = "Bravo ! Vous avez trouvé le bon nombre !"
		result.IsCorrect = true
	}

	return result
}
