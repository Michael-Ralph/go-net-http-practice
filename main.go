package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	// Serve static files (css, images, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	// Start the server
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "home.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
	}{
		Title: "About Us",
	}
	renderTemplate(w, "about.html", data)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Process form submission
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		// For demo purposes, just print the form data
		fmt.Fprintf(w, "Recieved contact form:\nName: %s\nEmail: %s\nMessage: %s", name, email, message)
		return
	}

	// Show contact form
	renderTemplate(w, "contact.html", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println(err)
	}
}
