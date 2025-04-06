package models

// Response is a struct containing the status code and data retrieved from paystack. Response.Data is a slice of
// byte that is JSON serializable.
type Response struct {
	// StatusCode is the http status code returned from making an http request to Paystack
	StatusCode int
	// Data is a json serializable slice of byte containing the response data returned from Paystack from calling
	// any of the client associated methods
	Data []byte
}