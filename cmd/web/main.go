package main

import (
	"flag"
	"log"
	"net/http"
)

// Config is our config struct at runtime.
type Config struct {
	Addr      string
	StaticDir string
}

func main() {
	// Command line for our  address
	// command line flag, default address, short descriptor
	// addr := flag.String("addr", ":4000", "Http network address") // returns a pointer
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	// Parse is used to Parse the flag
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet?id=1", showSnippet)     // ANY
	mux.HandleFunc("/snippet/create", createSnippet) // Post
	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", cfg.Addr)
	err := http.ListenAndServe(cfg.Addr, mux)
	log.Fatal(err)
}
