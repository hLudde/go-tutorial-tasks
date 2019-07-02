package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type newsAggPage struct {
	Title    string
	Location string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my very bad Washington News Aggregator!</h1><a href=\"/news\">click here to read the news</a>")
}

func newsAggregatorHandler(w http.ResponseWriter, r *http.Request) {
	news := newsAggPage{Title: "Clickbait!", Location: "https://www.example.com"}
	t, err := template.ParseFiles("newsTemplate.html")
	if err != nil {
		fmt.Println("There was an error reading the HTML template")
		fmt.Println(err)
		os.Exit(0)
	}
	err = t.Execute(w, news)
	if err != nil {
		fmt.Println("There was an error adding data to the HTML template")
		fmt.Println(err)
		os.Exit(0)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", newsAggregatorHandler)
	http.ListenAndServe(":80", nil)
}
