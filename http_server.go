package main

import (
	"flag"

	"github.com/gin-gonic/gin"
)

func main() {
	var prefix string
	var root string
	var port string
	var addr string

	flag.StringVar(&prefix, "prefix", "/", "The URI prefix path. Default is /")
	flag.StringVar(&root, "root", ".", "the directory to serve files from. Defaults to the current dir")
	flag.StringVar(&port, "port", "80", "The port to listen to. Defaults to 80")
	flag.StringVar(&addr, "addr", "0.0.0.0", "The address to listen to. Defaults to 0.0.0.0")
	flag.Parse()

	r := gin.Default()

	r.Static(prefix, root)

	// Listen and serve on 0.0.0.0:80
	r.Run(addr + ":" + port)
}
