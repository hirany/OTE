package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	//"text/template"

	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/stretchr/objx"
)

func init(){
	goth.UseProviders(
		github.New(os.Getenv("GOSIMPLEWEBAPP_GITHUB_ID"), os.Getenv("GOSIMPLEWEBAPP_GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
	)
}

type templateHandler struct {
	filename string
	once     sync.Once
	templ   *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	t.once.Do(func() {
		t.templ = template.Must(
			template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	data := make(map[string]interface{})

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["Userdata"] = objx.MustFromBase64(authCookie.Value)
	}

	t.templ.Execute(w, data)
}

func main() {

	/*server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}*/

	//http.HandleFunc("/", indexHandler)

	p := pat.New()
	p.Get("/auth/{provider}/callback", callbackHandler)
	p.Get("/auth/{provider}", gothic.BeginAuthHandler)
	p.Get("/logout", logoutHandler)
	p.Add("GET", "/login", (&templateHandler{filename: "login.html"}))
	p.Add("GET", "/", MustAuth(&templateHandler{filename: "index.html"}))

	//server.ListenAndServe()

	log.Fatal(http.ListenAndServe(":3000", p))

}



func indexHandler(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")
	checkError(err, "indexHandler")

	err = t.Execute(w, "str")
	checkError(err, "execute")

}

func checkError(err error, s string) {
	if err != nil {
		fmt.Println(s)
		os.Exit(1)
	}
}
