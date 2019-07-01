package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Hello World!</p><a href=\"/about\">About me</a> <a href=\"/stop\">Stop server, please only use if you are are the system admin (this is very secure)</a>")
}

func aboutHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>I am hLudde</p><a href=\"/\">Home</a> <a href=\"/stop\">Stop server, please only use if you are are the system admin (this is very secure)</a>")
}

func stopServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandle)
	http.HandleFunc("/stop", stopServer)
	http.ListenAndServe(":80", nil)
}
