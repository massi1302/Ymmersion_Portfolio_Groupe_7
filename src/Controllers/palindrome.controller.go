package controllers

import (
	temp "Ymmersion/assets/Templates"
	"net/http"
	"strings"
)

type PalindromeResult struct {
	Input   string
	Message string
	IsValid bool
}

func PalindromePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp.Temp.ExecuteTemplate(w, "palindrome", nil)
	} else if r.Method == http.MethodPost {
		word := r.FormValue("word")
		result := checkPalindrome(word)
		temp.Temp.ExecuteTemplate(w, "palindrome", result)
	}
}

func checkPalindrome(word string) PalindromeResult {
	if word == "" {
		return PalindromeResult{
			Message: "Veuillez entrer un mot ou une phrase.",
			IsValid: false,
		}
	}

	// Nettoyer le mot (enlever les espaces et convertir en minuscules)
	cleaned := strings.ToLower(strings.ReplaceAll(word, " ", ""))

	// Vérifier si c'est un palindrome
	length := len(cleaned)
	for i := 0; i < length/2; i++ {
		if cleaned[i] != cleaned[length-1-i] {
			return PalindromeResult{
				Input:   word,
				Message: "Ce n'est pas un palindrome.",
				IsValid: false,
			}
		}
	}

	return PalindromeResult{
		Input:   word,
		Message: "C'est un palindrome !",
		IsValid: true,
	}
}
