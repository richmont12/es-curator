package datastore

func SayHello() {
	sayHello()
}

func SayHelloToDataStoreInBackground() {
	client := connectAndGetClient()
	disconnect(client)
}

func Get() (records []CuratedRecord) {
	client := connectAndGetClient()
	if client == nil {
		return
	}
	records = get(client)

	disconnect(client)
	return
}

func Create(description string, headline string) (record CuratedRecord) {
	client := connectAndGetClient()
	if client == nil {
		return
	}
	record = create(client, description, headline)

	disconnect(client)
	return
}
