package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	address  = ":8443"
	certFile = "server.crt"
	keyFile  = "server.key"
)

func main() {
	http.HandleFunc("/", handler)
	log.Printf("Listening on %s", address)
	log.Fatal(http.ListenAndServeTLS(address, certFile, keyFile, nil))
}

const content = `<!DOCTYPE html>
<html lang="en">
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="Description" content="An HTTP Server-Timing example, written in Go. You can access source file in http://github.com/jomoespe/http-server-timing-example">
<title>HTTP Server-Timing example</title>
<h1>HTTP Server-Timing example</h1>
<article>This is a <a href=https://www.w3.org/TR/server-timing/ target=_new>HTTP Server-Timing</a> example.</article>
<article>If you request the resource with query string param <code>perf</code> it will generate the <code>server-timing</code> header.</article>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html; charset=utf-8")
	writeServerTimming(w, r)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, content)
}

const (
	perfQueryParam     = "perf"
	serverTimingHeader = "server-timing"
)

func writeServerTimming(w http.ResponseWriter, r *http.Request) {
	tt := [...]string{
		"sec;desc=\"Security\";dur=0.2,",
		"service-0;desc=\"Service 0\";dur=1.1,",
		"service-1;desc=\"Service 1\";dur=.8,",
		"db-1;desc=\"Database 1\";dur=.453,",
	}
	if _, ok := r.URL.Query()[perfQueryParam]; ok {
		var b strings.Builder
		for _, t := range tt {
			b.WriteString(t)
		}
		w.Header().Add(serverTimingHeader, b.String())
	}
}
