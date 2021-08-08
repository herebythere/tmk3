package ratelimit

import (
	"fmt"
	"net/http"
	"os"
	"time"

	jwtx "github.com/herebythere/jwtx/v0.1/golang"

	"webapi/details"
)

// send http request for INCR based on IP
//  get forwarded from
//  get reported ip

// 

// take in request
// increment HINCBY registry_ratelimit <JWT> 1
// inverse is HGET registry_ratelimit <JWT>
// 	get body

// determine if past ratelimit, bool err

const (
	hincrby = "HINCRBY"
	incrementValue = 1
	colonDelimiter = ":"
)

var (
	serviceName = os.Getenv("SERVICE_NAME")
	localCacheAddress = os.Getenv("LOCAL_CACHE_ADDRESS")

	// the "channel" of cache stuff
	errNoCookie = errors.New("request does not contain a valid cookie")
	errReqUnsuccessful = errors.New("cache request was not successful")
)

// untested
func PostJSONRequest(
	address string,
	entry interface{},
	err error,
) (
	*http.Response,
	error,
) {
	if err != nil {
		return nil, err
	}

	entryAsJSON, errEntryAsJSON := json.Marshal(entry)
	if errEntryAsJSON != nil {
		return nil, errEntryAsJSON
	}
	
	postBody := bytes.NewBuffer(entryAsJSON)

	return http.Post(address, applicationJSON, postBody)
}

func getLimitID(r http.Request, err error) (*string, *string, error) {
	if err != nil {
		return nil, err
	}

	tokenDetails, errTokenDetails := sessioncookiex.GetSessionCookie(r, nil)

	if err != nil {
		return nil, errCookie
	}
	if cookie == nil {
		return nil, errNoCookie
	}
		
	// previous_minute:current_minute:JWT
	currSeconds := time.Now().Unix()
	prevSeconds := currSeconds - 1
	prevPrevSeconds := prevSeconds - 1

	limitID := fmt.Sprint(prevSeconds, colonDelimiter, currentSeconds, colonDelimiter, cookie.Value)
	prevLimitID := fmt.Sprint(prevPrevSeconds, colonDelimiter, prevSeconds, colonDelimiter, cookie.Value)

	return &limitID, &prevLimitID nil
}

func requestIncrement(limitID *string, err error) (*http.Response, error) {
	if err != nil {
		return nil, err
	}


	instructions := []interface{}{hincrby, limitKeyName, cookie.Value, incrementValue}
	marshaledInstructions, errMarshaledInstructions := json.Marshal(instructions)
	if errMarshaledInstructions != nil {
        return false, errMarshaledInstructions
    }

	requestReader := bytes.NewReader(marshaledInstructions)

	return http.Post(
		details.ConfDetails.CacheAddress,
		ApplicationJson,
		requestReader,
	)
}

func digestCacheResponse() (*int64, error) {
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, err
	}

	var marshaledCount int64
	errMarshalCount := json.NewDecoder(r.Body).Decode(&count)

	return &marshaledCount, errMarshalCount
}

func passRequestThroughLimiter() {
	// if count is over details rate limite
}

func Limit(r *http.Request, err error) (*jwtx.TokenDetails, error) {
	if err != nil {
		return false, err
	}

	limitID, errLimitID := getLimitID(tokenDetails, errTokenDetails)
	resp, errResp := requestIncrement(limitID, errLimitID)
	rBody, errRBody := digestCacheResponse(r, errResp)
	hasPassedLimiter, errHasPassedLimiter := passRequestThroughLimiter(rBody, errRBody)

	if errHasPassedLimiter != nil {
		return nil, errHasPassedLimiter
	}

	return hasPassedLImiter, nil
}