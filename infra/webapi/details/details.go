// brian taylor vann
// details

package details

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type ConfigDetails struct {
	Filepath     string `json:"filepath"`
	FilepathTest string `json:"filepath_test"`
}

type ServerDetails struct {
	HTTPSPort    int `json:"https_port"`
	IdleTimeout  int `json:"idle_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
}

type CertPaths struct {
	Cert       string `json:"cert"`
	PrivateKey string `json:"private_key"`
}

type CacheDetails struct {
	Host        string        `json:"redis_host"`
	IdleTimeout time.Duration `json:"idle_timeout"`
	MaxActive   int           `json:"max_active"`
	MaxIdle     int           `json:"max_idle"`
	MaxSizeInMB string        `json:"max_size_in_mb"`
	Port        int           `json:"redis_port"`
	Protocol    string        `json:"protocol"`
}

type InfraDetails struct {
	ServiceName string        `json:"service_name"`
	Config      ConfigDetails `json:"config"`
	Cache       CacheDetails  `json:"cache"`
	CertPaths   CertPaths     `json:"cert_paths"`
	Server      ServerDetails `json:"server"`
}

type SkeletonKeyDetails struct {
	Password string   `json:"password"`
	Services []string `json:"services"`
}

type SkeletonKeysDetails = map[string]SkeletonKeyDetails

var (
	detailsPath = os.Getenv("CONFIG_FILEPATH")

	configDetailsPath   = os.Getenv("CONFIG_FILEPATH")
	skeletonDetailsPath = os.Getenv("SKELETON_FILEPATH")

	ConfDetails, errConfDetails         = parseConfigDetails(detailsPath)
	SkeletonDetails, errSkeletonDetails = parseSkeletonKeys(detailsPath)
)

func readFile(path string) (*[]byte, error) {
	detailsJSON, errDetiailsJSON := ioutil.ReadFile(path)
	return &detailsJSON, errDetiailsJSON
}

func parseConfigDetails(path string) (*InfraDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	if errDetailsJSON != nil {
		return nil, errDetailsJSON
	}

	var details InfraDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}

func parseSkeletonKeys(path string) (*SkeletonKeysDetails, error) {
	detailsJSON, errDetailsJSON := readFile(path)
	if errDetailsJSON != nil {
		return nil, errDetailsJSON
	}

	var details SkeletonKeysDetails
	errDetails := json.Unmarshal(*detailsJSON, &details)

	return &details, errDetails
}
