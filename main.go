package main

import (
	temp "Ymmersion/assets/Templates"
	_ "Ymmersion/src/Controllers"
	"Ymmersion/src/routes"
)

func main() {
	temp.InitTemplates()
	routes.InitServe()
}
