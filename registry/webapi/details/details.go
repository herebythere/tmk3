// brian taylor vann
// details

package details

import (
	"encoding/json"
	"errors"
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

type RegistryDetails struct {
	AvailableServices []string      `json:"available_services"`
	CacheAddress      string        `json:"cache_address"`
	CertPaths         CertPaths     `json:"cert_paths"`
	Config            ConfigDetails `json:"config"`
	Server            ServerDetails `json:"server"`
	ServiceName       string        `json:"service_name"`
}

type KeyDetails struct {
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type UserIdentity struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SkeletonKeyMap = map[string]KeyDetails
type SaltedKeyMap = map[string]passwordx.HashResults
type SkeletonKeyRoleMap = map[string]map[string]int
type AvailableServiceMap = map[string]int

const (
	skeletonKeyDoesNotExist = "SkeletonKey does not exist"
	roleDoesNotExist        = "Role does not exist"
)

var (
	detailsPath         = os.Getenv("CONFIG_FILEPATH")
	configDetailsPath   = os.Getenv("CONFIG_FILEPATH")
	skeletonDetailsPath = os.Getenv("SKELETON_FILEPATH")

	SaltedSkeletonKeys, SkeletonKeyRoles, errSkeletonKeys = parseSkeletonKeys(skeletonDetailsPath)
	ConfDetails, errConfDetails                           = parseConfigDetails(configDetailsPath)
	AvailableServices, errAvailableService                = buildAvailableServicesFromDetails(ConfDetails, errConfDetails)
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

func buildAvailableServicesFromDetails(
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
	for serviceIndex, service := range details.AvailableServices {
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

func VerifySkeletonKey(
	saltedKeys SaltedKeyMap,
	userIdentity *UserIdentity,
	err error,
) (bool, error) {
	if err != nil {
		return false, err
	}

	hashResults, hashResultsExists := saltedKeys[userIdentity.Username]
	if hashResultsExists {
		return passwordx.PasswordIsValid(
			userIdentity.Password,
			&hashResults,
		)
	}

	return false, errors.New(skeletonKeyDoesNotExist)
}

func VerifySkeletonKeyRoleAsAvailableService(
	services AvailableServiceMap,
	roles SkeletonKeyRoleMap,
	userIdentity *UserIdentity,
	role string,
	err error,
) (
	bool,
	error,
) {
	if err != nil {
		return false, err
	}

	roleMap, roleMapExists := roles[userIdentity.Username]
	if !roleMapExists {
		return false, errors.New(roleDoesNotExist)
	}

	_, roleExists := roleMap[role]
	if !roleExists {
		return false, errors.New(roleDoesNotExist)
	}

	_, roleExists = services[role]
	if roleExists {
		return true, nil
	}

	return false, errors.New(roleDoesNotExist)
}
