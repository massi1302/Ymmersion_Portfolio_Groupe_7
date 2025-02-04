package controllers

import (
	temp "Ymmersion/assets/Templates"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	temp.Temp.ExecuteTemplate(w, "index", nil)
}
