package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Error implements Response.
func (e *ErrorResponse) Error() string {
	return e.Message
}

// GetBody implements Response.
func (e *ErrorResponse) GetBody() ([]byte, error) {
	return json.Marshal(e)
}

// GetData implements Response.
func (e *ErrorResponse) GetData() interface{} {
	return nil
}

// StatusCode implements Response.
func (e *ErrorResponse) StatusCode() int {
	return e.Status
}

func erros(msg string, status int) Response {
	return &ErrorResponse{
		Status:  status,
		Message: msg,
	}
}

func InternalServerError(msg string) Response {
	return erros(msg, http.StatusInternalServerError)
}

func NotFound(msg string) Response {
	return erros(msg, http.StatusNotFound)
}

func Unauthorized(msg string) Response {
	return erros(msg, http.StatusUnauthorized)
}

func Forbidden(msg string) Response {
	return erros(msg, http.StatusForbidden)
}

func BadRequest(msg string) Response {
	return erros(msg, http.StatusBadRequest)
}

func Conflict(msg string) Response {
	return erros(msg, http.StatusConflict)
}
