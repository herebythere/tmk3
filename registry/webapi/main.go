package main

import (
	"fmt"
	"log"
	"net/http"

	"webapi/details"
	"webapi/mux"
)

var (
	httpsPort    = fmt.Sprint(":", details.ConfDetails.Server.HTTPSPort)
	certFilepath = details.ConfDetails.CertPaths.Cert
	keyFilepath  = details.ConfDetails.CertPaths.PrivateKey
)

func main() {
	// verify user
	//   ->
	//
	//

	proxyMux := mux.CreateMux()

	errServer := http.ListenAndServeTLS(
		httpsPort,
		certFilepath,
		keyFilepath,
		proxyMux,
	)

	log.Fatal(errServer)
}
