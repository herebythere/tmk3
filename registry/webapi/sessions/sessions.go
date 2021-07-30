package sessions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"webapi/details"
	"webapi/jwtx"
)

const (
	set = "SET"
	unableToWriteToCache = "unable to write jwt to cache"

	ApplicationJson = "application/json"
)


func writeJWTToCache(
	tokenPayload *jwtx.TokenPayload,
	err error,
) (
	bool,
	error,
) {
	if err != nil {
        return false, err
    }

	instructions := []interface{}{set, tokenPayload.Signature, tokenPayload}
	marshaledInstructions, errMarshaledInstructions := json.Marshal(instructions)
	if errMarshaledInstructions != nil {
        return false, errMarshaledInstructions
    }

	requestReader := bytes.NewReader(marshaledInstructions)

	resp, errResp := http.Post(
		details.ConfDetails.CacheAddress,
		ApplicationJson,
		requestReader,
	)
	fmt.Println(resp)
	if errResp != nil {
		return false, errResp
	}
	if resp != nil && resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, errors.New(unableToWriteToCache)
}

func readJWTFromCache(signature string, err error) {
	
}

func CreateSession(
	params *jwtx.CreateJWTParams,
	err error,
) (
	*jwtx.TokenPayload,
	error,
) {
	if err != nil {
		return nil, err
	}

	tokenPayload, errTokenPayload := jwtx.CreateJWT(params, err)
	if errTokenPayload != nil {
		return nil, errTokenPayload
	}

	fmt.Println(tokenPayload)
	successfulWrite, errSuccessfulWrite := writeJWTToCache(
		tokenPayload,
		errTokenPayload,
	)
	if successfulWrite {
		return tokenPayload, nil
	}
	
	return nil, errSuccessfulWrite
}

func ValidateSession(
	tokenPayload *jwtx.TokenPayload,
	err error,
) (bool, error) {
	if err != nil {
		return false, err
	}
	if tokenPayload == nil {
		return false, errors.New("nil token provided")
	}

	return jwtx.ValidateJWT(tokenPayload, nil)
}

func CreateAndStoreSession(
	params *jwtx.CreateJWTParams,
	err error,

) (
	*jwtx.TokenPayload,
	error,
) {

	return nil, nil
}

