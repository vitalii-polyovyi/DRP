package database

import (
	"context"
	"fmt"
	"log"
	"os"

	// "sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func getConnectionUri() string {
	return fmt.Sprintf(
		"mongodb://%s:%s/?authSource=%s",
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
		os.Getenv("MONGO_DATABASE"),
	)
}

func GetClient() *mongo.Client {
	if client == nil {
		var dbName string = os.Getenv("MONGO_DATABASE")
		var uri string = getConnectionUri()

		log.Printf("Connecting to Database %s...", dbName)

		credentials := options.Credential{
			AuthSource: dbName,
			Username:   os.Getenv("MONGO_USER"),
			Password:   os.Getenv("MONGO_PASSWORD"),
		}
		clientOptions := options.Client().SetAuth(credentials).ApplyURI(uri)
		client, _ = mongo.Connect(context.TODO(), clientOptions)

		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
			log.Fatal(err)
		}

		log.Printf("Connected to Database %s", dbName)
	}

	return client
}

func GetDatabase() *mongo.Database {
	return GetClient().Database(os.Getenv("MONGO_DATABASE"))
}

func CloseConnection() {
	if IsUp() {
		log.Println("Closing connection to database...")
		if err := GetClient().Disconnect(context.TODO()); err != nil {
			log.Fatal("Error closing connection to database: ", err)
		} else {
			log.Println("Closed connection to database")
		}
	}
}

func IsUp() bool {
	return client != nil
}
