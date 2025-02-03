package controllers

import (
	temp "Ymmersion/Templates"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	temp.Temp.ExecuteTemplate(w, "index", nil)
}
