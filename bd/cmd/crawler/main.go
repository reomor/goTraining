package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gobwas/bd/pool"
)

var (
	parallelism = flag.Int("parallelism", 5, "number of parallel requests")
)

func main() {
	flag.Parse()

	var (
		results = make(chan Result)   // Will be filled with crawling results.
		done    = make(chan struct{}) // Will be closed when all results received.
	)
	go func() {
		defer func() { close(done) }() // Signal that all results received.
		for r := range results {
			fmt.Printf("got result %+v\n", r)
		}
	}()

	// Initialize new goroutine pool limited up to parallelism number.
	p := pool.NewPool(*parallelism)

	for _, s := range flag.Args() {
		// Check that given uri is valid.
		// Note that we may omit this code because uri also checked during
		// http.Get() call.
		u, err := url.ParseRequestURI(s)
		if err != nil {
			log.Fatalf("invalid url: %q: %v", s, err)
		}
		// Schedule task execution on some idle goroutine.
		// Exec() will be blocked while all goroutines are busy.
		p.Exec(func() {
			r := Result{URL: u}
			r.Count, r.Err = countAt(u)
			results <- r
		})
	}

	// Close goroutine pool to be sure that all worker goroutines are done.
	// Close() returns only when all workers are done.
	p.Close()

	// Signal to goroutine above that no more results are expected – all workers
	// are done.
	close(results)

	// Wait for all results printed.
	<-done
}

// Result contains crawling result or error.
type Result struct {
	Err   error
	URL   *url.URL
	Count int
}

func countAt(u *url.URL) (int, error) {
	resp, err := http.Get(u.String())
	if err != nil {
		return 0, err
	}

	// FIXME: this is not efficient.
	// Try to optimize it somehow on large responses.
	// Try to write some benchmarks on this function for responses with body
	// greater than 8KB.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return bytes.Count(body, []byte("go")), nil
}

func dumpResponse(resp *http.Response) {
	dump, err := httputil.DumpResponse(resp, false)
	if err != nil {
		log.Printf("can not dump request from %q", resp.Request.URL)
		return
	}
	log.Printf("got response for %q:\n%s\n\n\n", resp.Request.URL, dump)
}
