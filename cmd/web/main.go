package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	data := GatherData()

	fs := http.FileServer(http.Dir("ui/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tmpl, err := template.ParseFiles("ui/html/pages/index.html", "ui/html/partials/project-template.html", "ui/html/partials/about-me.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	f, err := os.OpenFile("static-site/index.html", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening html file: %v", err)
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
	}
}
