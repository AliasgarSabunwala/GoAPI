/*
<- Getting Started with A Basic API
To get started we will have to create a very simple server which 
can handle HTTP requests. To do this we’ll create a new file 
called HomePage.go. Within this HomePage.go file we’ll want to 
define 3 distinct functions. A homePage function that will handle 
all requests to our root URL, a handleRequests function that will 
match the URL path hit with a defined function and a main function 
which will kick off our API.
*/

package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}