package helpers

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code,omitempty"`
}
