package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<h1>Hello World!</h1>
<p>This is a site made with multiple lines!!</p>
<p>It can even contain variables, like this: %d</p>
<a href="/stop">stop</a>		
`, 42)
}

func stopServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/stop", stopServer)
	http.ListenAndServe(":80", nil)
}
