package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ExternalLink struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}
type CuratedRecord struct {
	ID            string         `json:"id"`
	Headline      string         `json:"headline"`
	Description   string         `json:"description"`
	ExternalLinks []ExternalLink `json:"external-links"`
}

func main() {
	client := connectAndGetClient()

	disconnect(client)
}

func SayHello() {
	client := connectAndGetClient()
	if client == nil {
		return
	}

	disconnect(client)
}

func Insert(description string, headline string) (record CuratedRecord) {
	client := connectAndGetClient()
	if client == nil {
		return
	}
	record = insertCore(client, description, headline)

	disconnect(client)
	return
}

func disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func connectAndGetClient() (client *mongo.Client) {
	// Set client options
	// Connect to MongoDB
	// Check the connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

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

func insertCore(client *mongo.Client, description string, headline string) (record CuratedRecord) {
	collection := client.Database("es-curator").Collection("curatedRecords")
	record = CuratedRecord{
		ID:            getNextIdOfCurratedRecord(collection),
		Headline:      headline,
		Description:   description,
		ExternalLinks: []ExternalLink{},
	}

	insertResult, err := collection.InsertOne(context.TODO(), record)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return
}

func getNextIdOfCurratedRecord(recordCollection *mongo.Collection) (nextId string) {
	ctx := context.TODO()
	cur, err := recordCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
		return
	}

	var results []CuratedRecord

	err = cur.All(context.Background(), &results)
	if err != nil {
		log.Fatal(err)
		return
	}

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
