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
	port = strings.TrimLeft(port, ":")
	port = ":" + port
	srv := &cachelessHandler{
		dirHandler: http.FileServer(http.Dir(".")),
	}
	log.Fatal(http.ListenAndServe(port, srv))
}

type cachelessHandler struct {
	dirHandler http.Handler
}

func (h *cachelessHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	head := rw.Header()
	head.Set("Cache-Control", "no-store")
	h.dirHandler.ServeHTTP(rw, r)
}
