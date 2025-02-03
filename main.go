package main

import (
	temp "Ymmersion/Templates"
	"Ymmersion/src/routes"
)

func main() {
	temp.InitTemplates()
	routes.InitServe()
}
