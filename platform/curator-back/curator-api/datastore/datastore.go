package datastore

import "es-curator/curator-api/abstractions"

type DataStore struct {
}

func CreateDataStore() *DataStore {
	return &DataStore{}
}

func (store *DataStore) SayHelloToDataStoreInBackground() {
	client := connectAndGetClient()
	getAllCuratedRecords(client)
	disconnect(client)
}

func (store *DataStore) GetAllCuratedRecords() (records []abstractions.CuratedRecord) {
	client := connectAndGetClient()
	if client == nil {
		return
	}
	records = getAllCuratedRecords(client)

	disconnect(client)
	return
}

func (store *DataStore) GetCuratedRecord(id string) (record abstractions.CuratedRecord) {
	client := connectAndGetClient()
	if client == nil {
		return
	}
	record = getCuratedRecord(client, id)

	disconnect(client)
	return
}

func (store *DataStore) CreateCuratedRecord(description string, headline string) (record abstractions.CuratedRecord) {
	client := connectAndGetClient()
	if client == nil {
		return
	}
	record = create(client, description, headline)

	disconnect(client)
	return
}
