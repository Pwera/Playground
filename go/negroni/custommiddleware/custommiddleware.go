package custommiddleware

import (
	"fmt"
	"github.com/OneOfOne/go-utils/memory"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
	"strings"
	"time"
)

type CustomMiddleware struct{}

func (c CustomMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("CustomMiddleware")
	next(rw, r)
}

type SecondMiddleware struct {
	*negronilogrus.Middleware
}

func NewSecondMiddleware() *SecondMiddleware {
	middleware := &SecondMiddleware{Middleware: negronilogrus.NewMiddleware()}
	middleware.Before = customBefore
	middleware.After = customAfter
	fmt.Println("Middleware size {} bytes", memory.Sizeof(middleware))
	fmt.Println("Middleware size {} bytes", memory.Sizeof(*middleware))
	fmt.Println("Middleware size {} bytes", memory.Sizeof(&middleware))
	return middleware
}

func customBefore(entry *logrus.Entry, _ *http.Request, remoteAddr string) *logrus.Entry {
	return entry.WithFields(logrus.Fields{
		"REMOTE_ADDR": remoteAddr,
		"YELLING":     true,
	})
}

func customAfter(entry *logrus.Entry, res negroni.ResponseWriter, latency time.Duration, name string) *logrus.Entry {
	fields := logrus.Fields{
		"ALL_DONE":        true,
		"RESPONSE_STATUS": res.Status(),

		fmt.Sprintf("%s_LATENCY", strings.ToUpper(name)): latency,
	}

	// one way to replace an existing entry key
	if requestId, ok := entry.Data["request_id"]; ok {
		fields["REQUEST_ID"] = requestId
		delete(entry.Data, "request_id")
	}

	return entry.WithFields(fields)
}
