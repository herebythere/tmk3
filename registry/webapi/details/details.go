package details

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"webapi/typeflyweights/person"

	snowprintx "github.com/herebythere/snowprintx/v0.1/golang"
)

type ConfigDetails struct {
	Filepath     string `json:"filepath"`
	FilepathTest string `json:"filepath_test"`
}

type ServerDetails struct {
	CacheAddress       string `json:"cache_address"`
	GuestCookieLabel   string `json:"guest_cookie_label"`
	HTTPSPort          int64  `json:"https_port"`
	IdleTimeout        int64  `json:"idle_timeout"`
	RateLimit          int64  `json:"rate_limit"`
	ReadTimeout        int64  `json:"read_timeout"`
	SessionCookieLabel string `json:"session_cookie_label"`
	WriteTimeout       int64  `json:"write_timeout"`
}

type CertPaths struct {
	Cert       string `json:"cert"`
	PrivateKey string `json:"private_key"`
}

type RegistryDetails struct {
	CertPaths         CertPaths     `json:"cert_paths"`
	AvailableServices ConfigDetails `json:"available_services"`
	Config            ConfigDetails `json:"config"`
	Credentials       ConfigDetails `json:"credentials"`
	Server            ServerDetails `json:"server"`
	ServiceName       string        `json:"service_name"`
}

var (
	configDetailsPath      = os.Getenv("CONFIG_FILEPATH")
	credentialsDetailsPath = os.Getenv("CREDENTIALS_FILEPATH")

	Snowprint, errSnowprint = snowprintx.CreateSnowprint(
		ConfDetails.ServiceName,
	)

	ConfDetails, errConfDetails = parseConfigDetails(configDetailsPath)
	Credentials, errCredentials = parseCredentialDetails(credentialsDetailsPath)
)

func readFile(path string) (*[]byte, error) {
	detailsJSON, errDetiailsJSON := ioutil.ReadFile(path)
	return &detailsJSON, errDetiailsJSON
}

func parseConfigDetails(path string) (*RegistryDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	if errDetailsJSON != nil {
		return nil, errDetailsJSON
	}

	var details RegistryDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}

func parseCredentialDetails(path string) (*person.PersonCredentialDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	if errDetailsJSON != nil {
		return nil, errDetailsJSON
	}

	var details person.PersonCredentialDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}
