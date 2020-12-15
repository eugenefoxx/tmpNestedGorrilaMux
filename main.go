package main

import (
	//	"html/template"

	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	tpl    *template.Template
	router = mux.NewRouter()
)

//var router = mux.NewRouter()

func main() {
	//tpl, _ = template.ParseGlob("templates/*.html")
	//http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./js"))))
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources", http.FileServer(http.Dir("./js"))))
	tpl = template.Must(tpl.ParseGlob("templates/*.html"))
	//	if err != nil {
	//		log.Fatal("Error loading templates:" + err.Error())
	//	}
	router.HandleFunc("/", indexHandler())
	router.HandleFunc("/modul", modulHandler())

	fmt.Println("Start ...")
	log.Fatal(http.ListenAndServe(":8081", router))

}

func indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index2.html", nil)
	}
}

func modulHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "index3.html", nil)
	}
}
