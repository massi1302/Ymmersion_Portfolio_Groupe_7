package main

import (
	"Ymmersion/src/routes"
	temp "Ymmersion/templates"
)

func main() {
	temp.InitTemplates()
	routes.InitServe()
}
