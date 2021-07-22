package mux

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	// "webapi/setterx"
)

const (
	testKind    = "test kind"
	testMessage = "test message"
	statusOk    = 200
	statusNotOk = 400
)

func TestCreateMux(t *testing.T) {
	proxyMux := CreateMux()
	if proxyMux == nil {
		t.Fail()
		t.Logf("proxyMux was not created")
	}
}

func TestWriteError(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeError(testRecorder, testKind, testMessage)

	if testRecorder.Code != statusNotOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusNotOk, ", found: ", testRecorder.Code))
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

	if errors[0].Message != testMessage {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testMessage, ", found: ", errors[0].Message))
	}
}

func TestWriteResponse(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	writeResponse(testRecorder, testMessage, nil)

	if testRecorder.Code != statusOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusOk, ", found: ", testRecorder.Code))
	}

	var result string
	json.NewDecoder(testRecorder.Body).Decode(&result)

	if result == "" {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testMessage, ", found nil"))
		return
	}

	if result != testMessage {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testMessage, ", found: ", result))
	}
}

func TestGetBody(t *testing.T) {
	body := []interface{}{testMessage}
	bodyBytes := new(bytes.Buffer)
	errJson := json.NewEncoder(bodyBytes).Encode(&body)
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
}

func TestValidPost(t *testing.T) {
	bodyBytes := new(bytes.Buffer)
	errJson := json.NewEncoder(bodyBytes).Encode(testMessage)
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

	errValid := validPost(req)
	if errValid != nil {
		t.Fail()
		t.Logf(fmt.Sprint("expected an array, ", errValid.Error()))
		return
	}
}

func TestExec(t *testing.T) {
	getMessage := []interface{}{"INCR", "HELLO_WORLD_TEST"}
	bodyBytes := new(bytes.Buffer)
	errJson := json.NewEncoder(bodyBytes).Encode(getMessage)
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

	testRecorder := httptest.NewRecorder()
	exec(testRecorder, req)

	if testRecorder.Code != statusOk {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", statusOk, ", found: ", testRecorder.Code))
	}

	var result int64
	errJsonDecode := json.NewDecoder(testRecorder.Body).Decode(&result)
	if errJsonDecode != nil {
		t.Fail()
		t.Logf(errReq.Error())
		return
	}

	if result == 0 {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testMessage, ", found 0"))
		return
	}
}
