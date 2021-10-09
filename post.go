

func main() {
    post = []post{
        post{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        post{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}

type post struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global post array
// that we can then populate in our main function
// to simulate a database
var post []post