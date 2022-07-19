package api

import (
	"es-curator/curator-api/abstractions"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	dataStore abstractions.DataStore
}

func CreateApi(dataStore abstractions.DataStore) *Api {
	var api *Api
	api = &Api{}
	api.dataStore = dataStore
	return api
}

func (api *Api) SayHello() {
	fmt.Println("Hello world from api.")
}

func (api *Api) StartListen() {
	router := gin.Default()

	router.GET("/records", func(ctx *gin.Context) {
		getCuratedRecords(api, ctx)
	})
	router.POST("/records", func(ctx *gin.Context) {
		createCuratedRecord(api, ctx)
	})

	router.Run("localhost:8080")
}

func createCuratedRecord(api *Api, c *gin.Context) {
	var dto CreateCuratedRecordRequest
	if err := c.BindJSON(&dto); err != nil {
		return
	}
	record := api.dataStore.Create(dto.description, dto.headline)
	c.IndentedJSON(http.StatusOK, record)
}

// responds with all curated records
func getCuratedRecords(api *Api, c *gin.Context) {
	records := api.dataStore.Get()
	c.IndentedJSON(http.StatusOK, records)
}
