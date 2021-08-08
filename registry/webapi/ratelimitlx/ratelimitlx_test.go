package ratelimitlx

import (
	"net/http"
	"testing"
)

const (
	increment                    = "INCR"
	testJSONIncrement            = "test_json_increment"
)

func TestPostJSONRequest(t *testing.T) {
	instructions := []interface{}{increment, testJSONIncrement}
	resp, errResp := postJSONRequest(instructions)
	if errResp != nil {
		t.Fail()
		t.Logf(errResp.Error())
	}
	if resp == nil {
		t.Fail()
		t.Logf("set service was not successfuul")
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fail()
		t.Logf("response code was not 200")
	}
}
