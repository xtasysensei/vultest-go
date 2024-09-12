package handlers

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/anaskhan96/soup"
)

func Crawler(baseURL string, depth int, wg *sync.WaitGroup, mu *sync.Mutex) {
	visitedURL := make(map[string]bool)
	if depth <= 0 {
		return
	}

	urls, err := getLinks(baseURL, visitedURL)
	if err != nil {
		log.Fatalf("failed to get urls from %s: %v ", baseURL, err)
	}
	for _, lin := range urls {
		fmt.Println(lin)
	}

	if len(urls) > 0 {
		for _, singleURL := range urls {

			mu.Lock()
			if visitedURL[singleURL] {
				mu.Unlock()
				continue // Skip this URL if it has already been visited
			}
			visitedURL[singleURL] = true
			mu.Unlock()

			wg.Add(1)
			go func(goUrl string) {
				defer wg.Done()
				fmt.Printf("[%s]---------------------->[%s]\n", baseURL, goUrl)
				Crawler(goUrl, depth-1, wg, mu)
			}(singleURL)

			//ConnectAndRequest(singleURL)
		}
	}
}
func getLinks(baseURL string, visited map[string]bool) ([]string, error) {
	var lst []string
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
		} else {
			lst = append(lst, urlJoin)
			visited[urlJoin] = false
		}
	}
	return lst, nil
}
