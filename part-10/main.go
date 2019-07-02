package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	if err != nil {
		fmt.Println("There was an error fetching the data.")
		fmt.Println(err.Error())
		os.Exit(0)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("There was an error reading the data.")
		fmt.Println(err.Error())
		os.Exit(0)
	}
	stringBody := string(bytes)
	fmt.Print(stringBody)
	resp.Body.Close()
	os.Exit(0)
}
