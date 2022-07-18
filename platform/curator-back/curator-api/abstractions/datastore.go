package abstractions

type DataStore interface {
	SayHelloToDataStoreInBackground()
	Get() []CuratedRecord
}
