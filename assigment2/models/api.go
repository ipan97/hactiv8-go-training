package models

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
