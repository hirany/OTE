package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	file := newFile()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.Handle("/room", file)
	http.HandleFunc("/editors", editor)
	http.HandleFunc("/", indexHandler)
	go file.run()

	server.ListenAndServe()

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	checkError(err)

	err = t.Execute(w, "")
	checkError(err)

}

func editor(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/layout.html", "templates/editor.html")
	checkError(err)

	err = t.Execute(w, "file name")
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
