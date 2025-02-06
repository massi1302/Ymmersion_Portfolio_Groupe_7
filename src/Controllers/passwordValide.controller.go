package controllers

import (
	"net/http"
	"strings"
	"unicode"
	temp "Ymmersion/assets/Templates"
)

type PasswordValidationResult struct {
	Success string
	Error   string
}

func PasswordValidatorPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Afficher la page de validation de mot de passe
		temp.Temp.ExecuteTemplate(w, "password-validation", nil)
	} else if r.Method == http.MethodPost {
		// Récupérer le mot de passe soumis
		password := r.FormValue("password")
		
		// Valider le mot de passe
		result := validatePassword(password)
		
		// Rendre le template avec le résultat
		temp.Temp.ExecuteTemplate(w, "password-validation", result)
	}
}

func validatePassword(password string) PasswordValidationResult {
	// Vérifier la longueur
	if len(password) < 8 {
		return PasswordValidationResult{
			Error: "Le mot de passe doit contenir au moins 8 caractères.",
		}
	}

	// Initialiser les drapeaux de validation
	var (
		hasUpper    bool
		hasLower    bool
		hasDigit    bool
		hasSpecial  bool
	)

	// Caractères spéciaux à vérifier
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

	// Vérifier si tous les critères sont remplis
	if !hasUpper {
		return PasswordValidationResult{
			Error: "Le mot de passe doit contenir au moins une majuscule.",
		}
	}

	if !hasLower {
		return PasswordValidationResult{
			Error: "Le mot de passe doit contenir au moins une minuscule.",
		}
	}

	if !hasDigit {
		return PasswordValidationResult{
			Error: "Le mot de passe doit contenir au moins un chiffre.",
		}
	}

	if !hasSpecial {
		return PasswordValidationResult{
			Error: "Le mot de passe doit contenir au moins un caractère spécial.",
		}
	}

	// Si tous les critères sont remplis, retourner un succès
	return PasswordValidationResult{
		Success: "✅ Votre mot de passe est sécurisé ! Tous les critères sont respectés.",
	}
}