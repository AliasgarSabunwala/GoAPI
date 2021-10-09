# Backend API in Golang(Go)

Building a HTTP JSON API capable of the following operations,

1 Create an User
    Should be a POST request
    Use JSON request body
    URL should be ‘/users'

2 Get a user using id
    Should be a GET request
    Id should be in the url parameter
    URL should be ‘/users/<id here>’

3 Create a Post
    Should be a POST request
    Use JSON request body
    URL should be ‘/posts'

4 Get a post using id
    Should be a GET request
    Id should be in the url parameter
    URL should be ‘/posts/<id here>’

5 List all posts of a user
    Should be a GET request
    URL should be ‘/posts/users/<Id here>'

and here Passwords are securely stored such they can't be reverse engineered
and the server is thread safe
And has pagination to the list endpoint
And has unit tests

also it makes subsequent calls to a MongoDB database.
