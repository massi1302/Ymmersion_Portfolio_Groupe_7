package controllers

import (
	"Ymmersion/templates"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	temp.Temp.ExecuteTemplate(w, "home", nil)
}