package infrax


// ping
//	-> success means server is alive, proceed to register

// register rate limit 3 times
// register service

// user password
// password
// service name
// self-declared uuid
// service kind: broadcaster | receiver
// 
// assigned a name? db.b.1.tmk3.com db.r.0.tmk3.com

// when a new server registers

// -> infer ip address from request (forwarded from)
//	-> returns server session cookie, assigned name


// update service details
// payload
//   uuid
//   session
//	 -> infre ip from request (forwarded from)






// example cache service
//
//
// starts up
// pings registry until it is successful
// registers service
// then pings every 15seconds to update ip address

// when denied
// kick back to ping and start over
