package details

import (
	"os"
	"testing"
)

var (
	exampleDetailsPath = os.Getenv("CONFIG_FILEPATH_TEST")
	exampleSkeletonKeysPath = os.Getenv("SKELETON_KEYS_FILEPATH_TEST")
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