package main_test

import (
	. "jwt/api_service"
	"testing"
)

func Test_Authorized_Input_Token_Should_Be_Token_OK(t *testing.T) {
	expectedTokenStatus := true

	actualTokenStatus, _ := Authorized("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxNTQ2NTg5OTQ1IiwiZGF0YSI6eyJ1c2VybmFtZSI6ImRlYnVnZ2luZyJ9LCJpYXQiOiIxNTQ2NTg5OTQ1In0=.sCRlyTUP4wTySr6bzqzjnskY0Js5N4gb3Xwpy69x5o4=")

	if expectedTokenStatus != actualTokenStatus {
		t.Errorf("expect %v but it got %v", expectedTokenStatus, actualTokenStatus)
	}
}

func Test_Authorized_Input_Token_Should_Be_Invalid_Token_Format_Error(t *testing.T) {
	expectedTokenError := "Invalid token format"
	_, actualTokenError := Authorized("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxNTQ2NTg5OTQ1IiwiZGF0YSI6eyJ1c2VybmFtZSI6ImRlYnVnZ2luZyJ9LCJpYXQiOiIxNTQ2NTg5OTQ1In0=")

	if expectedTokenError != actualTokenError.Error() {
		t.Errorf("expect %v but it got %v", expectedTokenError, actualTokenError.Error())
	}
}

func Test_Authorized_Input_Token_Should_Be_Signature_Not_Match_Error(t *testing.T) {
	expectedTokenError := "Signature not match"
	_, actualTokenError := Authorized("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.sCRlyTUP4wTySr6bzqzjnskY0Js5N4gb3Xwpy69x5o4=")

	if expectedTokenError != actualTokenError.Error() {
		t.Errorf("expect %v but it got %v", expectedTokenError, actualTokenError.Error())
	}
}

func Test_Authorized_Input_Token_Should_Be_Username_Not_Found_In_Database_Error(t *testing.T) {
	expectedTokenError := "Username not found in database"
	_, actualTokenError := Authorized("eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxNTQ2NjI2OTI0IiwiZGF0YSI6eyJ1c2VybmFtZSI6ImludmFsaWRfdXNlcm5hbWUifSwiaWF0IjoiMTU0NjYyNjkyNCJ9.htbNUq68Mx7XpUX4dEXPto3+rOfePjwhfy2AVOshd/0=")

	if expectedTokenError != actualTokenError.Error() {
		t.Errorf("expect %v but it got %v", expectedTokenError, actualTokenError.Error())
	}
}
