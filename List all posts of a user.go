/*
	List all posts of a user
		Should be a GET request
		URL should be ‘/posts/users/<Id here>'
*/

// Request with fields[post] and fields[user] parameters:
GET /post?include=author&fields[post]=title,body,author&fields[user]=name HTTP/2