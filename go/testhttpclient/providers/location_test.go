package providers

import (
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryRestClientError(t *testing.T) {
	// given
	countryName := "AR"
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          fmt.Sprintf(url, countryName),
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: 100,
	})

	// when
	country, err := GetCountry(countryName)

	//then
	assert.NotNil(t, err)
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, fmt.Sprintf("invalid restclient reposnse when trying to get country %s", countryName), err.Message)
}

func TestGetCountryNotFound(t *testing.T) {
	// given
	rest.FlushMockups()
	countryName := "AR"
	rest.AddMockups(&rest.Mock{
		URL:          fmt.Sprintf(url, countryName),
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	})

	// when
	country, err := GetCountry(countryName)

	//then
	assert.NotNil(t, err)
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	// given
	countryName := "AR"
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          fmt.Sprintf(url, countryName),
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Country not found","error":"not_found","status":"404"","cause":[]}`,
	})

	// when
	country, err := GetCountry(countryName)

	//then
	assert.NotNil(t, err)
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, fmt.Sprintf("invalid error interface when getting country %s", countryName), err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	// given
	countryName := "AR"
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          fmt.Sprintf(url, countryName),
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 123,"name":"Argentina"}`,
	})

	// when
	country, err := GetCountry(countryName)

	//then
	assert.NotNil(t, err)
	assert.Nil(t, country)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, fmt.Sprintf("error when trying to unmarshal country data for %s", countryName), err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	// given
	countryName := "AR"
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          fmt.Sprintf(url, countryName),
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusOK,
		RespBody: `{"id":"AR","name":"Argentina","locale":"es_AR","currency_id":"ARS","decimal_separator":",","thousands_separator":".","time_zone":"GMT-03:00","geo_information":{"location":{"latitude":-38.416096,"longitude":-63.616673}}}`,
	})

	// when
	country, err := GetCountry(countryName)

	//then
	assert.NotNil(t, country)
	assert.Nil(t, err)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 0, len(country.States))
}
