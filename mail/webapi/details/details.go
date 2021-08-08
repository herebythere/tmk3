// brian taylor vann
// details

package details

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigDetails struct {
	Filepath     string `json:"filepath"`
	FilepathTest string `json:"filepath_test"`
}

type ServerDetails struct {
	HTTPPort     int `json:"http_port"`
	MXPort		 int `json:"mx_port"`
	IdleTimeout  int `json:"idle_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
}

type CertsDetails struct {
	CertFilePath       string `json:"cert_filepath"`
	PrivateKeyFilePath string `json:"private_key_filepath"`
}

type MailDetails struct {
	ServiceName string            `json:"service_name"`
	TLD			string            `json:"tld"`
	Config      ConfigDetails     `json:"config"`
	Certs	    CertsDetails	  `json:"certs"`
	Server      ServerDetails     `json:"server"`
}

var (
	detailsPath = os.Getenv("CONFIG_FILEPATH")
	Details, DetailsErr = ReadDetailsFromFile(detailsPath)
)

func readFile(path string) (*[]byte, error) {
	detailsJSON, errDetiailsJSON := ioutil.ReadFile(path)
	return &detailsJSON, errDetiailsJSON
}

func parseDetails(detailsJSON *[]byte, err error) (*MailDetails, error) {
	if err != nil {
		return nil, err
	}

	var details MailDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}

func ReadDetailsFromFile(path string) (*MailDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	return parseDetails(detailsJSON, errDetailsJSON)
}
