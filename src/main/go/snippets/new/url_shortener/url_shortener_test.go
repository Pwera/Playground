package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Redirect(t *testing.T) {
	data := []struct {
		name string
		path string
		url  string
	}{
		{name: "Redirect test", url: "https://github.com/gophercises/urlshort", path: "/urlshort"},
	}
	for _, c := range data {
		t.Run(c.name, func(t *testing.T) {
			//when
			request, err := http.NewRequest("GET", "http://localhost:8000"+c.path, nil)
			if err != nil {
				t.Fatalf("Couldn't create reuest %v", err)
			}
			recorder := httptest.NewRecorder()

			us := Urlshortener{PathsToUrls: map[string]string{
				"/urlshort-godoc": "http://p.com",
				"/yaml-godoc":     "http://o.com",
			},
			}
			mapHandler, err := us.MakeHandler()

			//given
			mapHandler(recorder, request)

			//then
			assert.Nil(t, err)
			assert.NotNil(t, recorder)
			assert.Equal(t, recorder.Header().Get("Location"), "https://github.com/gophercises/urlshort")
		})
	}

}
