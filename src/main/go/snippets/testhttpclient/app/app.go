package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pwera/testhttpclient/controllers"
)

var (
	router = gin.Default()
)

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}

func StartApp() {
	mapUrls()
	if err := router.Run(":8090"); err != nil {
		panic(err)
	}
}
