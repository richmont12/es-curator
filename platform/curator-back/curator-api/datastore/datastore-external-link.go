package datastore

import (
	"es-curator/curator-api/abstractions"

	"github.com/google/uuid"
)

func (store *DataStore) AddExternalLink(recordId string, linkUrl string) abstractions.ExternalLink {
	client := connectAndGetClient()
	record := getCuratedRecord(client, recordId)

	externalLink := abstractions.ExternalLink{Url: uuid.New().String()}

	record.ExternalLinks = append(record.ExternalLinks, externalLink)

	disconnect(client)

	return externalLink
}
