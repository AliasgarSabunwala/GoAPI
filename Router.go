/*
	Adding a Router using the standard library
*/

package main
 
import (
    "fmt"
    "html"
    "log"
    "net/http"
	"net/url"
 
    // "github.com/gorilla/mux"
)
 
func main() {
 
    // router := mux.NewRouter().StrictSlash(true)
	router := NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    log.Fatal(http.ListenAndServe(":8080", router))
}
 
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}