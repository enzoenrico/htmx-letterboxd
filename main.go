package main

import (
	"fmt"
	"html/template"
	"kyo/letterboxd/backend"
	"net/http"
	"time"
)

func main() {
	serveAPI()
	homepage := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("./html/index.html"))
		templ.Execute(w, nil)
	})
	http.HandleFunc("/", homepage)

	http.ListenAndServe(":8080", nil)
}

func serveAPI(){
	data := backend.RSSToJSON()

	handleData := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		for _, item := range data {
			retVal := string(item.Title + "\n" + item.Review + "\n")
			fmt.Fprintf(w, "------------------------------------------------\n")
			fmt.Fprintf(w, retVal)
			fmt.Fprintf(w, "------------------------------------------------\n")
		}

	})


	http.HandleFunc("/api/getMovies", handleData)
}

