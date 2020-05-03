package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

// Serve runs an HTTP server that serves static files in a directory
func Serve(r *gin.Engine, prefix string, root string, addr string) error {
	r.Static(prefix, root)

	// Listen and serve on 0.0.0.0:80
	return r.Run(addr)
}

func main() {
	var prefix string
	var root string
	var port string
	var addr string

	flag.StringVar(&prefix, "prefix", "/", "The URI prefix path.")
	flag.StringVar(&root, "root", ".", "the directory to serve files from.")
	flag.StringVar(&port, "port", "80", "The port to listen to.")
	flag.StringVar(&addr, "addr", "0.0.0.0", "The address to listen to.")
	flag.Parse()

	r := gin.Default()

	log.Fatal(Serve(r, prefix, root, addr+":"+port))
}
