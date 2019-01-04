package main

import (
	auth "jwt/api_service/authorization"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func handlerFunc(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("hello"))
}

func middlewareHandlerFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("Authorization")
		// check exists token in header
		if len(token) == 0 {
			responseWriter.WriteHeader(http.StatusUnauthorized)
			return
		}
		// check token prefix follow below url
		// http://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml
		if !strings.HasPrefix(token, "Bearer") {
			responseWriter.WriteHeader(http.StatusUnauthorized)
			return
		}

		// check authorized
		if auth.Authorized(token) != nil {
			responseWriter.WriteHeader(http.StatusUnauthorized)
			return
		}

		dumpRequest, _ := httputil.DumpRequest(request, true)
		log.Println(string(dumpRequest))
		next.ServeHTTP(responseWriter, request)
	}
}

func main() {
	http.HandleFunc("/", middlewareHandlerFunc(handlerFunc))
	http.ListenAndServe(":80", nil)
}
