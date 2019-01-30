package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func init() {
	os.Mkdir("fileData", 755)
}

func main() {

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	file := newFile()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/editors/", editorHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/", rootHandler)
	http.Handle("/room", file)
	go file.run()

	server.ListenAndServe()

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/editors/", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
}

func editorHandler(w http.ResponseWriter, r *http.Request) {
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
