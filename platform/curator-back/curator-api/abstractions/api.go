package abstractions

type Api interface {
	SayHello()
	StartListen(store DataStore)
}
