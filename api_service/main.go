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
		// check exists token in header
		token := request.Header.Get("Authorization")
		if len(token) == 0 {
			responseWriter.WriteHeader(http.StatusUnauthorized)
			responseWriter.Write([]byte("error"))
			return
		}

		// check token prefix follow below url
		// http://www.iana.org/assignments/http-authschemes/http-authschemes.xhtml
		if !strings.HasPrefix(token, "Bearer ") {
			responseWriter.WriteHeader(http.StatusUnauthorized)
			responseWriter.Write([]byte("error"))
			return
		}

		// remove prefix: Bearer
		token = strings.TrimPrefix(token, "Bearer ")

		// check authorized
		if auth.Authorized(token) != nil {
			responseWriter.WriteHeader(http.StatusUnauthorized)
			responseWriter.Write([]byte("error"))
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
