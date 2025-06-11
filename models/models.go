package models

// Response is a struct containing the status code and data retrieved from paystack. Response.Data is a slice of
// byte that is JSON serializable.
type Response[T any] struct {
	// StatusCode is the http status code returned from making an http request to Paystack
	StatusCode int    `json:"status_code"`
	Status     bool   `json:"status"`
	Message    string `json:"message"`
	// Data is a json serializable slice of byte containing the response data returned from Paystack from calling
	// any of the client associated methods
	Data T       `json:"data"`
	Meta *Meta   `json:"meta"`
	Type *string `json:"type"`
	Code *string `json:"code"`
	Raw  []byte
}

type Meta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage   int `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}
