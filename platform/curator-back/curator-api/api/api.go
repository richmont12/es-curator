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
	router.POST("/records", createCuratedRecord)

	router.Run("localhost:8080")
}

func createCuratedRecord(c *gin.Context) {
	var dto CreateCuratedRecordRequest
	if err := c.BindJSON(&dto); err != nil {
		return
	}
	record := dataStore.Create(dto.description, dto.headline)
	c.IndentedJSON(http.StatusOK, record)
}

// responds with all curated records
func getCuratedRecords(c *gin.Context) {
	records := dataStore.Get()
	c.IndentedJSON(http.StatusOK, records)
}
