package skeletonkeysx

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

const (
	testJSONIncrement = "test_json_increment"
	registryTest      = "registry_test"
	testService       = "test_service"
	testUser          = "test_user"
	testFailUser      = "test_failuser"
	testPassword      = "test_password"
	testBadPassword   = "test_badpassword"
)

var (
	testAvailableServicesPath = os.Getenv("AVAILABLE_SERVICES_FILEPATH_TEST")
	testSkeletonKeysPath      = os.Getenv("SKELETON_KEYS_FILEPATH_TEST")
)

func TestPostJSONRequest(t *testing.T) {
	instructions := []interface{}{"INCR", testJSONIncrement}
	resp, errResp := postJSONRequest(instructions)
	if resp == nil {
		t.Fail()
		t.Logf(fmt.Sprint("set service was not successfuul"))
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Fail()
		t.Logf(fmt.Sprint("response code was not:", http.StatusOK))
	}
	if errResp != nil {
		t.Fail()
		t.Logf(errResp.Error())
	}
}

func TestSetAvailableService(t *testing.T) {
	setSuccessful, errSetSuccessful := setAvailableService(registryTest, testService)
	if !setSuccessful {
		t.Fail()
		t.Logf(fmt.Sprint("set service was not successfuul"))
	}
	if errSetSuccessful != nil {
		t.Fail()
		t.Logf(errSetSuccessful.Error())
	}
}

func TestGetAvailableService(t *testing.T) {
	getSuccessful, errGetSuccessful := getAvailableService(registryTest, testService)
	if !getSuccessful {
		t.Fail()
		t.Logf(fmt.Sprint("set service was not successfuul"))
	}
	if errGetSuccessful != nil {
		t.Fail()
		t.Logf(errGetSuccessful.Error())
	}
}

func TestParseAvailableService(t *testing.T) {
	availableServices, errAvailableServices := parseAvailableServicesByFilepath(testAvailableServicesPath)
	if availableServices == nil {
		t.Fail()
		t.Logf(fmt.Sprint("parsing resulted in nil available services"))
	}
	if errAvailableServices != nil {
		t.Fail()
		t.Logf(errAvailableServices.Error())
	}
}

func TestParseAndSetAvailableServices(t *testing.T) {
	errParseAndSet := parseAndSetAvailableServices(registryTest, testAvailableServicesPath, nil)
	if errParseAndSet != nil {
		t.Fail()
		t.Logf(errParseAndSet.Error())
	}
}

func TestSetSkeletonKey(t *testing.T) {
	setSuccessful, errSetSuccessful := setSkeletonKey(registryTest, testUser, testPassword)
	if !setSuccessful {
		t.Fail()
		t.Logf(fmt.Sprint("set skeleton key was not successfuul"))
	}
	if errSetSuccessful != nil {
		t.Fail()
		t.Logf(errSetSuccessful.Error())
	}
}

func TestVerifySkeletonKey(t *testing.T) {
	verified, errVerified := VerifySkeletonKey(registryTest, testUser, testPassword)
	if !verified {
		t.Fail()
		t.Logf(fmt.Sprint("verify skeleton key was not successfuul"))
	}
	if errVerified != nil {
		t.Fail()
		t.Logf(errVerified.Error())
	}
}

func TestBadSkeletonKeyWontVerify(t *testing.T) {
	verified, errVerified := VerifySkeletonKey(registryTest, testFailUser, testPassword)
	if verified {
		t.Fail()
		t.Logf(fmt.Sprint("unidentified skeleton key was successfuul"))
	}
	if errVerified == nil {
		t.Fail()
		t.Logf(fmt.Sprint("an error should have been returned on bad skeleton key"))
	}
}

func TestBadSkeletonKeyPasswordWontVerify(t *testing.T) {
	verified, errVerified := VerifySkeletonKey(registryTest, testUser, testBadPassword)
	if verified {
		t.Fail()
		t.Logf(fmt.Sprint("badd password successfuul"))
	}
	if errVerified != nil {
		t.Fail()
		t.Logf(errVerified.Error())
	}
}

func TestSetSkeletonKeyService(t *testing.T) {
	setSuccessful, errSetSuccessful := setSkeletonKeyService(registryTest, testUser, testService)
	if !setSuccessful {
		t.Fail()
		t.Logf(fmt.Sprint("set service was not successfuul"))
	}
	if errSetSuccessful != nil {
		t.Fail()
		t.Logf(errSetSuccessful.Error())
	}
}

func TestGetSkeletonKeyService(t *testing.T) {
	getSuccessful, errGetSuccessful := getSkeletonKeyService(registryTest, testService)
	if !getSuccessful {
		t.Fail()
		t.Logf(fmt.Sprint("set service was not successfuul"))
	}
	if errGetSuccessful != nil {
		t.Fail()
		t.Logf(errGetSuccessful.Error())
	}
}

func TestParseSkeletonKeys(t *testing.T) {
	skeletonKeys, errParseSkeletonKeys := parseSkeletonKeysByFilepath(testSkeletonKeysPath)
	if skeletonKeys == nil {
		t.Fail()
		t.Logf(fmt.Sprint("parsing resulted in nil skeleton keys"))
	}
	if errParseSkeletonKeys != nil {
		t.Fail()
		t.Logf(errParseSkeletonKeys.Error())
	}
}

func TestParseAndSetSkeletonKeys(t *testing.T) {
	errParseAndSet := parseAndSetSkeletonKeys(registryTest, testSkeletonKeysPath, nil)
	if errParseAndSet != nil {
		t.Fail()
		t.Logf(errParseAndSet.Error())
	}
}

func TestVerifySkeletonKeyAndService(t *testing.T) {
	verified, errVerified := VerifySkeletonKeyAndService(
		registryTest,
		testService,
		testUser,
		testPassword,
	)
	if !verified {
		t.Fail()
		t.Logf(fmt.Sprint("verify skeleton key was not successfuul"))
	}
	if errVerified != nil {
		t.Fail()
		t.Logf(errVerified.Error())
	}
}

func TestSetupSkeletonKeysAndAvailableServices(t *testing.T) {
	errSetupSkeletonKeys := SetupSkeletonKeysAndAvailableServices(
		registryTest,
		testAvailableServicesPath,
		testSkeletonKeysPath,
	)
	if errSetupSkeletonKeys != nil {
		t.Fail()
		t.Logf(errSetupSkeletonKeys.Error())
	}
}
