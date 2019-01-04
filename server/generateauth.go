package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const secret = "1234567890abc"

type jwtHeader struct {
	AuthenticationType string `json:"typ"`
	Algorithm          string `json:"alg"`
}

type jwtPayload struct {
	ReferenceUserID string      `json:"sub"`
	Data            interface{} `json:"data"`
	CreatedUnixTime string      `json:"iat"`
}

func hmac256(src string, secret string) []byte {
	key := []byte(secret)
	hmacKey := hmac.New(sha256.New, key)
	hmacKey.Write([]byte(src))
	return hmacKey.Sum(nil)
}

func GenerateJWT(requestedTime time.Time, data interface{}) string {
	header := jwtHeader{
		AuthenticationType: "JWT",
		Algorithm:          "HS256",
	}
	JSONHeader, _ := json.Marshal(header)
	encodedHeader := base64.StdEncoding.EncodeToString(JSONHeader)
	payload := jwtPayload{
		ReferenceUserID: strconv.FormatInt(requestedTime.Unix(), 10),
		Data:            data,
		CreatedUnixTime: strconv.FormatInt(requestedTime.Unix(), 10),
	}
	JSONPayload, _ := json.Marshal(payload)
	encodedPayload := base64.StdEncoding.EncodeToString(JSONPayload)

	src := fmt.Sprintf("%s.%s", encodedHeader, encodedPayload)
	signature := hmac256(src, secret)
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	jwtMessage := fmt.Sprintf("%s.%s", src, encodedSignature)
	return jwtMessage
}

type apiHandler struct{}

func (apiHandler apiHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	requestedTime := time.Now()
	data := struct {
		Username string `json:"username"`
	}{
		Username: "debugging",
	}
	responseBody := struct {
		RequestedTime string `json:"requested_time"`
		Token         string `json:"token"`
	}{
		RequestedTime: requestedTime.Format(time.RFC3339Nano),
		Token:         GenerateJWT(requestedTime, data),
	}
	responseJSONBody, _ := json.Marshal(responseBody)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(responseJSONBody)
}

func main() {
	http.ListenAndServe(":8080", &apiHandler{})
}
