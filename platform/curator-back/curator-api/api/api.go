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
		api.getCuratedRecords(ctx)
	})
	router.GET("/records/:id", func(ctx *gin.Context) {
		api.getCuratedRecord(ctx)
	})
	router.POST("/records", func(ctx *gin.Context) {
		api.createCuratedRecord(ctx)
	})

	router.Run("localhost:8080")
}

func (api *Api) createCuratedRecord(c *gin.Context) {
	var dto CreateCuratedRecordRequest

	// jsonData, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println("Failed to read body.")
	// 	return
	// }
	// jsonBody := string(jsonData)
	// fmt.Println("Read body: ", jsonBody)

	if err := c.BindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "VALIDATEERR-1",
				"message": err.Error()})
		c.AbortWithError(http.StatusBadRequest, err)
		fmt.Println("Failed to bind json to create record.")
		return
	}
	record := api.dataStore.CreateCuratedRecord(dto.Description, dto.Headline)
	c.IndentedJSON(http.StatusOK, record)
}

func (api *Api) getCuratedRecords(c *gin.Context) {
	records := api.dataStore.GetAllCuratedRecords()
	c.IndentedJSON(http.StatusOK, records)
}
func (api *Api) getCuratedRecord(c *gin.Context) {
	id := c.Param("id")
	record := api.dataStore.GetCuratedRecord(id)
	c.IndentedJSON(http.StatusOK, record)
}
