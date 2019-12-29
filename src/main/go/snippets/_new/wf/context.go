package wf

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Context struct {
	Request    *http.Request
	Response   http.ResponseWriter
	statusCode int
	App        *App
	DidSent    bool
}

func NewContext(res http.ResponseWriter, req *http.Request, app *App) *Context {
	ctx := Context{
		Request:    req,
		Response:   res,
		App:        app,
		statusCode: http.StatusOK,
		DidSent:    false,
	}
	ctx.Request.ParseForm()
	return &ctx
}

func (ctx *Context) reset() {
	ctx.statusCode = http.StatusOK
	ctx.DidSent = false
}
func (ctx *Context) SendStatus(statusCode int) {
	ctx.statusCode = statusCode
	ctx.Response.WriteHeader(statusCode)
}
func (ctx *Context) StatusCode() int {
	return ctx.statusCode
}
func (ctx *Context) SetHeader(key, value string) {
	ctx.Response.Header().Set(key, value)
}
func (ctx *Context) AddHeader(key, value string) {
	ctx.Response.Header().Add(key, value)
}
func (ctx *Context) GetHeader(key string) string {
	return ctx.Response.Header().Get(key)
}
func (ctx *Context) Query(key string, index ...int) (string, error) {
	if val, ok := ctx.Request.Form[key]; ok {
		if len(index) == 1 {
			return val[index[0]], nil
		}
		return val[0], nil
	}
	return "", errors.New("Query:key not found")
}
func (ctx *Context) Json(data interface{}) {
	json, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Send(json)
}
func (ctx *Context) Send(body interface{}) {
	if ctx.DidSent {
		return
	}
	switch body.(type) {
	case []byte:
		ctx.Response.Write(body.([]byte))
	case string:
		ctx.Response.Write([]byte(body.(string)))
	default:
		log.Fatal()

	}

}
func (ctx *Context) Redirect(url string) {
	ctx.SetHeader("Location", url)
	ctx.SendStatus(302)
}
func (ctx *Context) ContentType(contentType string) {
	ctx.SetHeader("Contenty-Type", contentType)
}
