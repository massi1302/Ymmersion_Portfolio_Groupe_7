package controllers

import (
	temp "Ymmersion/Templates"
	"net/http"
)

func ProjetsPage(w http.ResponseWriter, r *http.Request) {

	temp.Temp.ExecuteTemplate(w, "projets", nil)
}
