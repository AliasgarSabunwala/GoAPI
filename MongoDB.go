// To get started with the driver, import the mongo package, create a mongo.Client:
import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))


// And connect it to your running MongoDB server
// To do this in a single step, you can use the Connect function:
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))


//Make sure to defer a call to Disconnect after instantiating your client:
defer func() {
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
}()