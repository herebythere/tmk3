package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	testKind    = "test kind"
	testMessage = "test message"
	statusOk    = 200
	statusNotOk = 400
)

func TestWriteError(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	WriteError(testRecorder, testKind, testMessage)

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
	WriteResponse(testRecorder, testMessage, nil)

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
