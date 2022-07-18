package main

import (
	"es-curator/curator-api/datastore"
	"fmt"
)

func main() {
	fmt.Println("Hello world from curator-api.")
	datastore.SayHello()
	datastore.SayHelloToDataStoreInBackground()
	datastore.Get()
}
