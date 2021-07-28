// brian taylor vann
// details

package details

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"webapi/passwordx"
	"webapi/typeflyweights/person"
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

type RegistryDetails struct {
	AvailableRoles []string      `json:"available_roles"`
	CacheAddress   string        `json:"cache_address"`
	CertPaths      CertPaths     `json:"cert_paths"`
	Config         ConfigDetails `json:"config"`
	Credentials    ConfigDetails `json:"credentials"`
	Server         ServerDetails `json:"server"`
	ServiceName    string        `json:"service_name"`
}

type KeyDetails struct {
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type SkeletonKeyMap = map[string]KeyDetails
type SaltedKeyMap = map[string]passwordx.HashResults
type SkeletonKeyRoleMap = map[string]map[string]int
type AvailableServiceMap = map[string]int


var (
	configDetailsPath   = os.Getenv("CONFIG_FILEPATH")
	skeletonDetailsPath = os.Getenv("SKELETON_KEYS_FILEPATH")
	credentialsDetailsPath = os.Getenv("CREDENTIALS_FILEPATH")

	SaltedSkeletonKeys, SkeletonKeyRoles, errSkeletonKeys = parseSkeletonKeys(skeletonDetailsPath)
	ConfDetails, errConfDetails                           = parseConfigDetails(configDetailsPath)
	Credentials, errCredentials                     	  = parseCredentialDetails(credentialsDetailsPath)
	AvailableRoles, errAvailableRoles                     = buildAvailableRolesFromDetails(ConfDetails, errConfDetails)
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

func buildAvailableRolesFromDetails(
	details *RegistryDetails,
	err error,
) (
	AvailableServiceMap,
	error,
) {
	if err != nil {
		return nil, err
	}
	if details == nil {
		return nil, errors.New("details should not be nil")
	}

	services := AvailableServiceMap{}
	for serviceIndex, service := range details.AvailableRoles {
		services[service] = serviceIndex
	}

	return services, nil
}

func parseSkeletonKeys(path string) (
	SaltedKeyMap,
	SkeletonKeyRoleMap,
	error,
) {
	skeletonKeysJSON, errSkeletonKeysJSON := readFile(path)
	if errSkeletonKeysJSON != nil {
		return nil, nil, errSkeletonKeysJSON
	}

	var skeletonKeys SkeletonKeyMap
	errSkeletonKeys := json.Unmarshal(*skeletonKeysJSON, &skeletonKeys)
	if errSkeletonKeys != nil {
		return nil, nil, errSkeletonKeys
	}

	saltedSkeletonKeys := SaltedKeyMap{}
	SkeletonKeyRoles := SkeletonKeyRoleMap{}

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

		saltedSkeletonKeys[username] = *saltedPassword
		SkeletonKeyRoles[username] = roles
	}

	return saltedSkeletonKeys, SkeletonKeyRoles, errSkeletonKeys
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