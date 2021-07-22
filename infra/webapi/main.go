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
	proxyMux := mux.CreateMux()

	err := http.ListenAndServeTLS(
		httpsPort,
		certFilepath,
		keyFilepath,
		proxyMux,
	)

	log.Fatal(err)
}
