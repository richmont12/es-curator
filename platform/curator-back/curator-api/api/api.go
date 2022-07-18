package api

import (
	"es-curator/curator-api/abstractions"
	"net/http"

	"github.com/gin-gonic/gin"
)

var dataStore abstractions.DataStore

func StartListen(store abstractions.DataStore) {
	dataStore = store
	router := gin.Default()

	router.GET("/records", getCuratedRecords)

	router.Run("localhost:8080")
}

var records = []curatedRecord{
	{
		ID:          "1",
		Headline:    "Head1",
		Description: "desc1",
		ExternalLinks: []externalLink{
			{ID: "1", Url: "url1"},
		},
	},
}

// responds with all curared records
func getCuratedRecords(c *gin.Context) {
	records := dataStore.Get()
	c.IndentedJSON(http.StatusOK, records)
}
