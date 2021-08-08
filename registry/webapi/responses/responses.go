package responses

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ErrorEntity struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}
type ErrorDeclarations = []ErrorEntity

const (
	applicationJSON = "application/json"
	contentType     = "Content-Type"
	failedToExec    = "failed to exec"
)

func WriteError(w http.ResponseWriter, kind string, message string) {
	setErrors := ErrorDeclarations{
		ErrorEntity{
			Kind:    kind,
			Message: message,
		},
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(setErrors)
}

func WriteResponse(w http.ResponseWriter, entry interface{}, err error) {
	if err != nil {
		WriteError(w, failedToExec, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(contentType, applicationJSON)
	if entry != nil {
		json.NewEncoder(w).Encode(entry)
	}
}
