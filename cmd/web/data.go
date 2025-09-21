package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"personal-vibesite/internal/github"
)

const GITHUB_USERNAME = "Matthew-Berthoud"
const PROJECT_NAMES = "projects.txt"

type PageData struct {
	Projects []github.Project
	AboutMe  template.HTML
	CSS      template.CSS
	JS       template.JS
}

func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func GatherData() *PageData {
	repos, err := ReadLines(PROJECT_NAMES)
	if err != nil {
		log.Fatalf("failed to read repos: %v", err)
	}

	gh := github.NewGithubConnection(GITHUB_USERNAME)

	projects, err := gh.GetProjects(repos)
	if err != nil {
		log.Fatalf("Error getting projects: %v", err)
	}

	aboutMe, err := gh.GetAboutMe()
	if err != nil {
		log.Fatalf("Error getting About Me: %v", err)
	}

	staticDir := filepath.Join("ui", "static")
	cssFile := filepath.Join(staticDir, "css", "style.css")
	jsFile := filepath.Join(staticDir, "js", "script.js")

	css, err := os.ReadFile(cssFile)
	if err != nil {
		log.Fatalf("No CSS files found: %v", err)
	}

	js, err := os.ReadFile(jsFile)
	if err != nil {
		log.Fatalf("No JS files found: %v", err)
	}

	return &PageData{
		Projects: projects,
		AboutMe:  aboutMe,
		CSS:      template.CSS(css),
		JS:       template.JS(js),
	}
}
