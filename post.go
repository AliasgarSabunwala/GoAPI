type post struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global post array
// that we can then populate in our main function
// to simulate a database
var post []post