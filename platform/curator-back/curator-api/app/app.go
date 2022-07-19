package app

import (
	"es-curator/curator-api/api"
	"es-curator/curator-api/datastore"
	"fmt"
)

func Start() {
	fmt.Println("Hello world from app.")

	dataStore := datastore.CreateDataStore()
	dataStore.SayHelloToDataStoreInBackground()

	api := api.CreateApi(dataStore)
	api.SayHello()
	api.StartListen()
}
