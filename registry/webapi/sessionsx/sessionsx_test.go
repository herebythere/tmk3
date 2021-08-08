package sessionsx

import (
	"fmt"
	"testing"

	"webapi/details"
	jwtx "github.com/herebythere/jwtx/v0.1/golang"
)

const (
	couldNotWriteToCache = "could not write to cache"
	testRegistryHost = "registry_host_test"
	tmk3 = "tmk3"
)

func TestCreateSession(t *testing.T) {
	jwtxParams := jwtx.CreateJWTParams{
		Aud: []string{testRegistryHost},
		Iss: tmk3,
		Sub: details.Credentials.Username,
		Lifetime: 3600,
	}

	tokenPayload, errTokenPayload := CreateSession(&jwtxParams, nil)
	if errTokenPayload != nil {
		t.Fail()
		t.Logf(errTokenPayload.Error())
	}
	
	fmt.Println(tokenPayload)
	fmt.Println(errTokenPayload)
	if errTokenPayload != nil {
		t.Fail()
		t.Logf(errTokenPayload.Error())
	}
}