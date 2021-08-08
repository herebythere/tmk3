package registryhanlders

import (
	"net/http"

	"webapi/responses"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	// rate limit
	//

	responses.WriteResponse(h, nil, nil)
}

func RequestGuestSession(w http.ResponseWriter, r *http.Request) {
	// rate limit
	//
	// return okay! easy

	// ping
	//	-> success means server is alive, proceed to register
}

func RegisterServer(w http.ResponseWriter, r *http.Request) {
	// register rate limit 3 times
	// register service

	// -> infer ip address from request (forwarded from)

	// user password
	// password
	// service name
	// self-declared uuid
	// service kind: broadcaster | receiver
	//
	// assigned a name? db.b.1.tmk3.com db.r.0.tmk3.com
}

func UpdateServerDetails(w http.ResponseWriter, r *http.Request) {
	// verifiy session
}

func RequestServiceDetails() {
	// verifiy session

	// -> infer ip address from request (forwarded from)

	// update service details
	// payload
	//   uuid
	//   session
	//	 -> infre ip from request (forwarded from)
}

// when a new server registers

//	-> returns server session cookie, assigned name

// example cache service
//
//
// starts up
// pings registry until it is successful
// registers service
// then pings every 15seconds to update ip address

// when denied
// kick back to ping and start over
