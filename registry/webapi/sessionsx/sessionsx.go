package sessionsx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	
	jwtx "github.com/herebythere/jwtx/v0.1/golang"
)

const (
	applicationJson = "application/json"
	set = "SET"
	unableToWriteToCache = "unable to write jwt to cache"

)

var (
	localCacheAddress = os.Getenv("LOCAL_CACHE_ADDRESS")
	sessionCookieName = os.Getenv("SESSION_COOKIE_LABEL")

	errSuccessfulWrite = errors.New("nil entry was provided")
	errNoCookie = errors.New("no session cookie found")
)


// create session
// write session to cache

// read from cache

// verify session
//	 -> read from cache
//   -> not before
//	 -> issued at is after now or equal to no
//   -> expiry is not greater than now



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
		localCacheAddress,
		applicationJson,
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

	return nil, errTokenPayload

	// fmt.Println(tokenPayload)
	// successfulWrite, errSuccessfulWrite := writeJWTToCache(
	// 	tokenPayload,
	// 	errTokenPayload,
	// )
	// if successfulWrite {
	// 	return tokenPayload, nil
	// }
	
	// return nil, errSuccessfulWrite
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

func GetSessionCookie(r *http.Request, err error) (*jwtx.TokenDetails, error) {
	cookie, errCookie := r.Cookie(sessionCookieName)
	if err != nil {
		return nil, errCookie
	}
	if cookie == nil {
		return nil, errNoCookie
	}

	details, errDetails := jwtx.RetrieveTokenDetails(cookie.Value)
	if errDetails != nil {
		return nil, errDetails
	}

	currNow := time.Now().Unix()

	// check if not before

	// check iss at
	    // -> if issued before now, bail
	// check expiry
	    // -> if expiry is below now, bail with invalid request	
}

// // hello hello
// func CreateAndStoreSession(
// 	params *jwtx.CreateJWTParams,
// 	err error,

// ) (
// 	*jwtx.TokenPayload,
// 	error,
// ) {

// 	return nil, nil
// }

// // hello hello
// func CreateAndStoreGuestSession(
// 	params *jwtx.CreateJWTParams,
// 	err error,

// ) (
// 	*jwtx.TokenPayload,
// 	error,
// ) {

// 	// create guest session
// 	// store it in the redis cache
// 	return nil, nil
// }

