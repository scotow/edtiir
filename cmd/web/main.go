package main

import (
	"github.com/scotow/edtiir"
	"log"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	year := edtiir.Year2018
	http.Redirect(w, r, year.CurrentWeek().Link(), http.StatusTemporaryRedirect)
}

func listeningAddress() string {
	port, set := os.LookupEnv("PORT")
	if !set {
		port = "8080"
	}

	return ":" + port
}

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(listeningAddress(), nil))
}
