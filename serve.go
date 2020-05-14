package main

import (
	"flag"
	"log"
	"net/http"
)

var prefix string
var root string
var port string
var addr string

func parseFlags() {
	flag.StringVar(&prefix, "prefix", "/", "The URI prefix path.")
	flag.StringVar(&root, "root", ".", "the directory to serve files from.")
	flag.StringVar(&port, "port", "80", "The port to listen to.")
	flag.StringVar(&addr, "addr", "0.0.0.0", "The address to listen to.")
	flag.Parse()
}

func newServer() *http.ServeMux {
	return http.NewServeMux()
}

// RegisterStatic runs an HTTP server that serves static files in a directory
func RegisterStatic(r *http.ServeMux, prefix string, root string) {
	r.Handle(prefix, http.FileServer(http.Dir(root)))
}

func main() {
	parseFlags()

	r := newServer()

	RegisterStatic(r, prefix, root)

	// Listen and serve on 0.0.0.0:80
	// log.Fatal(r.Run(addr+":"+port))
	log.Fatal(http.ListenAndServe(addr+":"+port, r))
}
