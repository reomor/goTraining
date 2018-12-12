package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "addr to bind to")
)

func main() {
	// This example shows how to use http.ServeMux in some common cases.

	http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "login page")
	})
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "index page")
	})

	// Create sub routing mux, which will route requests under "/my-custom-mux"
	// path.
	mux := http.NewServeMux()
	// This will serve "/my-custom-mux/hello" path.
	mux.Handle("/hello", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// Fsprintf is like Printf() but writes to the io.Writer instead.
		fmt.Fprintf(res, "hello from my-custom-mux")
	}))
	// Register mux in http.DefaultServeMux.
	http.Handle("/my-custom-mux/", http.StripPrefix("/my-custom-mux", mux))

	// Register static files server.
	http.Handle("/static/", http.StripPrefix("/static/",
		// We will remove /static/ from request path and then lookup by the
		// remaining part of path in this directory.
		http.FileServer(http.Dir("/absolute/path/to/web/assets/directory")),
	))

	// Note that nil Handler here leads http to use its DefaultServeMux, which
	// is fulfilled by global function calls like http.Handle() and so on.
	//
	// Please jump into http.ListenAndServe() implementation for more details.
	log.Fatal(http.ListenAndServe(*addr, nil))
}
