//	brian taylor vann
//	gateway

package main

import (
	"fmt"
	"net/http"

	"webapi/details"
	"webapi/mux"
)

var (
	httpPort = fmt.Sprint(":", details.Details.Server.HTTPPort)
)

func main() {
	mux := mux.CreateMux()

	http.ListenAndServe(
		httpPort,
		mux,
	)
}
