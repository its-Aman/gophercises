package main

import (
	"Sitemap_Builder/HTMLparser"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type empty struct{}
type loc struct {
	Value string `xml:"loc"`
}
type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	siteUrl := "https://www.calhoun.io"
	// siteUrl := "https://gophercises.com"

	url := flag.String("URL", siteUrl, "URL to build the sitemap")
	depth := flag.Int("Depth", 1<<3, "Depth to crawl upto the sitemap")
	flag.Parse()

	pages := bfs(*url, *depth)

	toXml := urlset{
		Urls:  make([]loc, len(pages)),
		Xmlns: xmlns,
	}

	for i, link := range pages {
		// toXml.Urls = append(toXml.Urls, loc{link})
		toXml.Urls[i] = loc{link}
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}

	fmt.Println()
}

func bfs(urlStr string, maxDepth int) []string {
	seen := make(map[string]empty)
	var q map[string]empty
	nq := map[string]empty{
		urlStr: empty{},
	}

	for i := 0; i <= maxDepth && len(nq) > 0; i++ {
		q, nq = nq, make(map[string]empty)

		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = empty{}

			for _, link := range get(url) {
				if _, ok := seen[link]; !ok {
					nq[link] = empty{}
				}
			}
		}
	}

	ret := make([]string, 0, len(seen))

	for url := range seen {
		ret = append(ret, url)
	}

	return ret
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		return []string{}
	}

	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
	links, _ := HTMLparser.Parse(r)
	var ret []string

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string

	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}

func start(url string, depth int, visited *map[string]struct{}) []HTMLparser.Link {
	domain := ""

	if _, ok := (*visited)[url]; ok || depth == 0 {
		return []HTMLparser.Link{}
	}

	(*visited)[url] = struct{}{}
	// fmt.Println("URL: ", url)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error while crawling...", err)
		return []HTMLparser.Link{}
	}

	defer resp.Body.Close()
	links, err := HTMLparser.Parse(resp.Body)

	if err != nil {
		fmt.Println("Some error while fetching the links")
		return []HTMLparser.Link{}
	}

	var nextLinks []HTMLparser.Link

	for _, link := range links {
		if _, ok := (*visited)[link.Href]; !ok {

			if strings.HasPrefix(link.Href, domain) {
				nextLinks = append(nextLinks, start(link.Href, depth-1, visited)...)
			}

			if strings.HasPrefix(link.Href, "/") {
				nextLinks = append(nextLinks, start(domain+link.Href, depth-1, visited)...)
			}
		}
	}

	links = append(links, nextLinks...)

	return links
}
