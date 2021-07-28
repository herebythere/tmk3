package sessions

import (
	"webapi/jwtx"
)

// jwt payload

// we get a request
// parse the session in the request
//

// json parse body

// create session from details
// -> save to cache, return session
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
) {
	if err != nil {
		return nil, err
	}

	return jtwx.ValidateJWT(params, nil)
}

// validate session
// -> request from cache, return ok not okay
