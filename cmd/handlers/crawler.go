package handlers

import (
	"log"
	"net/url"
	"strings"

	"github.com/anaskhan96/soup"
)

var visited []string
var lst []string

func Crawler(baseURL string) {
	urls, err := getLinks(baseURL)
	if err != nil {
		log.Fatalf("failed to get urls from %s: %v ", baseURL, err)
	}
	for _, singleURL := range urls {
		go func() {
			ConnectAndRequest(singleURL)
		}()
	}
}
func getLinks(baseURL string) ([]string, error) {

	resp, err := soup.Get(baseURL)
	if err != nil {
		log.Fatalf("failed to get %s with soup: %v", baseURL, err)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("a")
	for _, link := range links {
		linkURL := link.Attrs()["href"]

		urlJoin, err := url.JoinPath(baseURL, linkURL)
		if err != nil {
			log.Fatalf("link join failed: %v", err)
		}
		if strings.HasPrefix(linkURL, "http://") || strings.HasPrefix(linkURL, "https://") {
			continue
		} else if strings.HasPrefix(linkURL, "mailto:") || strings.HasPrefix(linkURL, "javascript:") {
			continue
		} else if stringInSlice(urlJoin, visited) {
			continue
		} else {
			lst = append(lst, urlJoin)
			visited = append(visited, urlJoin)
		}
	}
	return lst, nil
}

// checks if the joined url is in the visited slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
