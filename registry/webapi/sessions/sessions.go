package sessions

import (
	"errors"

	"webapi/details"
	"webapi/jwtx"
)


func writeJWTToCache() {

}

func readJWTFromCache() {
	
}

func CreateGuestSession() {

}

func CreateServerSession(
	params *jwtx.CreateJWTParams,
	err error,
) (
	*TokenPayload,
	error,
) {
	if err != nil {
		return nil, err
	}

	return jtwx.CreateJWT(params, nil)
}

func ValidateSession(
	tokenPayload *TokenPayload,
	err error,
) (bool, error) {
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errors.New("nil token provided")
	}

	return jtwx.ValidateJWT(tokenPayload, nil)
}

