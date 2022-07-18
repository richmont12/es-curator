package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartListen() {
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
	c.IndentedJSON(http.StatusOK, records)
}
