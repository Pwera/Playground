package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type Urlshortener struct {
	PathsToUrls map[string]string
	yamlBytes   []byte
	mux         *http.ServeMux
}

func (u *Urlshortener) ReadYml() {
	bb, err := ioutil.ReadFile("test.yml")
	if err != nil {
		panic(err)
	}
	u.yamlBytes = bb
}
func (u *Urlshortener) MapHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := u.PathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		if w == nil {
			println("recorder")
		}
		if r == nil {
			println("request")
		}
		u.mux.ServeHTTP(w, r)
	}
}

func (u *Urlshortener) YAMLHandler() (http.HandlerFunc, error) {
	pathUrls, err := u.parseYaml()
	if err != nil {
		return nil, err
	}
	//TODO: merge maps
	u.PathsToUrls = u.buildMap(pathUrls)
	return u.MapHandler(), nil
}

func (u *Urlshortener) parseYaml() ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(u.yamlBytes, &pathUrls)
	return pathUrls, err
}

func (u *Urlshortener) buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)

	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.Url
	}
	return pathsToUrls
}

func (u *Urlshortener) DefaultMux() {
	u.mux = http.NewServeMux()
	u.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
}
func (u *Urlshortener) MakeHandler() (http.HandlerFunc, error) {
	u.DefaultMux()
	u.MapHandler()
	u.ReadYml()
	return u.YAMLHandler()
}

func main() {
	urlshortener := Urlshortener{PathsToUrls: map[string]string{
		"/urlshort-godoc": "http://p.com",
		"/yaml-godoc":     "http://o.com",
	},
	}
	yamlHandler, err := urlshortener.MakeHandler()
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8000", yamlHandler)
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
