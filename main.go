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
	sitesView *views.View
	coursesView *views.View
	mediaView *views.View
	jobsView *views.View
	merchandiseView *views.View
	groupsView *views.View
	faqView *views.View
	contactView *views.View
	aboutView *views.View
	feedbackView *views.View
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

func sites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := sitesView.Template.ExecuteTemplate(w, sitesView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func courses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := coursesView.Template.ExecuteTemplate(w, coursesView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func media(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := mediaView.Template.ExecuteTemplate(w, mediaView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func jobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := jobsView.Template.ExecuteTemplate(w, jobsView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func merchandise(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := merchandiseView.Template.ExecuteTemplate(w, merchandiseView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func groups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := groupsView.Template.ExecuteTemplate(w, groupsView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := aboutView.Template.ExecuteTemplate(w, aboutView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func feedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := feedbackView.Template.ExecuteTemplate(w, feedbackView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := faqView.Template.ExecuteTemplate(w, faqView.Layout, nil)
	if err != nil {
		fmt.Printf("Error executing contactView template: %v\n", err)
	}
}

func main() {
	// Call NewView func from views package
	homeView = views.NewView("materialize", "views/home.gohtml")
	newsView = views.NewView("materialize", "views/news.gohtml")
	sitesView = views.NewView("materialize", "views/sites.gohtml")
	coursesView = views.NewView("materialize", "views/courses.gohtml")
	mediaView = views.NewView("materialize", "views/media.gohtml")
	jobsView = views.NewView("materialize", "views/jobs.gohtml")
	merchandiseView = views.NewView("materialize", "views/merchandise.gohtml")
	groupsView = views.NewView("materialize", "views/groups.gohtml")
	aboutView = views.NewView("materialize", "views/about.gohtml")
	contactView = views.NewView("materialize", "views/contact.gohtml")
	feedbackView = views.NewView("materialize", "views/feedback.gohtml")
	faqView = views.NewView("materialize", "views/faq.gohtml")

	// Create router
	r := mux.NewRouter()
	
	// Asset Routes
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)

	// Define paths with assigned functions
	r.HandleFunc("/", home)
	r.HandleFunc("/news", news)
	r.HandleFunc("/sites", sites)
	r.HandleFunc("/courses", courses)
	r.HandleFunc("/media", media)
	r.HandleFunc("/jobs", jobs)
	r.HandleFunc("/merchandise", merchandise)
	r.HandleFunc("/groups", groups)
	r.HandleFunc("/about", about)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/feedback", feedback)
	r.HandleFunc("/faq", faq)

	// Start up web server and
	// pass in router as deafult handler for web requests
	http.ListenAndServe(":8080", r)
}
