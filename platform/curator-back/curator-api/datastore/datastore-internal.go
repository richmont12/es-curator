package datastore

import (
	"context"
	"es-curator/curator-api/abstractions"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func sayHello() {
	fmt.Println("Hello world from datastore.")
}

func disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func connectAndGetClient() (client *mongo.Client) {
	clientOptions := options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func get(client *mongo.Client) (records []abstractions.CuratedRecord) {
	collection := client.Database("es-curator").Collection("curatedRecords")
	records = getCore(collection)
	return
}

func getCore(collection *mongo.Collection) (records []abstractions.CuratedRecord) {
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return
	}

	var results []abstractions.CuratedRecord

	err = cur.All(context.Background(), &results)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
func create(client *mongo.Client, description string, headline string) (record abstractions.CuratedRecord) {
	collection := client.Database("es-curator").Collection("curatedRecords")
	record = abstractions.CuratedRecord{
		ID:            getNextIdOfCurratedRecord(collection),
		Headline:      headline,
		Description:   description,
		ExternalLinks: []abstractions.ExternalLink{},
	}

	insertResult, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return
}

func getNextIdOfCurratedRecord(recordCollection *mongo.Collection) (nextId string) {
	results := getCore(recordCollection)
	var highestId int64
	highestId = 0

	for _, record := range results {
		id, err := strconv.ParseInt(record.ID, 10, 64)
		if err == nil {
			if id > highestId {
				highestId = id
			}
		}
	}
	var nxtId int64
	nxtId = highestId + 1
	nextId = strconv.FormatInt(nxtId, 10)
	return
}
