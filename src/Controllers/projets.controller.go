package controllers

import (
	temp "Ymmersion/assets/Templates"
	"net/http"
)

func ProjetsPage(w http.ResponseWriter, r *http.Request) {

	temp.Temp.ExecuteTemplate(w, "projets", nil)
}
