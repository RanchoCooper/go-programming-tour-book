package main

import (
	"flag"
)

var port string

func main() {
	flag.StringVar(&port, "p", "8000", "port")
	flag.Parse()
}
