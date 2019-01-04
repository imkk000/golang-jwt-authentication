package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
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

func hmacAlgorithm(src string, hashFunction func() hash.Hash, secret string) []byte {
	key := []byte(secret)
	hmacKey := hmac.New(hashFunction, key)
	hmacKey.Write([]byte(src))
	return hmacKey.Sum(nil)
}

func GenerateJWT(requestedTime time.Time, data interface{}) (string, error) {
	header := jwtHeader{
		AuthenticationType: "JWT",
		Algorithm:          "HS256",
	}
	JSONHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	encodedHeader := base64.StdEncoding.EncodeToString(JSONHeader)
	payload := jwtPayload{
		ReferenceUserID: strconv.FormatInt(requestedTime.Unix(), 10),
		Data:            data,
		CreatedUnixTime: strconv.FormatInt(requestedTime.Unix(), 10),
	}
	JSONPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	encodedPayload := base64.StdEncoding.EncodeToString(JSONPayload)

	signatureData := fmt.Sprintf("%s.%s", encodedHeader, encodedPayload)
	signature := hmacAlgorithm(signatureData, sha256.New, secret)
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	jwtMessage := fmt.Sprintf("%s.%s", signatureData, encodedSignature)
	return jwtMessage, nil
}

type apiHandler struct{}

func (apiHandler apiHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	data := struct {
		Username string `json:"username"`
	}{
		Username: "debugging",
	}
	requestedTime := time.Now()
	token, _ := GenerateJWT(requestedTime, data)
	responseBody := struct {
		RequestedTime string `json:"requested_time"`
		Token         string `json:"token"`
	}{
		RequestedTime: requestedTime.Format(time.RFC3339Nano),
		Token:         token,
	}
	responseJSONBody, _ := json.Marshal(responseBody)

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Write(responseJSONBody)
}

func main() {
	http.ListenAndServe(":8080", &apiHandler{})
}
