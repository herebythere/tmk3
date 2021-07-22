package jwtx

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Claims struct {
	Aud []string `json:"aud"`
	Exp int64    `json:"exp"`
	Iat int64    `json:"iat"`
	Iss string   `json:"iss"`
	Nbf *int64   `json:"nbf,omitempty"`
	Sub string   `json:"sub"`
}

type CreateJWTParams struct {
	Aud      []string `json:"aud"`
	Iss      string   `json:"iss"`
	Sub      string   `json:"sub"`
	Lifetime int64    `json:"lifetime"`
	Delay    *int64   `json:"delay,omitempty"`
}

type TokenChunks struct {
	Header    string `json:"header"`
	Claims    string `json:"claims"`
	Signature string `json:"signature"`
}

type TokenPayload struct {
	Token  string `json:"token"`
	Secret []byte `json:"secret"`
}

type TokenDetails struct {
	Header Header `json:"header"`
	Claims Claims `json:"claims"`
}

const (
	periodRune = "."
)

var (
	headerDefaultParams = Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	headerBase64, errHeaderBase64 = convertToBase64(headerDefaultParams)
)

func convertToBase64(h interface{}) (*string, error) {
	marhalledHeader, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}

	header64 := base64.RawStdEncoding.EncodeToString(marhalledHeader)

	return &header64, nil
}

// untested
func decodeFromBase64(str *string, err error) (*string, error) {
	if err != nil {
		return nil, err
	}

	data, errData := base64.RawStdEncoding.DecodeString(*str)
	dataStr := string(data)

	return &dataStr, errData
}

func generateRandomByteArray(n int, err error) (*[]byte, error) {
	if err != nil {
		return nil, err
	}

	token := make([]byte, n)
	_, errRandom := rand.Read(token)
	if errRandom != nil {
		return nil, errRandom
	}

	return &token, nil
}

func getNowAsSecond() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

func generateSignature(
	header *string,
	claims *string,
	secret *[]byte,
	err error,
) (*string, error) {
	if err != nil {
		return nil, err
	}

	hmacSecret := hmac.New(sha256.New, *secret)
	headerAndClaims := fmt.Sprint(*header, periodRune, *claims)
	hmacSecret.Write([]byte(headerAndClaims))
	signature := hmacSecret.Sum(nil)

	return convertToBase64(signature)
}

func createJWTClaims(params *CreateJWTParams, err error) (*string, error) {
	if err != nil {
		return nil, err
	}
	if params == nil {
		return nil, errors.New("nil params CreateJWTParams")
	}

	nowAsSecond := getNowAsSecond()
	expiration := nowAsSecond + params.Lifetime

	var notBefore int64
	if params.Delay != nil {
		notBefore = nowAsSecond + *params.Delay
	}

	claims := Claims{
		Aud: params.Aud,
		Exp: expiration,
		Iat: nowAsSecond,
		Iss: params.Iss,
		Nbf: &notBefore,
		Sub: params.Sub,
	}

	return convertToBase64(claims)
}

func retrieveTokenChunks(token *string, err error) (*TokenChunks, error) {
	if err != nil {
		return nil, err
	}

	chunks := strings.Split(*token, ".")
	if len(chunks) != 3 {
		return nil, errors.New("invalid token")
	}

	tokenChunks := TokenChunks{
		Header:    chunks[0],
		Claims:    chunks[1],
		Signature: chunks[2],
	}

	return &tokenChunks, nil
}

// untested
func unmarshalHeader(header *string, err error) (*Header, error) {
	if err != nil {
		return nil, err
	}

	var headerDetails Header
	errHeaderMarshal := json.Unmarshal([]byte(*header), &headerDetails)

	return &headerDetails, errHeaderMarshal
}

// untested
func unmarshalClaims(claims *string, err error) (*Claims, error) {
	if err != nil {
		return nil, err
	}

	var claimsDetails Claims
	errClaimsMarshal := json.Unmarshal([]byte(*claims), &claimsDetails)

	return &claimsDetails, errClaimsMarshal
}

func CreateJWT(params *CreateJWTParams, err error) (*TokenPayload, error) {
	if err != nil {
		return nil, err
	}

	claims, errClaims := createJWTClaims(params, nil)
	secret, errSecret := generateRandomByteArray(128, errClaims)
	signature, errSignature := generateSignature(headerBase64, claims, secret, errSecret)

	token := fmt.Sprint(*headerBase64, periodRune, *claims, periodRune, *signature)
	tokenPayload := TokenPayload{
		Token:  token,
		Secret: *secret,
	}

	return &tokenPayload, errSignature
}

func ValidateJWT(tokenPayload *TokenPayload, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	chunks, errChunks := retrieveTokenChunks(&tokenPayload.Token, nil)
	signature, errSignature := generateSignature(&chunks.Header, &chunks.Claims, &tokenPayload.Secret, errChunks)
	signatureIsValid := *signature == chunks.Signature

	return signatureIsValid, errSignature
}

// untested
func RetrieveTokenDetails(token *string, err error) (*TokenDetails, error) {
	if err != nil {
		return nil, err
	}

	chunks, errChunks := retrieveTokenChunks(token, nil)
	header, errHeader := decodeFromBase64(&chunks.Header, errChunks)
	headerDetails, errHeaderDetails := unmarshalHeader(header, errHeader)
	claims, errClaims := decodeFromBase64(&chunks.Claims, errHeaderDetails)
	claimsDetails, errClaimsDetails := unmarshalClaims(claims, errClaims)

	tokenDetails := TokenDetails{
		Header: *headerDetails,
		Claims: *claimsDetails,
	}

	return &tokenDetails, errClaimsDetails
}
