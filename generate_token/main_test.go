package main_test

import (
	. "jwt/generate_token"
	"testing"
	"time"
)

func Test_GenerateJWT_Should_Be_JWT_Message(t *testing.T) {
	expectedJWTMessage := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxNTQ2MzAwODAwIiwiZGF0YSI6eyJ1c2VybmFtZSI6ImRlYnVnZ2luZyJ9LCJpYXQiOiIxNTQ2MzAwODAwIn0=.lbjBBFPXKCPj3+ZjN9cuQL4FB2L+HdYTLzxvVbks68k="
	inputTime := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)
	inputData := struct {
		Username string `json:"username"`
	}{
		Username: "debugging",
	}

	actualJWTMessage, _ := GenerateJWT(inputTime, inputData)

	if expectedJWTMessage != actualJWTMessage {
		t.Errorf("expect\n%s\nbut it got\n%s", expectedJWTMessage, actualJWTMessage)
	}
}
