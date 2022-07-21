package abstractions

type DataStore interface {
	SayHelloToDataStoreInBackground()
	CreateCuratedRecord(description string, headline string) CuratedRecord
	AddExternalLink(recordId string, linkUrl string) ExternalLink
	GetAllCuratedRecords() []CuratedRecord
	GetCuratedRecord(id string) CuratedRecord
}
