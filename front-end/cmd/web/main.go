package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.layout.html")
	})

	http.HandleFunc("/microservices", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.layout.html")
	})

	fmt.Println("Starting front end service on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	brokerURL := os.Getenv("BROKER_URL")

	partials := []string{
		"./cmd/web/templates/base.layout.html",
		"./cmd/web/templates/header.layout.html",
		"./cmd/web/templates/footer.layout.html",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("./cmd/web/templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, brokerURL); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
