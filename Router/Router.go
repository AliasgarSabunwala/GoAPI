package router

import "net/http"

func main() http.Handler {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" || r.Method != http.MethodGet {
    http.NotFound(w, r)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("ok"))
})
  })
  return mux
}

s := http.Server{
  Addr:    listenAddr,
  Handler: router.NewRouter(),  // Our own instance of inbuilt servemux
}

