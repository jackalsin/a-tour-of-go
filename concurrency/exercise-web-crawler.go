package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Status int

const (
	started Status = iota
)

type Tasks struct {
	c     sync.Mutex
	tasks map[string]Status
}

var tasks = Tasks{
	tasks: make(map[string]Status),
}

func (t *Tasks) PutIfAbsent(url string, status Status) bool {
	t.c.Lock()
	defer t.c.Unlock()
	_, exists := t.tasks[url]
	if !exists {
		t.tasks[url] = status
	}
	return !exists
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c chan bool) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	//defer sendCloseSignal(c)
	if depth <= 0 {
		sendCloseSignal(c, url)
		return
	}
	if added := tasks.PutIfAbsent(url, started); !added {
		sendCloseSignal(c, url)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		sendCloseSignal(c, url)
		return
	}
	childChan := make(chan bool, len(urls))
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, childChan)
	}
	for range urls {
		<-childChan
	}
	sendCloseSignal(c, url)
	return
}

func sendCloseSignal(c chan bool, url string) {
	fmt.Println("\t\t\t\tBefore sending turn off", url)
	c <- true
}

func main() {
	c := make(chan bool)
	Crawl("https://golang.org/", 4, fetcher, c)
	<-c
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
