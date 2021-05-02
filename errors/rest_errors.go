package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	M string        `json:"message"`
	S int           `json:"status"`
	E string        `json:"error"`
	C []interface{} `json:"causes"`
}

func (r *restErr) Message() string {
	return r.M
}

func (r *restErr) Status() int {
	return r.S
}

func (r *restErr) Error() string {
	return r.E
}

func (r *restErr) Causes() []interface{} {
	return r.C
}

func NewRestError(message string, status int, error string, causes []interface{}) RestErr {
	return &restErr{
		M: message,
		S: status,
		E: error,
		C: causes,
	}
}

func NewRestErrorFromBytes(b []byte) (RestErr, error) {
	var r restErr
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, errors.New("invalid json")
	}
	return &r, nil
}

func NewBadRequestError(m string) RestErr {
	return &restErr{
		M: m,
		S: http.StatusBadRequest,
		E: "Bad Request",
	}
}

func NewNotFoundError(m string) RestErr {
	return &restErr{
		M: m,
		S: http.StatusNotFound,
		E: "Not found",
	}
}

func NewInternalServerError(m string, err error) RestErr {
	r := &restErr{
		M: m,
		S: http.StatusInternalServerError,
		E: "Internal server error",
	}
	if err != nil {
		r.C = append(r.C, err.Error())
	}
	return r
}

func NewUnauthorizedError(m string) RestErr {
	return &restErr{
		M: m,
		S: 401,
		E: "Unauthorized",
	}
}
