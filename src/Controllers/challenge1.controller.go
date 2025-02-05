package controllers

import (
	temp "Ymmersion/assets/Templates"
	"net/http"
)

func Challenge1Page(w http.ResponseWriter, r *http.Request) {
	
	temp.Temp.ExecuteTemplate(w, "challenge1", nil)
}