package controllers

import (
	"math/rand"
	"net/http"
	"strconv"
	temp "Ymmersion/assets/Templates"
	"time"
)

type PasswordGeneratorResult struct {
	Password string
	Error    string
}

func PasswordGeneratorPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp.Temp.ExecuteTemplate(w, "password-generator", nil)
	} else if r.Method == http.MethodPost {
		lengthStr := r.FormValue("length")
		length, err := strconv.Atoi(lengthStr)
		
		result := generatePassword(length, err)
		temp.Temp.ExecuteTemplate(w, "password-generator", result)
	}
}

func generatePassword(length int, err error) PasswordGeneratorResult {
	 // Valide que la longueur est entre 8 et 16 caractères
	if err != nil || length < 8 || length > 16 {
		return PasswordGeneratorResult{
			Error: "Password length must be between 8 and 16 characters.",
		}
	}

	rand.Seed(time.Now().UnixNano())
	password := ""
	 // Génère un mot de passe aléatoire selon la longueur spécifiée
	for i := 0; i < length; i++ {
		password += string(rune(rand.Intn(94) + 33))// Utilise des caractères ASCII entre 33 et 126 (caractères imprimables)
	}

	return PasswordGeneratorResult{
		Password: password,
	}
}