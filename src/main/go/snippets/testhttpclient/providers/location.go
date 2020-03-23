package providers

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/pwera/testhttpclient/domain"
	"github.com/pwera/testhttpclient/utils/error"
	"net/http"
)

var (
	url        = "https://api.mercadolibre.com/countries/%s"
	logRawJson = true
)

func GetCountry(countryId string) (*domain.Country, *error.ApiError) {
	response := rest.Get(fmt.Sprintf(url, countryId))
	if response == nil || response.Response == nil {
		return nil, &error.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("invalid restclient reposnse when trying to get country %s", countryId),
		}
	}
	if logRawJson {
		fmt.Println(string(response.Bytes()))
	}

	if response.StatusCode > 299 {
		var apiErr error.ApiError
		if err := json.Unmarshal(response.Bytes(), &apiErr); err != nil {
			return nil, &error.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("invalid error interface when getting country %s", countryId),
			}
		}
		return nil, &apiErr
	}
	var c domain.Country
	if err := json.Unmarshal(response.Bytes(), &c); err != nil {
		return nil, &error.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryId),
		}
	}

	return &c, nil
}
