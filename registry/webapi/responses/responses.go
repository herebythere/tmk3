package responses

func writeError(w http.ResponseWriter, kind string, message string) {
	setErrors := ErrorDeclarations{
		ErrorEntity{
			Kind:    kind,
			Message: message,
		},
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set(contentType, applicationJson)
	json.NewEncoder(w).Encode(setErrors)
}

func writeResponse(w http.ResponseWriter, entry interface{}, err error) {
	if err != nil {
		writeError(w, failedToExec, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(contentType, applicationJson)
	json.NewEncoder(w).Encode(entry)
}
