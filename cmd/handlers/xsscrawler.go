package handlers

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/anaskhan96/soup"
)

func XSSCrawler(baseURL string, depth int, wg *sync.WaitGroup, mu *sync.Mutex, useragent []string) {
	visitedURL := make(map[string]bool)
	if depth <= 0 {
		return
	}

	urls, err := getLinks(baseURL, visitedURL)
	if err != nil {
		log.Printf("failed to get urls from %s: %v", baseURL, err)
		return
	}

	if len(urls) > 0 {
		for _, singleURL := range urls {

			mu.Lock()
			if visitedURL[singleURL] {
				mu.Unlock()
				continue
			}
			visitedURL[singleURL] = true
			mu.Unlock()

			wg.Add(1)
			go func(goUrl string) {
				defer wg.Done()
				XSSCrawler(goUrl, depth-1, wg, mu, useragent)
			}(singleURL)

			ConnectAndRequest(singleURL)
		}
	}
}
func getLinks(baseURL string, visited map[string]bool) ([]string, error) {
	var lst []string
	resp, err := soup.Get(baseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s with soup: %w", baseURL, err)
	}
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("a")
	for _, link := range links {
		linkURL := link.Attrs()["href"]

		urlJoin, err := url.JoinPath(baseURL, linkURL)
		if err != nil {
			log.Printf("link join failed: %v", err)
			continue
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
