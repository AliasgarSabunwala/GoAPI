package main

/*
    To get started with the driver, we import the mongo package, and create a mongo.Client:
    here only 1 function is imported because of using MongoDB Atlas ie it's online version 
*/
import "go.mongodb.org/mongo-driver/mongo"

clientOptions := options.Client().
    // ApplyURI("mongodb+srv://alif:<password>@ag.znuaz.mongodb.net/myFirstDatabase?
    ApplyURI("mongodb+srv://alif:vZJVWHXNtq5X-#c@ag.znuaz.mongodb.net/AG?
    retryWrites=true&w=majority")


/*
    connecting it to your running MongoDB server:
    We can also do this in a single step, you can use the Connect function shown below:
*/
    ctx, cancel := 
    context.WithTimeout(context.Background(), 
    10*time.Second)

/* 
    Calling Connect does not block for server discovery. 
    If you wish to know if a MongoDB server has been found and connected to, use the Ping method:
*/
    defer cancel()
    err = client.Ping(ctx, readpref.Primary())



// Make sure to defer a call to Disconnect after instantiating your client to avoid memory leak:
    defer cancel()
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
}


/*
To insert a document into a collection, 
first retrieve a Database and then Collection instance from the Client:
*/
collection := client.Database("AG").Collection("users")
collection := client.Database("AG").Collection("userid")
collection := client.Database("AG").Collection("posts")
collection := client.Database("AG").Collection("postsid")
collection := client.Database("AG").Collection("allposts")

func main() {
}
