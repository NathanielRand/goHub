package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/NathanielRand/goHub/views"
)

var (
	homeView *views.View
	newsView *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing homeView template: %v\n", err)
	}
}

func news(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := newsView.Template.ExecuteTemplate(w, newsView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing homeView template: %v\n", err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ</h1>\n<h2>Coming Soon!</h2>")
}

func main() {
	// Call NewView func from views package
	homeView = views.NewView("materialize", "views/home.gohtml")
	newsView = views.NewView("materialize", "views/news.gohtml")
	contactView = views.NewView("materialize", "views/contact.gohtml")

	// Create router
	r := mux.NewRouter()
	
	// Asset Routes
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	// Define paths with assigned functions
	r.HandleFunc("/", home)
	r.HandleFunc("/news", news)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)

	// Start up web server and
	// pass in router as deafult handler for web requests
	http.ListenAndServe(":8080", r)
}
