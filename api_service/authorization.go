package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"strings"
)

const secret = "1234567890abc"
const username = "debugging"

func hmacAlgorithm(src string, hashFunction func() hash.Hash, secret string) []byte {
	key := []byte(secret)
	hmacKey := hmac.New(hashFunction, key)
	hmacKey.Write([]byte(src))
	return hmacKey.Sum(nil)
}

// TODO: ยังไม่ได้เช็คว่า token timeout หรือยัง?
func Authorized(token string) (bool, error) {
	// check token format
	tempSplitToken := strings.Split(token, ".")
	if len(tempSplitToken) != 3 {
		return false, errors.New("Invalid token format")
	}

	// check make new signature equal request signature
	// 0: header
	// 1: payload
	// 2: signature
	signatureData := fmt.Sprintf("%s.%s", tempSplitToken[0], tempSplitToken[1])
	newSignature := hmacAlgorithm(signatureData, sha256.New, secret)
	requestSignature, _ := base64.StdEncoding.DecodeString(tempSplitToken[2])
	if !hmac.Equal(requestSignature, newSignature) {
		return false, errors.New("Signature not match")
	}

	// find user
	payload, _ := base64.StdEncoding.DecodeString(tempSplitToken[1])
	fmt.Println(string(payload))
	if !strings.Contains(string(payload), fmt.Sprintf(`{"username":"%s"}`, username)) {
		return false, errors.New("Username not found in database")
	}

	// if ok return true with no error - authorized ok
	return true, nil
}
