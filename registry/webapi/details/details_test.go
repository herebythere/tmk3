package details

import (
	"encoding/json"
	"os"
	"testing"

	"fmt"
)

var (
	exampleDetailsPath                                     = os.Getenv("CONFIG_FILEPATH_TEST")
	exampleSkeletonKeysPath                                = os.Getenv("SKELETON_KEYS_FILEPATH_TEST")
	testSaltedUsers, testUserRoles, errTestSkeletonDetails = parseSkeletonKeys(exampleSkeletonKeysPath)
	testSkeletonKeysJSON, errTestSkeletonKeysJSON          = readFile(exampleSkeletonKeysPath)
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

func TestParseDetails(t *testing.T) {
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

func TestParseSkeletonKeys(t *testing.T) {
	testSaltedUsers, testUserRoles, errTestSkeletonDetails := parseSkeletonKeys(exampleSkeletonKeysPath)

	if testSaltedUsers == nil {
		t.Fail()
		t.Logf("There should be salted users that can be parsed")
	}

	if testUserRoles == nil {
		t.Fail()
		t.Logf("There should be user roles that can be parsed")
	}

	if errTestSkeletonDetails != nil {
		t.Fail()
		t.Logf(errTestSkeletonDetails.Error())
	}
}

func TestBuildAvailableServicesFromDetails(t *testing.T) {
	exampleDetails, errExampleDetails := parseConfigDetails(exampleDetailsPath)
	services, errServices := buildAvailableServicesFromDetails(
		exampleDetails,
		errExampleDetails,
	)

	if errServices != nil {
		t.Fail()
		t.Logf(errServices.Error())
	}

	if services == nil {
		t.Fail()
		t.Logf("There should be services that can be parsed")
	}
}

func TestVerifySkeletonKey(t *testing.T) {
	var testSkeletonKeys SkeletonKeyMap
	errTestSkeletonKeys := json.Unmarshal(*testSkeletonKeysJSON, &testSkeletonKeys)
	if errTestSkeletonKeys != nil {
		t.Fail()
		t.Logf(errTestSkeletonKeys.Error())
	}

	for user, details := range testSkeletonKeys {
		userIdentity := UserIdentity{
			Username: user,
			Password: details.Password,
		}
		skeletonKeyIsVerified, errVerifySkeletonKey := VerifySkeletonKey(testSaltedUsers, &userIdentity, nil)
		if errVerifySkeletonKey != nil {
			t.Fail()
			t.Logf(errVerifySkeletonKey.Error())
		}

		if !skeletonKeyIsVerified {
			t.Fail()
			t.Logf("Skeleton Key was not valid")
		}
	}
}

func TestVerifySkeletonKeyRoleAsAvailableService(t *testing.T) {
	exampleDetails, errExampleDetails := parseConfigDetails(exampleDetailsPath)
	exampleServices, errExampleServices := buildAvailableServicesFromDetails(
		exampleDetails,
		errExampleDetails,
	)
	if errExampleServices != nil {
		t.Fail()
		t.Logf(errExampleServices.Error())
	}

	var testSkeletonKeys SkeletonKeyMap
	errTestSkeletonKeys := json.Unmarshal(*testSkeletonKeysJSON, &testSkeletonKeys)
	if errTestSkeletonKeys != nil {
		t.Fail()
		t.Logf(errTestSkeletonKeys.Error())
	}

	for user, details := range testSkeletonKeys {
		userIdentity := UserIdentity{
			Username: user,
			Password: details.Password,
		}

		for roleIndex, role := range details.Roles {
			roleExists, errRoleExists := VerifySkeletonKeyRoleAsAvailableService(
				exampleServices,
				testUserRoles,
				&userIdentity,
				role,
				nil,
			)
			if errRoleExists != nil {
				t.Fail()
				t.Logf(errRoleExists.Error())
			}

			if !roleExists {
				t.Fail()
				t.Logf(fmt.Sprint(roleIndex, role, roleDoesNotExist))
			}
		}
	}
}
