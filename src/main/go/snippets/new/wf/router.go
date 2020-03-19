package wf

import (
	"net/url"
	"regexp"
	"strings"
)

type Router struct {
	routes []*RoutePath
}
type RoutePath struct {
	regex  *regexp.Regexp
	method string
	h      HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		routes: make([]*RoutePath, 0),
	}
}

func newRoute() *RoutePath {
	return &RoutePath{}
}

func (rt *Router) MatchRoute(urlString string, method string) (*RoutePath, error) {
	urlString = url.QueryEscape(urlString)
	if !strings.HasSuffix(urlString, "%2F") {
		urlString += "%2F"
	}
	urlString = strings.Replace(urlString, "%2F", "/", -1)
	for _, route := range rt.routes {
		matched, _ := regexp.MatchString(route.regex.String(), urlString)
		if matched {
			return route, nil
		}
	}
	return nil, nil
}

func (rt *Router) Route(pattern string, method string, handler HandlerFunc) *RoutePath {
	r := newRoute()
	r.regex = parsePattern(pattern)
	r.method = method
	r.h = handler
	return r
}

func parsePattern(pattern string) *regexp.Regexp {
	pattern = url.QueryEscape(pattern)
	if !strings.HasSuffix(pattern, "%2F") {
		pattern += "%2F"
	}
	pattern = strings.Replace(pattern, "%2F", "/", -1)
	regex := regexp.MustCompile("^" + pattern + "$")
	return regex
}
