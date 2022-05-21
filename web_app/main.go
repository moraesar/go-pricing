package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func main() {
	// routing
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// create a server GO Routine
	wg := sync.WaitGroup{}
	wg.Add(1)
	go serve(&wg)

	fmt.Println("Listening on port 5000")
	wg.Wait()
}

func serve(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// template

	files := []string{
		"templates/_header.html",
		"templates/home.html",
	}

	tmpls, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = tmpls.ExecuteTemplate(w, "home", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
