package main

import (
	"es-curator/curator-api/api"
	"es-curator/curator-api/datastore"
	"fmt"
)

func main() {
	fmt.Println("Hello world from curator-api.")

	dataStore := datastore.CreateDataStore()
	dataStore.SayHelloToDataStoreInBackground()
	api.StartListen(dataStore)
}
