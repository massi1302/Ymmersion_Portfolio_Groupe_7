package routes

import (
	controllers "Ymmersion/src/Controllers"
	"net/http"
)

func tableauRoutes() {
	http.HandleFunc("/tableau", controllers.TableauPage)
}
