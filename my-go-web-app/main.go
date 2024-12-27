package main

import (
	"net/http"
	"html/template"
)

var tmpl = template.Must(template.ParseGlob("static/*.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "products.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "about.html", nil)
}

func main() {
	// Serve static files (CSS, JS, images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define routes for the pages
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the server
	port := ":8080"
	println("Server is running on http://localhost" + port)
	http.ListenAndServe(port, nil)
}

