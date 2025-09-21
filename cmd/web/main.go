package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	data := GatherData()

	templateDir := "templates"
	outputDir := "github-pages"

	indexPage := filepath.Join(templateDir, "index.html")
	aboutMeTemplate := filepath.Join(templateDir, "about-me.html")
	projectTemplate := filepath.Join(templateDir, "project.html")

	tmpl, err := template.ParseFiles(indexPage, aboutMeTemplate, projectTemplate)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	f, err := os.OpenFile(filepath.Join(outputDir, "index.html"), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Error opening html file: %v", err)
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Printf("Error rendering template: %v", err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
