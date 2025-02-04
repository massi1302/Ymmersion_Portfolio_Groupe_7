package main

import (
	temp "Ymmersion/assets/Templates"
	"Ymmersion/src/routes"
)

func main() {
	temp.InitTemplates()
	routes.InitServe()
}
