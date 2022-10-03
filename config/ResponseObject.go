package config

import (
	"net/http"
	"time"
)

type ResponseBuilder interface {
	IsSuccess_(bool) ResponseBuilder
	Message_(string) ResponseBuilder
	Body_(any) ResponseBuilder
	Build() any
}

type ResponseObject struct {
	TimeStamp  time.Time
	Message    string
	IsSuccess  bool
	StatusCode uint64
	Body       any
}

func (cb *ResponseObject) Build() any {
	return &ResponseObject{
		time.Now(),
		cb.Message,
		cb.IsSuccess,
		cb.StatusCode,
		cb.Body,
	}
}

func (cb *ResponseObject) IsSuccess_(isSuccess bool) ResponseBuilder {
	cb.IsSuccess = isSuccess
	cb.StatusCode = http.StatusBadRequest

	if isSuccess {
		cb.StatusCode = http.StatusOK
		cb.Message = "Success"
	}

	return cb
}

func (cb *ResponseObject) Message_(message string) ResponseBuilder {
	cb.Message = message
	return cb
}

func (cb *ResponseObject) Body_(body any) ResponseBuilder {
	cb.Body = body
	return cb
}

func Builder() ResponseBuilder {
	return &ResponseObject{}
}
