package routes

import (
	"fmt"
	"net/http"
)

func InitServe() {

	homeRoutes()
	ProjetsRoutes()
	tableauRoutes()
	Challenge2Routes()
	fmt.Println("Le serveur est opérationel : http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}
