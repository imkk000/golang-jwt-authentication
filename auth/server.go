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

func GenerateJWT(data interface{}) []byte {
	header := jwtHeader{
		AuthenticationType: "JWT",
		Algorithm:          "HS256",
	}
	JSONHeader, _ := json.Marshal(header)
	encodedHeader := base64.StdEncoding.EncodeToString(JSONHeader)
	payload := jwtPayload{
		ReferenceUserID: strconv.FormatInt(time.Now().Unix(), 10),
		Data:            data,
		CreatedUnixTime: strconv.FormatInt(time.Now().Unix(), 10),
	}
	JSONPayload, _ := json.Marshal(payload)
	encodedPayload := base64.StdEncoding.EncodeToString(JSONPayload)

	src := fmt.Sprintf("%s.%s", encodedHeader, encodedPayload)
	signature := hmac256(src, secret)
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	jwtMessage := fmt.Sprintf("%s.%s", src, encodedSignature)
	return []byte(jwtMessage)
}

type apiHandler struct{}

func (apiHandler apiHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	data := struct {
		Username string `json:"username"`
	}{
		Username: "debugging",
	}

	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(GenerateJWT(data))
}

func main() {
	http.ListenAndServe(":8080", &apiHandler{})
}
