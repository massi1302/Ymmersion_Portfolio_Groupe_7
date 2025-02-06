package controllers

import (
	temp "Ymmersion/assets/Templates"
	"net/http"
	"strings"
)

type TextAnalysisResult struct {
	Text        string
	WordCount   int
	CharCount   int
	LongestWord string
	Error       string
}

func TextAnalyzerPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		text := r.Form.Get("text")
		result := analyzeText(text)

		temp.Temp.ExecuteTemplate(w, "text-analyzer", result)
		return
	}

	// GET request, render empty form
	temp.Temp.ExecuteTemplate(w, "text-analyzer", TextAnalysisResult{})
}

func analyzeText(text string) TextAnalysisResult {
	// Supprimer les espaces superflus
	text = strings.TrimSpace(text)

	// Vérification du texte vide
	if text == "" {
		return TextAnalysisResult{Error: "Veuillez entrer un texte à analyser"}
	}

	// Découper en mots
	words := strings.Fields(text)
	wordCount := len(words)
	charCount := len(strings.ReplaceAll(text, " ", "")) // Ne pas compter les espaces

	// Trouver le mot le plus long
	longestWord := ""
	for _, word := range words {
		if len(word) > len(longestWord) {
			longestWord = word
		}
	}

	return TextAnalysisResult{
		Text:        text,
		WordCount:   wordCount,
		CharCount:   charCount,
		LongestWord: longestWord,
	}
}