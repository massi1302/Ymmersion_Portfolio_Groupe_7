package main

import (
	"net/http"
	"os"
)

// creation d'un simple serveur web
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
