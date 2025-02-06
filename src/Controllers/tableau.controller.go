package controllers

import (
	temp "Ymmersion/assets/Templates"
	"net/http"
)

func TableauPage(w http.ResponseWriter, r *http.Request) {

	temp.Temp.ExecuteTemplate(w, "tableau", nil)
}
