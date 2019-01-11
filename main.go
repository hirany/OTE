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

	http.HandleFunc("/", indexHandler)

	server.ListenAndServe()

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")
	checkError(err, "indexHandler_parsefile")

	var user UserInfo
	user.Name = "megane"

	err = t.Execute(w, user)
	checkError(err, "indexHandler_execute")

}

func checkError(err error, s string) {
	if err != nil {
		fmt.Println(s)
		os.Exit(1)
	}
}
