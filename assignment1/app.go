package main

import (
	"github.com/drone/routes"
	"log"
	"net/http"
)

func main() {
	mux := routes.New()

	mux.Get("/profile/:email", GetProfile)

	http.Handle("/", mux)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	email := params.Get(":email")
	w.Write([]byte(email + "'s Profile"))
}
