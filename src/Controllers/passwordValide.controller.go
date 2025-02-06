package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	temp "Ymmersion/assets/Templates"
)

// Fonction pour évaluer la force du mot de passe
func checkPasswordStrength(password string) string {
	var hasUpper, hasLower, hasDigit, hasSpecial bool
	specialChars := "!@#$%^&*()-_=+[]{}<>?/"

	// Vérifier chaque caractère du mot de passe
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsRune(specialChars, char):
			hasSpecial = true
		}
	}

	// Évaluer la force
	length := len(password)
	switch {
	case length >= 12 && hasUpper && hasLower && hasDigit && hasSpecial:
		return "💪 Fort (Bien sécurisé) 🔒"
	case length >= 8 && ((hasUpper && hasLower && hasDigit) || (hasLower && hasDigit && hasSpecial)):
		return "⚠️ Moyen (Ajoutez des majuscules et des symboles)"
	default:
		return "❌ Faible (Trop court ou manque de diversité)"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🔍 Vérificateur de Force de Mot de Passe")
	fmt.Println("➡️  Entrez un mot de passe (ou tapez 'exit' pour quitter) :")

	for {
		fmt.Print("> ")
		scanner.Scan()
		password := scanner.Text()

		// Vérifier si l'utilisateur veut quitter
		if strings.ToLower(password) == "exit" {
			fmt.Println("👋 Merci d'avoir utilisé le vérificateur ! À bientôt !")
			break
		}

		// Vérifier la force du mot de passe
		strength := checkPasswordStrength(password)
		fmt.Printf("🔑 Sécurité : %s\n\n", strength)
	}
}