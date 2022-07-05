package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/records", getCuratedRecords)

	router.Run("localhost:8080")
}

type externalLink struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}

type curatedRecord struct {
	ID            string         `json:"id"`
	Headline      string         `json:"headline"`
	Description   string         `json:"description"`
	ExternalLinks []externalLink `json:"external-links"`
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
