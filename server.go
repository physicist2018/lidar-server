package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/kataras/golog"
// 	"lidarserver.sqlc/app/lidarserver/v3/database"
// )

import (
	"context"

	"github.com/kataras/golog"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://37.46.134.113:27017"))
	if err != nil {
		golog.Fatal(err)
		panic(err)
	}
	ctx := context.Background()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			golog.Fatal(err)
			panic(err)
		}
	}()

	collection := client.Database("Lidar").Collection("Experiment")
	if collection == nil {
		golog.Info(err)
	}

	v := R{Id: 1}
	res, eee := collection.InsertOne(context.Background(), v)
	if eee != nil {
		golog.Fatal(eee)
	}
	golog.Info(res.InsertedID)

	filter := bson.D{{"id", 1}}
	ress, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(ress)
	}
	golog.Info(ress.DeletedCount)
	// LoadAppConfig()

	// database.Connect(AppConfig.ConnectionString)
	// database.Migrate()

	// router := mux.NewRouter().StrictSlash(true)

	// golog.Info(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	// golog.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}
