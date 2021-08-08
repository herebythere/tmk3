package details

import (
	"os"
	"testing"
)

var (
	exampleDetailsPath      = os.Getenv("CONFIG_FILEPATH_TEST")
	exampleSkeletonKeysPath = os.Getenv("SKELETON_KEYS_FILEPATH_TEST")
	exampleCredentialsPath  = os.Getenv("CREDENTIALS_FILEPATH_TEST")

	// testSaltedUsers, testUserRoles, errTestSkeletonDetails = parseSkeletonKeys(exampleSkeletonKeysPath)
	// testSkeletonKeysJSON, errTestSkeletonKeysJSON          = readFile(exampleSkeletonKeysPath)
)

func TestReadFile(t *testing.T) {
	currentFile, errCurrentFile := readFile(exampleDetailsPath)
	if currentFile == nil {
		t.Fail()
		t.Logf("There should be an example init file available.")
	}

	if errCurrentFile != nil {
		t.Fail()
		t.Logf(errCurrentFile.Error())
	}
}

func TestParseConfigDetails(t *testing.T) {
	exampleDetails, errExampleDetails := parseConfigDetails(exampleDetailsPath)

	if exampleDetails == nil {
		t.Fail()
		t.Logf("There should be details that can be parsed")
	}

	if errExampleDetails != nil {
		t.Fail()
		t.Logf(errExampleDetails.Error())
	}
}

func TestParseCredentialDetails(t *testing.T) {
	credentials, errCredentials := parseCredentialDetails(exampleCredentialsPath)
	if credentials == nil {
		t.Fail()
		t.Logf("There should be credentials that can be parsed")
	}
	if errCredentials != nil {
		t.Fail()
		t.Logf(errCredentials.Error())
	}
}

// func TestParseSkeletonKeys(t *testing.T) {
// 	testSaltedUsers, testUserRoles, errTestSkeletonDetails := parseSkeletonKeys(exampleSkeletonKeysPath)

// 	if testSaltedUsers == nil {
// 		t.Fail()
// 		t.Logf("There should be salted users that can be parsed")
// 	}

// 	if testUserRoles == nil {
// 		t.Fail()
// 		t.Logf("There should be user roles that can be parsed")
// 	}

// 	if errTestSkeletonDetails != nil {
// 		t.Fail()
// 		t.Logf(errTestSkeletonDetails.Error())
// 	}
// }

// func TestBuildAvailableServicesFromDetails(t *testing.T) {
// 	exampleDetails, errExampleDetails := parseConfigDetails(exampleDetailsPath)
// 	services, errServices := buildAvailableServicesFromDetails(
// 		exampleDetails,
// 		errExampleDetails,
// 	)

// 	if errServices != nil {
// 		t.Fail()
// 		t.Logf(errServices.Error())
// 	}

// 	if services == nil {
// 		t.Fail()
// 		t.Logf("There should be services that can be parsed")
// 	}
// }
