//	brian taylor vann

package mux

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type MailRequestBody struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

type MailResponseBody struct {
	CmdOutput string `json:"cmd_output"`
}

type ErrorEntity struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

type ErrorDeclarations = []ErrorEntity

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
	execRoute       = "/"
	post            = "POST"
	failedToExec    = "failed to exec mail command"
	mail            = "mail"
	mailSubject     = "-s"
	quotationMark   = "\""
)

var (
	errMethodMessage  = errors.New("request method is not POST")
	errNilRequestBody = errors.New("request body is nil")
)

func validPost(r *http.Request) error {
	if r.Method == post {
		return nil
	}

	return errMethodMessage
}

func getBody(r *http.Request, err error) (*MailRequestBody, error) {
	if err != nil {
		return nil, err
	}

	if r.Body == nil {
		return nil, errNilRequestBody
	}

	var rBody MailRequestBody
	errRBody := json.NewDecoder(r.Body).Decode(&rBody)

	return &rBody, errRBody
}

func execMailCommand(rBody *MailRequestBody, err error) (*string, error) {
	if err != nil {
		return nil, err
	}

	subject := fmt.Sprint(quotationMark, rBody.Subject, quotationMark)
	cmd := exec.Command(mail, mailSubject, subject, rBody.Recipient)
	cmd.Stdin = strings.NewReader(rBody.Body)
	results, errResults := cmd.Output()
	if errResults != nil {
		return nil, errResults
	}

	resultsAsStr := string(results)

	return &resultsAsStr, errResults
}

func writeError(w http.ResponseWriter, kind string, err error) {
	if err == nil {
		return
	}

	setErrors := ErrorDeclarations{
		ErrorEntity{
			Kind:    kind,
			Message: err.Error(),
		},
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set(contentType, applicationJson)
	json.NewEncoder(w).Encode(setErrors)
}

func writeResponse(w http.ResponseWriter, cmdResult *string, err error) error {
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set(contentType, applicationJson)

	return err
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	errPost := validPost(r)
	rBody, errRBody := getBody(r, errPost)
	cmdResult, errCmdResult := execMailCommand(rBody, errRBody)
	errResponse := writeResponse(w, cmdResult, errCmdResult)
	writeError(w, failedToExec, errResponse)
}

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(execRoute, sendMail)

	return mux
}
