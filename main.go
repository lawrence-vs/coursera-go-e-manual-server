package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
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

func newsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved: " + r.URL.Path)
	t, err := template.ParseFiles("templates/news.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%v Server error\n", http.StatusNotFound)
		fmt.Fprintf(w, "Description: %s\n", err)
		return
	}
	// get today's 2025-07-03
	date := time.Now().String()
	t.Execute(w, date)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/news", newsHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/manuals/", http.StripPrefix("/manuals/", http.FileServer(http.Dir("manuals"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))
	fmt.Println("Listening on port 30000...")
	http.ListenAndServe(":3000", nil)
}
