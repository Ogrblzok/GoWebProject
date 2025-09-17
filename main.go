package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	text := "Here's my template!\n"
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, nil)
	check(err)
	html, err := template.ParseFiles("templates/view.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
