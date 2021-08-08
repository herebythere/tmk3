package person

type PersonCredentialDetails struct {
	Kind     string `json:"kind"`
	Username string `json:"username"`
	Password string `json:"password"`
	Service  string `json:"service"`
}

type PersonIdentity struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


// type PersonCredentialDetails struct {
// 	Kind     string `json:"kind"`
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Service  string `json:"service"`
// }

// type PersonIdentity struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// type ConfirmedPersonIdentity struct {
// 	Kind     string `json:"kind"`
// 	Username string `json:"username"`
// 	Service  string `json:"service"`
// }
