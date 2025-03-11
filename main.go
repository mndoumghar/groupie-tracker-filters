package main

import (
	"fmt"
	"log"
	"net/http"

	"g/function"
)

func main() {
	http.HandleFunc("/static/", function.Static)
	http.HandleFunc("/", function.HomePage)
	http.HandleFunc("/artist", function.ArtistPage)
	fmt.Println("Server running at http://localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}
