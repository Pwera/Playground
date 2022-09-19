package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/pwera/testhttpclient/services"
	"net/http"
)

/*type CountryController struct {
	locationService *services.locationServiceInterface
}

func NewController(locationService *services.locationServiceInterface) *CountryController {
	return &CountryController{locationService: locationService}
}
*/

func GetCountry(c *gin.Context) {
	coutry, err := services.LocationService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, coutry)
}
