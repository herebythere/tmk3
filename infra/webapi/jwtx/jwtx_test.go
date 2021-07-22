package jwtx

import (
	"fmt"
	"testing"
	"time"
)

const (
	testHeaderBase64 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
)

var (
	testClaims = CreateJWTParams{
		Aud:      []string{"hello", "world"},
		Iss:      "tmk3.com",
		Sub:      "test_jwt",
		Lifetime: 1000000,
	}
)

func TestConvertToBase64(t *testing.T) {
	headerBase, errHeaderBase := convertToBase64(headerDefaultParams)
	if errHeaderBase != nil {
		t.Fail()
		t.Logf(errHeaderBase.Error())
	}

	if *headerBase != testHeaderBase64 {
		t.Fail()
		t.Logf(fmt.Sprint("expected: ", testHeaderBase64, ", instead found: ", *headerBase))
	}
}

func TestGenerateRandomByteArray(t *testing.T) {
	testLength := 12

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

func TestGetNowAsSecond(t *testing.T) {
	oldNow := getNowAsSecond()
	time.Sleep(time.Second)
	nowNow := getNowAsSecond()

	if oldNow >= nowNow {
		t.Fail()
		t.Logf("oldNow should be less than nowNow")
	}
}

func TestGenerateSignature(t *testing.T) {
	payload := "Hello World, this is not a a valid JWT!"
	secret, errSecret := generateRandomByteArray(256, nil)
	if errSecret != nil {
		t.Fail()
		t.Logf(errSecret.Error())
	}

	signature, errSignature := generateSignature(
		headerBase64,
		&payload,
		secret,
		errSecret,
	)

	if errSignature != nil {
		t.Fail()
		t.Logf(errSignature.Error())
	}

	if signature == nil {
		t.Fail()
		t.Logf("signature is nil")
	}
}

func TestCreateJWTClaims(t *testing.T) {
	claims, errClaims := createJWTClaims(&testClaims, nil)
	if claims == nil {
		t.Fail()
		t.Logf("claims should not be nil")
	}

	if errClaims != nil {
		t.Fail()
		t.Logf(errClaims.Error())
	}
}

func TestCreateJWT(t *testing.T) {
	tokenPayload, errTokenPayload := CreateJWT(&testClaims, nil)
	if tokenPayload == nil {
		t.Fail()
		t.Logf("token should not be nil")
	}

	if errTokenPayload != nil {
		t.Fail()
		t.Logf(errTokenPayload.Error())
	}
}

func TestRetrieveTokenChunks(t *testing.T) {
	tokenPayload, errTokenPayload := CreateJWT(&testClaims, nil)
	if tokenPayload == nil {
		t.Fail()
		t.Logf("token should not be nil")
	}

	if errTokenPayload != nil {
		t.Fail()
		t.Logf(errTokenPayload.Error())
	}

	tokenChunks, errTokenChunks := retrieveTokenChunks(&tokenPayload.Token, nil)
	if tokenChunks == nil {
		t.Fail()
		t.Logf("token chunks should not be nil")
	}

	if errTokenChunks != nil {
		t.Fail()
		t.Logf(errTokenChunks.Error())
	}
}

func TestValidateJWT(t *testing.T) {
	tokenPayload, errTokenPayload := CreateJWT(&testClaims, nil)
	if tokenPayload == nil {
		t.Fail()
		t.Logf("token should not be nil")
	}

	if errTokenPayload != nil {
		t.Fail()
		t.Logf(errTokenPayload.Error())
	}

	signatureIsValid, errSignatureIsValid := ValidateJWT(tokenPayload, errTokenPayload)
	if !signatureIsValid {
		t.Fail()
		t.Logf("token is not valid")
	}

	if errSignatureIsValid != nil {
		t.Fail()
		t.Logf(errSignatureIsValid.Error())
	}
}
