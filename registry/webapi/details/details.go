// brian taylor vann
// details

package details

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"webapi/passwordx"
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

type InfraDetails struct {
	CacheAddress	string			`json:"cache_address"`
	CertPaths		CertPaths		`json:"cert_paths"`
	Config      	ConfigDetails	`json:"config"`
	Server			ServerDetails	`json:"server"`
	ServiceName		string			`json:"service_name"`
}

type KeyDetails struct {
	Password	string		`json:"password"`
	Roles		[]string	`json:"roles"`
}

type SkeletonKeysDetails = map[string]KeyDetails
type SaltedUserDetails = map[string]*passwordx.HashResults
type UserRoleDetails = map[string]*map[string]int

var (
	detailsPath = os.Getenv("CONFIG_FILEPATH")

	configDetailsPath   = os.Getenv("CONFIG_FILEPATH")
	skeletonDetailsPath = os.Getenv("SKELETON_FILEPATH")

	ConfDetails, errConfDetails = parseConfigDetails(configDetailsPath)
	SaltedUsers, UserRoles, errSkeletonKeys = parseSkeletonKeys(skeletonDetailsPath)
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

func parseSkeletonKeys(path string) (
	*SaltedUserDetails,
	*UserRoleDetails,
	error,
) {
	skeletonKeysJSON, errSkeletonKeysJSON := readFile(path)
	if errSkeletonKeysJSON != nil {
		return nil, nil, errSkeletonKeysJSON
	}

	var skeletonKeys SkeletonKeysDetails
	errSkeletonKeys := json.Unmarshal(*skeletonKeysJSON, &skeletonKeys)
	if errSkeletonKeys != nil {
		return nil, nil, errSkeletonKeys
	}

	saltedUsers := SaltedUserDetails{}
	userRoles := UserRoleDetails{}

	for username, details := range skeletonKeys {
		saltedPassword, errSaltedPassword := passwordx.HashPassword(
			details.Password,
			&passwordx.DefaultHashParams,
		)

		if errSaltedPassword != nil {
			continue
		}

		roles := map[string]int{}
		for roleIndex, role := range details.Roles {
			roles[role] = roleIndex
		}

		saltedUsers[username] = saltedPassword
		userRoles[username] = &roles
	}

	return &saltedUsers, &userRoles, errSkeletonKeys
}
