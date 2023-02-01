package main

import (
	"html/template"
	"os"

	"io.github.topages/assets"
)

func render(repo *Repository) {
	categoryHTMLBytes, err := assets.Dir.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	categoryHTML := string(categoryHTMLBytes)
	categoryTemplate := template.Must(template.New("category").Parse(categoryHTML))

	categoryTemplate.Execute(os.Stdout, repo)
}
