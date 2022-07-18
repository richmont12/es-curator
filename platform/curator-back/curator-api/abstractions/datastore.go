package abstractions

type DataStore interface {
	SayHelloToDataStoreInBackground()
	Create(description string, headline string) CuratedRecord
	Get() []CuratedRecord
}
