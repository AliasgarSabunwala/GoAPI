package main

import (
	"fmt"
	"net/http"
)

func runHttp(listenAddr string) error {
	s := http.Server{
		Addr:    listenAddr,
		Handler: http.DefaultServeMux,
	}
	fmt.Printf("Starting HTTP listener at %s\n", listenAddr)
	return s.ListenAndServe()
}
