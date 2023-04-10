package main

import (
	"fmt"
	"net/http"
)

func main() {
	request()
}

func request() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/aboutme", aboutMe)

	http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to the Fo web API")
	fmt.Println("Homepage end point")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	who := "Micheal"
	fmt.Fprintf(w, "A little about me")
	fmt.Println("aboutME end point", who)
}
