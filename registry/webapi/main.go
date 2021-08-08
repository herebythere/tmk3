package main

import (
	"fmt"
	"log"
	"net/http"

	"webapi/details"
	"webapi/mux"

	skeletonkeyslx "github.com/herebythere/skeletonkeyslx/v0.1/golang"
)

var (
	httpsPort = fmt.Sprint(":", details.ConfDetails.Server.HTTPSPort)
)

func main() {
	// Setup Skeleton Keys
	errSkeletonKeys := skeletonkeyslx.SetupSkeletonKeysAndAvailableServices(
		details.ConfDetails.ServiceName,
		details.ConfDetails.AvailableServices.Filepath,
		details.ConfDetails.SkeletonKeys.Filepath,
	)
	if errSkeletonKeys != nil {
		log.Fatal(errSkeletonKeys.Error())
	}

	// start server
	proxyMux := mux.CreateMux()
	errServer := http.ListenAndServeTLS(
		httpsPort,
		details.ConfDetails.CertPaths.Cert,
		details.ConfDetails.CertPaths.PrivateKey,
		proxyMux,
	)

	log.Fatal(errServer.Error())
}
