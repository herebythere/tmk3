package responses

func WriteError(w http.ResponseWriter, kind string, message string) {
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

func WriteResponse(w http.ResponseWriter, entry interface{}, err error) {
	if err != nil {
		writeError(w, failedToExec, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(contentType, applicationJson)
	if entry != nil {
		json.NewEncoder(w).Encode(entry)
	}
}
