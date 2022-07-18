package abstractions

type DataStore interface {
	SayHello()
	SayHelloToDataStoreInBackground()
	Get() []CuratedRecord
}
