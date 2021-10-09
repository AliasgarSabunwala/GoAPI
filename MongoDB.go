import "go.mongodb.org/mongo-driver/mongo"

clientOptions := options.Client().
    ApplyURI("mongodb+srv://alif:<vZJVWHXNtq5X-#c>@ag.znuaz.mongodb.net/AG?retryWrites=true&w=majority")
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
client, err := mongo.Connect(ctx, clientOptions)
if err != nil {
    log.Fatal(err)
}
