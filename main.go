package main

import (
	"log"
	"net/http"
	"student/ascii-art-web/servers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", servers.MainPage)
	mux.HandleFunc("/ascii-art", servers.AsciiPage)
	log.Println("server : http:localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
