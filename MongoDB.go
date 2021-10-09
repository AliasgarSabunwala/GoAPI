import "go.mongodb.org/mongo-driver/mongo"

clientOptions := options.Client().
    ApplyURI("mongodb+srv://alif:<password>@ag.znuaz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}
