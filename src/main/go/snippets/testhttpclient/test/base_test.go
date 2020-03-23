package test

/*
func TestMain(m *testing.M) {
	go app.StartApp()
	os.Exit(m.Run())
}


func TestGetCountryNoError(t *testing.T) {
	response := rest.Get("http://localhost:8090/locations/countries/AR")
	assert.NotNil(t, response)
	assert.Nil(t, response.Err)

	var country domain.Country
	json.NewDecoder(response.Body).Decode(&country)
	fmt.Println(country)
}



func TestGetCountryNotFound(t *testing.T) {
	// given
	rest.StartMockupServer()
	countryName := "AR5"
	rest.AddMockups(&rest.Mock{
		URL:          fmt.Sprintf("https://api.mercadolibre.com/countries/%s", countryName),
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"Country not found","error":"not_found","status":404,"cause":[]}`,
	})
	// when
	response := rest.Get("http://localhost:8090/locations/countries/AR5")

	//then
	assert.NotNil(t, response)
	assert.Nil(t, response.Err)

	var apiErr error.ApiError
	json.NewDecoder(response.Body).Decode(&apiErr)
	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}
*/