package services

import (
	"github.com/pwera/testhttpclient/domain"
	"github.com/pwera/testhttpclient/providers"
	"github.com/pwera/testhttpclient/utils/error"
)

type locationServiceInterface interface {
	GetCountry(countryId string) (*domain.Country, *error.ApiError)
}

type locationService struct {
}

var (
	LocationService locationServiceInterface
)

func init(){
	LocationService = &locationService{}
}

func (s *locationService) GetCountry(countryId string) (*domain.Country, *error.ApiError) {
	return providers.GetCountry(countryId)
}
