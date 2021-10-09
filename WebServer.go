/*
A Basic Web Server
A RESTful service starts with fundamentally being a web service first. Here is a
really basic web server that responds to any requests by simply outputting the request
url:
Running this example will spin up a server on port 8080,
and can be accessed at http://localhost:8080
*/

package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
