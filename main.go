package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "8080", "listen port")
	flag.Parse()
	port = strings.TrimPrefix(port, ":")
	port = ":" + port
	srv := http.FileServer(http.Dir("."))
	log.Fatal(http.ListenAndServe(port, srv))
}
