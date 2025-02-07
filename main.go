package main

import (
	temp "Ymmersion/assets/Templates"
	_ "Ymmersion/src/Controllers"
	"Ymmersion/src/routes"
)

//main.go
func main() {
	// Initialise les templates HTML pour le rendu des pages
	temp.InitTemplates()
	   // Démarre le serveur web et configure les routes
	routes.InitServe()
}
