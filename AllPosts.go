/*
	List all posts of a user
		Should be a GET request
		URL should be ‘/posts/users/<Id here>'
*/
package main

// GET Request with fields[post] and fields[user] parameters:
GET /post?include=author&fields[post]=caption,image,userid&fields[user]=name HTTP/2

// Here we want post's objects to have fields caption, image and userid only and people objects to 
// have name field only.
HTTP/2 200 OK
Content-Type: application/vnd.api+json

func main() {
{
  "data": [{
    "type": "post",
    "id": "1",
    "attributes": {
      "caption": "JSON:API paints my bikeshed!",
      "image": "The shortest article. Ever."
    },
    "relationships": {
      "userid": {
        "data": {"id": "42", "type": "people"}
      }
    }
  }],
  "included": [
    {
      "type": "people",
      "id": "42",
      "attributes": {
        "name": "John"
      }
    }
  ]
}