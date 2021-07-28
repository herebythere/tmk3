package verifyperson

import (
	"errors"
	"net/http"

	"webapi/details"
	"webapi/passwordx"
	"webapi/typeflyweights/person"
)

const (
	skeletonKeyDoesNotExist = "SkeletonKey does not exist"
	roleDoesNotExist        = "Role does not exist"
	roleIsNotAuthentic      = "Role is not authentic"
	unableToVerifyPerson	= "unable to verify person or role"
	unknownCredentialKind	= "unknown credential kind found"

	RegistryHost = "registry_host"
	SkeletonKey = "skeleton_key"
)

func VerifySkeletonKey(
	personIdentity *person.PersonIdentity,
	err error,
) (bool, error) {
	if err != nil {
		return false, err
	}

	hashResults, hashResultsExists := details.SaltedSkeletonKeys[personIdentity.Username]
	if hashResultsExists {
		return passwordx.PasswordIsValid(
			personIdentity.Password,
			&hashResults,
		)
	}

	return false, errors.New(skeletonKeyDoesNotExist)
}

func VerifySkeletonKeyRoleAsAvailableService(
	personIdentity *person.PersonIdentity,
	role string,
	err error,
) (
	bool,
	error,
) {
	if err != nil {
		return false, err
	}

	roleMap, roleMapExists := details.SkeletonKeyRoles[personIdentity.Username]
	if !roleMapExists {
		return false, errors.New(roleDoesNotExist)
	}
	_, roleExists := roleMap[role]
	if !roleExists {
		return false, errors.New(roleDoesNotExist)
	}

	_, roleExists = details.AvailableRoles[role]
	if roleExists {
		return true, nil
	}

	return false, errors.New(roleDoesNotExist)
}

func VerifyPersonAndRoleWithSkeletonKeys(
	personIdentity *person.PersonIdentity,
	role string,
	err error,
) (
	bool,
	error,
) {
	roleIsAuthentic, errVerifyUser := VerifySkeletonKeyRoleAsAvailableService(
		personIdentity,
		role,
		nil,
	)
	if !roleIsAuthentic {
		return false, errVerifyUser
	}

	return VerifySkeletonKey(personIdentity, err)
}

func VerifyPersonAndRole(
	credentialDetails *person.PersonCredentialDetails,
	role string,
	err error,
) (
	bool,
	error,
) {
	personIdentity := person.PersonIdentity{
		Username: credentialDetails.Username,
		Password: credentialDetails.Password,
	}

	if credentialDetails.Kind == SkeletonKey {
		return VerifyPersonAndRoleWithSkeletonKeys(
			&personIdentity,
			role,
			nil,
		)
	}

	// TODO: try verify through users remotely

	return false, errors.New(unknownCredentialKind)
}

// get request body
// railway style
func VerifyPersonAndRoleHandler(
	r *http.Request,
	err error,
) (
	bool,
	error,
) {
	if err != nil {
		return false, err
	}

	// get guest session
	
	// get body

	// call VerifyPersonAndRole

	// write response

	// write errors

	return false, errors.New(unableToVerifyPerson)
}
