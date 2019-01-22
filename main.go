package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type UserInfo struct {
	Name string
}

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/editors", editor)

	server.ListenAndServe()

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	checkError(err, "indexHandler_parsefile")

	var user UserInfo
	user.Name = "megane"

	err = t.Execute(w, user)
	checkError(err, "indexHandler_execute")

}

func editor(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/layout.html", "templates/editor.html")
	checkError(err, "editor_parsefile")

	var user UserInfo
	user.Name = "megane"

	err = t.Execute(w, user)
	checkError(err, "editor_execute")

}

func checkError(err error, s string) {
	if err != nil {
		fmt.Println(s)
		os.Exit(1)
	}
}
