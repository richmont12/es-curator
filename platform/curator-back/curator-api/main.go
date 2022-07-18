package main

import (
	"es-curator/curator-api/api"
	"es-curator/curator-api/datastore"
	"fmt"
)

func main() {
	fmt.Println("Hello world from curator-api.")

	dataStore := datastore.CreateDataStore()
	dataStore.SayHello()
	dataStore.SayHelloToDataStoreInBackground()
	// datastore.Get()
	api.StartListen(dataStore)
}
