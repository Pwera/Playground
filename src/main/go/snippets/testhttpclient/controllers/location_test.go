package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/pwera/testhttpclient/domain"
	"github.com/pwera/testhttpclient/services"
	"github.com/pwera/testhttpclient/utils/error"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

var someFunc func(countryId string) (*domain.Country, *error.ApiError)

type LocationServiceMock struct{}

func (c *LocationServiceMock) GetCountry(countryId string) (*domain.Country, *error.ApiError) {
	return someFunc(countryId)
}

func TestGetCountryNotFound(t *testing.T) {
	// given
	someFunc = func(countryId string) (*domain.Country, *error.ApiError) {
		return nil, &error.ApiError{
			Status:  http.StatusNotFound,
			Message: "CountryNotFound",
		}
	}
	services.LocationService = &LocationServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	// when
	GetCountry(c)

	//then
	assert.Equal(t, http.StatusNotFound, response.Code)
	var apiError error.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiError)
	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiError.Status)
	assert.EqualValues(t, "CountryNotFound", apiError.Message)
}

func TestGetCountryNoError(t *testing.T) {
	// given
	someFunc = func(countryId string) (*domain.Country, *error.ApiError) {
		return &domain.Country{
			Id:   "AR",
			Name: "Argentina",
		}, nil
	}
	services.LocationService = &LocationServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	// when
	GetCountry(c)

	//then
	assert.Equal(t, http.StatusOK, response.Code)
	var country domain.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
