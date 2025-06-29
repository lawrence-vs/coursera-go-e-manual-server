package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request: " + r.URL.Path)
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server error\n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s\n", err)
		return
	}
	pages, _ := ScanDir("./manuals")
	fmt.Println(pages)

	t.Execute(w, pages)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port 30000...")
	http.ListenAndServe(":3000", nil)
}
