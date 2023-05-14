package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/about", handleAbout)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello, world!")
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Page about!")
}
