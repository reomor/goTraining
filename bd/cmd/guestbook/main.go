package main

import (
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "addr to bind to")
)

const index = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.Title}}</title>
	</head>
	<body>
		<form method="post" action="/">
			<input type="text" name="author"/>
			<textarea name="message"></textarea>
			<input type="submit" value="post!"/>
		</form>
		{{range .Posts}}
			<p>
				<div><strong>{{ .Author }}</strong></div>
				<div>{{ .Message }}</div>
			</p>
		{{else}}
			<div><strong>no posts yet</strong></div>
		{{end}}
	</body>
</html>`

type Index struct {
	Title string
	Posts []Post
}

type Post struct {
	Author  string
	Message string
}

type Server struct {
	Pages map[string]func(io.Writer) error

	mu    sync.RWMutex
	Posts []Post
}

func (s *Server) GetPosts() []Post {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Posts
}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("got request: %s", req.URL)
	if req.Method == "POST" {
		if err := req.ParseForm(); err != nil {
			res.WriteHeader(400)
			return
		}
		post := Post{
			Author:  req.PostFormValue("author"),
			Message: req.PostFormValue("message"),
		}

		// FIXME: Replace this logic by some Storage interface capable for
		// storing/loading posts.
		// First, use this logic (simply storing in slice) as a default
		// implementation.
		// Then try to store it somewhere else – BoltDB, Sqlite, or other
		// storage.
		s.mu.Lock()
		s.Posts = append(s.Posts, post)
		s.mu.Unlock()
	}
	fn, ok := s.Pages[req.URL.Path]
	if !ok {
		res.WriteHeader(404)
		return
	}
	if err := fn(res); err != nil {
		res.WriteHeader(500)
	}
}

func main() {
	tmpl, err := template.New("index").Parse(index)
	if err != nil {
		log.Fatal(err)
	}
	var s *Server
	s = &Server{
		// FIXME: use http.NewServeMux instead.
		// This mapping is only for studying purposes.
		Pages: map[string]func(io.Writer) error{
			"/": func(w io.Writer) error {
				return tmpl.Execute(w, Index{
					Title: "My Cool Guestbook",
					Posts: s.GetPosts(),
				})
			},
		},
	}
	log.Fatal(http.ListenAndServe(*addr, s))
}
