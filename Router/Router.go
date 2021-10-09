func NewRouter() http.Handler {
  mux := http.NewServeMux()
  mux.HandleFunc("/", taskListPage())
  return mux
}
var tasks = []map[string]string{
{"id": "task-01", "title": "Post 1", "details": "a test task", "status": "active"},
{"id": "task-02", "title": "Post 2", "details": "# Title\n\n+ Some Stuff to do\n+ Another item\n\nSomething else written here", "status": "active"},
{"id": "task-03", "title": "Post 3", "details": "a test task 3", "status": "active"},
{"id": "task-04", "title": "Post 4", "details": "a test task 4", "status": "active"},
{"id": "task-05", "title": "Post 5", "details": "a test task 5", "status": "active"},
{"id": "task-06", "title": "Post 6", "details": "a test task 6", "status": "pending"},
{"id": "task-07", "title": "Post 7", "details": "a test task 7", "status": "pending"},
{"id": "task-08", "title": "Test 8", "details": "a test task 8", "status": "active"},
{"id": "task-09", "title": "Test 9", "details": "a test task 9", "status": "inactive"},
{"id": "task-10", "title": "Test 10", "details": "a test task 10", "status": "inactive"},
}

// Change the http.NotFound to notFoundPage(w, r)
func taskListPage() http.HandlerFunc {
  ...
  return func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" || r.Method != http.MethodGet {
      notFoundPage(w, r)
      return
    }
    ...
  }
}
... at the end of the file add this function
var notFoundPage = func() http.HandlerFunc {
  files := tmplLayout("./web/templates/_notFound.gohtml")
  tmpl := template.Must(template.New("index").Funcs(defaultFuncs).ParseFiles(files...))
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var buf bytes.Buffer
    if err := tmpl.ExecuteTemplate(&buf, "base", nil); err != nil {
      fmt.Printf("ERR: %v\n", err)
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }
    w.WriteHeader(http.StatusOK)
    io.Copy(w, &buf)
  })
}()