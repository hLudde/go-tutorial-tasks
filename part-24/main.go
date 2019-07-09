package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

//Sitemapindex to store the parsed xml document
type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

//News elements in the sitemap
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

//NewsMap map the Keyword and Url to their title
type NewsMap struct {
	Keyword  string
	Location string
}

type newsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", newsAggregatorHandler)
	http.ListenAndServe(":80", nil)
}

func printSitemapLoop(s Sitemapindex) {
	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my very bad Washington News Aggregator!</h1><a href=\"/news\">click here to read the news</a>")
}

func newsAggregatorHandler(w http.ResponseWriter, r *http.Request) {
	var sitemap Sitemapindex
	queue := make(chan News, 30)
	newsMap := make(map[string]NewsMap)
	sitemap = parseSitemapWebsite("https://www.washingtonpost.com/news-sitemaps/index.xml")

	for _, Location := range sitemap.Locations {
		wg.Add(1)
		go parseNewsXMLWebsite(queue, Location)
	}

	wg.Wait()
	close(queue)

	for i := range queue {
		for id := range i.Keywords {
			newsMap[i.Titles[id]] = NewsMap{i.Keywords[id], i.Locations[id]}
		}
	}

	newsAgg := newsAggPage{Title: "News Aggregator!", News: newsMap}
	t, err := template.ParseFiles("newsTemplate.html")
	if err != nil {
		fmt.Println("There was an error reading the HTML template")
		fmt.Println(err)
		os.Exit(0)
	}
	err = t.Execute(w, newsAgg)
	if err != nil {
		fmt.Println("There was an error adding data to the HTML template")
		fmt.Println(err)
		os.Exit(0)
	}
}

func requestWebsite(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("There was an error fetching the data.")
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return resp
}

func getResponseBody(resp *http.Response) []byte {
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("There was an error reading the data.")
		fmt.Println(err.Error())
		os.Exit(0)
	}
	return bytes
}

func bytesToString(bytes []byte) string {
	return string(bytes)
}

func parseSitemapWebsite(url string) Sitemapindex {
	var sitemap Sitemapindex
	resp := requestWebsite(url)
	bytes := getResponseBody(resp)
	resp.Body.Close()
	xml.Unmarshal(bytes, &sitemap)
	for i := 0; i < len(sitemap.Locations); i++ {
		sitemap.Locations[i] = strings.TrimPrefix(strings.TrimSuffix(sitemap.Locations[i], "\n"), "\n")
	}
	return sitemap
}

func parseNewsXMLWebsite(cahnnel chan News, url string) {
	defer wg.Done()
	fmt.Println("Requesting from", url)
	var news News
	resp := requestWebsite(url)
	bytes := getResponseBody(resp)
	resp.Body.Close()
	xml.Unmarshal(bytes, &news)
	cahnnel <- news
}
