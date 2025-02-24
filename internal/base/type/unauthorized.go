package basetype

import "net/http"

type UnauthorizedResponse struct {
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}

type ForbiddenAccessResponse struct {
	Message    string `json:"message"`
	StatusCode uint   `json:"statusCode"`
}

func NewUnauthorizedResponse() UnauthorizedResponse {
	return UnauthorizedResponse{
		Message:    "Unauthorized User",
		StatusCode: http.StatusUnauthorized,
	}
}

func NewForbiddenAccessResponse() ForbiddenAccessResponse {
	return ForbiddenAccessResponse{
		Message:    "Forbidden Access",
		StatusCode: http.StatusForbidden,
	}
}
