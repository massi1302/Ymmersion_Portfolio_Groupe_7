package controllers

import (
	epreuve "Ymmersion/epreuve"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RenderTemplate(w http.ResponseWriter, tmplName string, data interface{}) error {
	tmpl, err := template.ParseFiles("Ymmersion/assets/Templates/guess-number.html")
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, "guess-number", data)
}

func GuessNumberGameHandler(w http.ResponseWriter, r *http.Request, ctx interface{}) error {
	// Adapter method to use the existing Fiber controller with standard library
	guessStr := ctx.(struct{ FormValue func(string) string }).FormValue("guess")
	// Convert guessStr to an integer
	guessInt, err := strconv.Atoi(guessStr)
	if err != nil {
		http.Error(w, "Invalid guess number", http.StatusBadRequest)
		return err
	}

	// Use existing PlayGame logic from epreuve package
	result := epreuve.PlayGame(guessInt)

	// Render template with game result
	return RenderTemplate(w, "guess-number", fiber.Map{
		"message":      result.Message,
		"targetNumber": result.TargetNumber,
		"guessNumber":  result.GuessNumber,
		"attempts":     result.Attempts,
		"isCorrect":    result.IsCorrect,
	})
}
