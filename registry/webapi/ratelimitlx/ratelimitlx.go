package ratelimitlx

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"encoding/json"
	"bytes"
)

const (
	applicationJson = "application/json"
	hset            = "HSET"
	hget            = "HGET"
	increment       = "HINCRBY"

	unableToWriteToCache = "unable to write jwt to cache"
	ratelimits           = "rate_limits"
	colonDelimiter       = ":"

)

var (
	localCacheAddress = os.Getenv("LOCAL_CACHE_ADDRESS")
	errInstructionsAreNil = errors.New("instructions are nil")
)

// create session
func getCacheSetID(categories ...string) string {
	return strings.Join(categories, colonDelimiter)
}

func parseCachedInteger(resp *http.Response) (*int64, error) {
	var countAsBase64 int64
	errJSONResponse := json.NewDecoder(resp.Body).Decode(&countAsBase64)
	if errJSONResponse != nil {
		return nil, errJSONResponse
	}

	// serviceAsBytes, errServiceAsBytes := base64.URLEncoding.DecodeString(
	// 	serviceAsBase64,
	// )
	// if errServiceAsBytes != nil {

	// 	return nil, errServiceAsBytes
	// }

	// serviceReturned := string(serviceAsBytes)
	return &countAsBase64, nil
}

func postJSONRequest(
	instructions []interface{},
) (
	*http.Response,
	error,
) {
	if instructions == nil {
		return nil, errInstructionsAreNil
	}

	instructionsAsJSON, errInstructionsAsJSON := json.Marshal(instructions)
	if errInstructionsAsJSON != nil {
		return nil, errInstructionsAsJSON
	}

	requestBody := bytes.NewBuffer(instructionsAsJSON)

	return http.Post(localCacheAddress, applicationJson, requestBody)
}

func minInt(x, y int64) int64 {
    if x > y {
        return y
    }
    return x
}

func passesSlidingWindowLimit(
	currentTime int64,
	interval int64,
	prevCount int64,
	currCount int64,
	limit int64,
) bool {
	if currCount > limit {
		return false
	}

	adjPrevCount = minInt(prevCount, limit)
	adjCurrCount = minInt(currCount, limit)

	intervalValue := currentTime % interval
	intervalDelta := 1 - (float64(intervalValue) / float64(interval))

	windowValue := int64(intervalDelta * float64(adjPrevCount))
	totalCount := windowValue + adjCurrCount

	if totalCount < limit {
		return true
	}

	return false
}

func getIntervals(
	serverName string,
	interval int64, // by seconds
) (*string, *string) {
	currentTime := time.Now().Unix()
	currentInterval := currentTime / interval
	previousInterval := currentInterval - 1
	
	setID := getCacheSetID(serverName, ratelimits)

	prevIntervalID := getCacheSetID(
		r.RemoteAddress,
		previousInterval,
	)
	currIntervalID := getCacheSetID(
		r.RemoteAddress,
		currentInterval,
	)

	return &prevIntervalID, &currIntervalID
}

func LimitByIP(
	r *http.Request,
	serverName string,
	interval int64, // by seconds
	limit int64,
) (bool, error) {
	prevIntervalID, currIntervalID := getIntervals(serverName, interval)

	incrPrevInstructions := []interface{}{hget, setID, prevIntervalID}
	incrInstructions := []interface{}{increment, setID, intervalID}

	// set cache

	// get cache


	// get cache name
}

// rate limit by session

// (r, serverName, cookieName)
func LimitByID(
	uniqueID *string,
	serverName string,
	interval int64, // by seconds
	limit int64,
) (bool, error) {
	prevIntervalID, currIntervalID := getIntervals(serverName, interval)

}