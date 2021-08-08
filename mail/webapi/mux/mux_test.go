//	brian taylor vann

package mux

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testKind = "test kind"
)

var (
	testBody = MailRequestBody{
		Recipient: "person@example.com",
		Subject:   "this is a test email",
		Body:      "Please disregard this email, it is a test.",
	}
	errTestBody      = errors.New("test body error")
	expectedResponse = "this is a default test response"
	expectedSendMail = `echo "Please disregard this email, it is a test." | mailx -s "this is a test email" person@example.com`
)

func TestCreateMux(t *testing.T) {
	proxyMux := CreateMux()
	if proxyMux == nil {
		t.Fail()
		t.Logf("mux was not created")
	}
}

func TestValidPost(t *testing.T) {
	req, errReq := http.NewRequest("POST", "/", nil)
	if errReq != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	errValid := validPost(req)
	if errValid != nil {
		t.Fail()
		t.Logf(fmt.Sprint("expected an array, ", errValid.Error()))
		return
	}
}

func TestGetBody(t *testing.T) {
	bodyBytes := new(bytes.Buffer)
	errJson := json.NewEncoder(bodyBytes).Encode(&testBody)
	if errJson != nil {
		t.Fail()
		t.Logf(errJson.Error())
		return
	}

	req, errReq := http.NewRequest("POST", "/", bodyBytes)
	if errReq != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	reqBody, errReqBody := getBody(req, nil)
	if errReqBody != nil {
		t.Fail()
		t.Logf(fmt.Sprint("expected an array, ", errReqBody.Error()))
		return
	}
	if reqBody == nil {
		t.Fail()
		t.Logf(fmt.Sprint("request body is nil"))
		return
	}

	if reqBody.Body != testBody.Body {
		t.Fail()
		t.Logf(fmt.Sprint("expected body to equal: ", testBody.Body))
	}
}

func TestExecMailCommand(t *testing.T) {
	results, errResults := execMailCommand(&testBody, nil)
	if errResults != nil {
		t.Fail()
		t.Logf(errResults.Error())
	}

	if results != nil {
		fmt.Println(*results, errResults)
	}
}

func TestWriteError(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeError(testRecorder, testKind, errTestBody)

	if testRecorder.Code != http.StatusBadRequest {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", http.StatusBadRequest, ", found: ", testRecorder.Code))
	}

	var errors ErrorDeclarations
	json.NewDecoder(testRecorder.Body).Decode(&errors)

	if len(errors) == 0 {
		t.Fail()
		t.Logf("error array has a length of zero")
		return
	}

	if errors[0].Kind != testKind {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testKind, ", found: ", errors[0].Kind))
	}

	testBodyError := errTestBody.Error()
	if errors[0].Message != testBodyError {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testBodyError, ", found: ", errors[0].Message))
	}
}

func TestWriteResponse(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeResponse(testRecorder, &expectedResponse, nil)

	if testRecorder.Code != http.StatusOK {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", http.StatusOK, ", found: ", testRecorder.Code))
	}
}

func TestSendMail(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	bodyBytes := new(bytes.Buffer)
	errJson := json.NewEncoder(bodyBytes).Encode(&testBody)
	if errJson != nil {
		t.Fail()
		t.Logf(errJson.Error())
		return
	}

	req, errReq := http.NewRequest("POST", "/", bodyBytes)
	if errReq != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	sendMail(testRecorder, req)

	if testRecorder.Code != http.StatusOK {
		t.Fail()
		t.Logf(fmt.Sprint("expected status: ", http.StatusOK, ", instead found: ", testRecorder.Code))
		return
	}
}
