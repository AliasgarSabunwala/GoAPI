/*
	List all posts of a user
		URL should be ‘/posts/users/<Id here>'
*/
package main

import (
    "fmt" 
    "http/url"
    "database/sql"
)    

func main() {

    // creating an URL object with a target URL string that accepts the JSON data via HTTP POST method:
    URL url = new URL ("mongodb+srv://alif:vZJVWHXNtq5X-#c@ag.znuaz.mongodb.net/AG?
    retryWrites=true&w=majority");


    /*
        -> Opening a Connection
        From the above URL object, we can invoke the openConnection method to get the HttpURLConnection object.

        We can't instantiate HttpURLConnection directly, as it's an abstract class:
    */
    HttpURLConnection con = (HttpURLConnection)url.openConnection();


    /*
        Set the Request Method
        To send a GET request, we'll have to set the request method property to GET:
    */
        con.setRequestMethod("GET");

        
    /*
        Set the Request Content-Type Header Parameter
        Set the “content-type” request header to “application/json” to send the request content in JSON form. 
        This parameter has to be set to send the request body in JSON format.

        Failing to do so, the server returns HTTP status code “400-bad request”:
    */
    con.setRequestProperty("Content-Type", "application/json; utf-8");
    // Also, note that we've mentioned charset encoding along with content type. This is useful 
    // if the request content encoding is different from UTF-8 encoding, which is the default encoding.


    /*
        Setting up Response Format Type
        Set the “Accept” request header to “application/json” to read the response in the desired format:
    */
        con.setRequestProperty("Accept", "application/json");


    /*
        Ensuring the Connection Will Be Used to Send Content
        To send request content, let's enable the URLConnection object's doOutput property to true.
        Otherwise, we won't be able to write content to the connection output stream:
    */  
        con.setDoOutput(true);


    /*
        Create the Request Body
        After creating a custom JSON String:
    */
        String jsonInputString = "{"name": "Upendra", "job": "Programmer"}";
    
    //We would need to write it:
        try(OutputStream os = con.getOutputStream()) {
            byte[] input = jsonInputString.getBytes("utf-8");
            os.write(input, 0, input.length);			
        }


    /*        
    Read the Response From Input Stream
    Get the input stream to read the response content. 
    Remember to use try-with-resources to close the response stream automatically.
    Read through the whole response content, and print the final response string:
    */
    freestar
    try(BufferedReader br = new BufferedReader(
    new InputStreamReader(con.getInputStream(), "utf-8"))) {
        StringBuilder response = new StringBuilder();
        String responseLine = null;
        while ((responseLine = br.readLine()) != null) {
            response.append(responseLine.trim());
        }
        System.out.println(response.toString());
    }

    
    /*
      In the Above Code, If the response is in JSON format, 
      we can use any 3rd-party JSON parsers to parse the response.
    */
    

    // Showing all posts 
{
  "data": [{
    "type": "post",
    "id": "1",
    "attributes": {
      "caption": "JSON:API!",
      "image": "https://www.google.com/search?q=welcome&rlz=1C1FHFK_enIN966IN966&sxsrf=AOaemvLHeSoIV8pKnfbO5kePJlhe7GA_HQ:1633796258488&tbm=isch&source=iu&ictx=1&fir=x2tw04iHmein1M%252CsSYd5x4bkIAjWM%252C_%253Bc4qkAXIC3Bh4kM%252C0qGCk2f6ihu80M%252C_%253BqrfDxlkCa7VaxM%252C4pNK5z8-7ld4HM%252C_%253BCmcMrNJJX6OKIM%252ClaVEV2CnnE4sGM%252C_%253BlIq6whI-Spp37M%252C9vti9rG8MwYXtM%252C_%253BpDnWszSTxtuBwM%252CIkep7NGQyvaRzM%252C_%253BrQ730mfy8QvNJM%252CxtGH6tKXgXET3M%252C_%253BRnBGunY_-A2P-M%252CxfeDuXhKUFJvVM%252C_%253BnigP9KVBocKUXM%252CRBDFHuF7VYc08M%252C_%253BMKVB06E6nP5NsM%252CWSH2swCJ7qtAsM%252C_&vet=1&usg=AI4_-kTbzCUbdGWyol2mHnj0HLfjxUr37A&sa=X&sqi=2&ved=2ahUKEwjO9Pys3b3zAhVbqZUCHbQ9D9IQ9QF6BAgKEAE#imgrc=x2tw04iHmein1M"
    },
    "relationships": {
      "userid": {
        "data": {"id": "1", "type": "data"}
      }
    }
  }],
  "included": [
    {
      "type": "post",
      "id": "1",
      "attributes": {
        "user": "John"
        "caption": "JSON:API!",
        "image": "https://www.google.com/search?q=welcome&rlz=1C1FHFK_enIN966IN966&sxsrf=AOaemvLHeSoIV8pKnfbO5kePJlhe7GA_HQ:1633796258488&tbm=isch&source=iu&ictx=1&fir=x2tw04iHmein1M%252CsSYd5x4bkIAjWM%252C_%253Bc4qkAXIC3Bh4kM%252C0qGCk2f6ihu80M%252C_%253BqrfDxlkCa7VaxM%252C4pNK5z8-7ld4HM%252C_%253BCmcMrNJJX6OKIM%252ClaVEV2CnnE4sGM%252C_%253BlIq6whI-Spp37M%252C9vti9rG8MwYXtM%252C_%253BpDnWszSTxtuBwM%252CIkep7NGQyvaRzM%252C_%253BrQ730mfy8QvNJM%252CxtGH6tKXgXET3M%252C_%253BRnBGunY_-A2P-M%252CxfeDuXhKUFJvVM%252C_%253BnigP9KVBocKUXM%252CRBDFHuF7VYc08M%252C_%253BMKVB06E6nP5NsM%252CWSH2swCJ7qtAsM%252C_&vet=1&usg=AI4_-kTbzCUbdGWyol2mHnj0HLfjxUr37A&sa=X&sqi=2&ved=2ahUKEwjO9Pys3b3zAhVbqZUCHbQ9D9IQ9QF6BAgKEAE#imgrc=x2tw04iHmein1M"
      }
    }
  ]
}