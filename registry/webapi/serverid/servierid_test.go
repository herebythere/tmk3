package serverid

import (
	"fmt"
	"testing"
	// "time"
)


func TestGenerateRandomByteArray(t *testing.T) {
	testLength := 128

	randomBytes, errRandomBytes := generateRandomByteArray(testLength, nil)
	if errRandomBytes != nil {
		t.Fail()
		t.Logf(errRandomBytes.Error())
	}

	if randomBytes == nil {
		t.Fail()
		t.Logf("randomBytes should not be nil")
	}

	randomByteLength := len(*randomBytes)

	if randomByteLength != testLength {
		t.Fail()
		t.Logf(
			fmt.Sprint(
				"randomBytes should be:",
				testLength,
				", instead found:",
				randomByteLength,
			),
		)
	}
}

func TestCreateServerID(t *testing.T) {
	potentialServerID, errPotentialServerID := createServerID(16)
	if errPotentialServerID != nil {
		t.Fail()
		t.Logf(errPotentialServerID.Error())
	}
	if potentialServerID == nil {
		t.Fail()
		t.Logf("potentialServerID should not be nil")
	}
}