package person

type PersonCredentialDetails struct {
	Kind     string `json:"kind"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type PersonIdentity struct {
	Username string `json:"username"`
	Password string `json:"password"`
}