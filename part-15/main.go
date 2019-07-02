package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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

func parseNewsXMLWebsite(url string) News {
	var news News
	resp := requestWebsite(url)
	bytes := getResponseBody(resp)
	resp.Body.Close()
	xml.Unmarshal(bytes, &news)
	return news
}

func getWebsiteAsString(url string) string {
	resp := requestWebsite(url)
	bytes := getResponseBody(resp)
	resp.Body.Close()
	return bytesToString(bytes)
}

func printSitemapLoop(s Sitemapindex) {
	for _, Location := range s.Locations {
		fmt.Printf("%s\n", Location)
	}
}

func main() {
	var sitemap Sitemapindex
	var news News

	newsMap := make(map[string]NewsMap)
	sitemap = parseSitemapWebsite("https://www.washingtonpost.com/news-sitemaps/index.xml")

	for _, Location := range sitemap.Locations {
		news = parseNewsXMLWebsite(Location)
		for i := range news.Titles {
			newsMap[news.Titles[i]] = NewsMap{news.Keywords[i], news.Locations[i]}
		}
	}

	for i, data := range newsMap {
		fmt.Printf("Title: %s, Link: %s\n", i, data.Location)
	}

	os.Exit(0)
}
