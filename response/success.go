package response

import (
	"encoding/json"
	"net/http"

	"github.com/DuvanM9/go_meta_paginator/meta"
)

type SuccesResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Meta    *meta.Meta  `json:"meta,omitempty"`
}

// Error implements Response.
func (s *SuccesResponse) Error() string {
	return ""
}

// GetBody implements Response.
func (s *SuccesResponse) GetBody() ([]byte, error) {
	return json.Marshal(s)
}

// GetData implements Response.
func (s *SuccesResponse) GetData() interface{} {
	return s.Data
}

// StatusCode implements Response.
func (s *SuccesResponse) StatusCode() int {
	return s.Status
}

func success(msg string, data interface{}, meta *meta.Meta, code int) Response {
	return &SuccesResponse{
		Message: msg,
		Status:  code,
		Data:    data,
		Meta:    meta,
	}
}

func Ok(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusOK)
}

func Created(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusCreated)
}

func Accepted(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusAccepted)
}

func NonAuthoritativeInfo(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusNonAuthoritativeInfo)
}

func NoContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusNoContent)
}

func ResetContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusResetContent)
}

func PartialContent(msg string, data interface{}, meta *meta.Meta) Response {
	return success(msg, data, meta, http.StatusPartialContent)
}
