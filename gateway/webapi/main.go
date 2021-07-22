package main

import (
	"fmt"
	"net/http"

	"webapi/details"
	"webapi/muxrouter"
)

var (
	httpsPort    = fmt.Sprint(":", details.Details.Server.HTTPSPort)
	certFilepath = details.Details.CertPaths.Cert
	keyFilepath  = details.Details.CertPaths.PrivateKey
)

func main() {
	proxyMux, errProxyMux := muxrouter.CreateHTTPSMux(&details.Details.Routes)
	if errProxyMux != nil {
		return
	}

	http.ListenAndServeTLS(
		httpsPort,
		certFilepath,
		keyFilepath,
		proxyMux,
	)
}
