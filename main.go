package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: " + r.URL.Path)
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 30000...")
	http.ListenAndServe(":3000", nil)
}
