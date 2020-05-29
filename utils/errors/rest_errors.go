package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status    int  `json:"code"`
	Error   string `json:"error"`
}

func GetBadRequest(message string) *RestErr {
	return  &RestErr {
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}
