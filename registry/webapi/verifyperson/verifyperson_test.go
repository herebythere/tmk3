package verifyperson

import (
	"fmt"
	"os"
	"testing"

	"webapi/details"
	"webapi/typeflyweights/person"
)

var (
	credentialsPath = os.Getenv("CREDENTIALS_FILEPATH")
)

func TestVerifySkeletonKey(t *testing.T) {
	personIdentity := person.PersonIdentity{
		Username: details.Credentials.Username,
		Password: details.Credentials.Password,
	}

	skeletonKeyIsVerified, errVerifySkeletonKey := VerifySkeletonKey(&personIdentity, nil)
	if errVerifySkeletonKey != nil {
		t.Fail()
		t.Logf(errVerifySkeletonKey.Error())
	}

	if !skeletonKeyIsVerified {
		t.Fail()
		t.Logf("Skeleton Key was not valid")
	}
}

func TestVerifySkeletonKeyRoleAsAvailableService(t *testing.T) {
	personIdentity := person.PersonIdentity{
		Username: details.Credentials.Username,
		Password: details.Credentials.Password,
	}

	roleExists, errRoleExists := VerifySkeletonKeyRoleAsAvailableService(
		&personIdentity,
		RegistryHost,
		nil,
	)
	if errRoleExists != nil {
		t.Fail()
		t.Logf(errRoleExists.Error())
	}

	if !roleExists {
		t.Fail()
		t.Logf(fmt.Sprint(roleDoesNotExist, ": ", RegistryHost))
	}
}

func TestVerifyPersonAndRoleWithSkeletonKeys(t *testing.T) {
	personIdentity := person.PersonIdentity{
		Username: details.Credentials.Username,
		Password: details.Credentials.Password,
	}

	roleExists, errRoleExists := VerifyPersonAndRoleWithSkeletonKeys(
		&personIdentity,
		RegistryHost,
		nil,
	)
	if errRoleExists != nil {
		t.Fail()
		t.Logf(errRoleExists.Error())
	}

	if !roleExists {
		t.Fail()
		t.Logf(fmt.Sprint(roleDoesNotExist, ": ", RegistryHost))
	}
}

func TestVerifyPersonAndRoleThroughSkeletonKey(t *testing.T) {
	roleExists, errRoleExists := VerifyPersonAndRole(
		details.Credentials,
		RegistryHost,
		nil,
	)
	if errRoleExists != nil {
		t.Fail()
		t.Logf(errRoleExists.Error())
	}

	if !roleExists {
		t.Fail()
		t.Logf(fmt.Sprint(roleDoesNotExist, ": ", RegistryHost))
	}
}

// TODO: test verify through users remotely
