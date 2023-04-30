package middleware

import (
	"log"
	"net/http"
)

type LogMiddleware struct {
	Handler http.Handler
}

func NewLogMiddleware(handler http.Handler) *LogMiddleware {
	return &LogMiddleware{Handler: handler}
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("GoInfo : HTTP api request sent to %s from %s", request.URL.Path, request.RemoteAddr)
	middleware.Handler.ServeHTTP(writer, request)

}
