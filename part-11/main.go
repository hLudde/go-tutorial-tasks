package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Sitemapindex to store the parsed xml document
type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

//Location points to the url's in the xml document
type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
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
	return sitemap
}

func getWebsiteAsString(url string) string {
	resp := requestWebsite(url)
	bytes := getResponseBody(resp)
	resp.Body.Close()
	return bytesToString(bytes)
}

func main() {
	fmt.Println(getWebsiteAsString("https://www.washingtonpost.com/news-sitemaps/index.xml"))
	fmt.Println(parseSitemapWebsite("https://www.washingtonpost.com/news-sitemaps/index.xml").Locations)
	os.Exit(0)
}
