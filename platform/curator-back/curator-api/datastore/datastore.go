package datastore

import "es-curator/curator-api/abstractions"

type DataStore struct {
}

func CreateDataStore() *DataStore {
	return &DataStore{}
}

func (store *DataStore) SayHelloToDataStoreInBackground() {
	client := connectAndGetClient()
	get(client)
	disconnect(client)
}

func (store *DataStore) Get() (records []abstractions.CuratedRecord) {
	client := connectAndGetClient()
	if client == nil {
		return
	}
	records = get(client)

	disconnect(client)
	return
}

func (store *DataStore) Create(description string, headline string) (record abstractions.CuratedRecord) {
	client := connectAndGetClient()
	if client == nil {
		return
	}
	record = create(client, description, headline)

	disconnect(client)
	return
}
