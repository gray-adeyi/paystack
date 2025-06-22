package paystack

// OptionalPayload is a type for storing optional parameters used by some APIClient methods that needs
// to accept optional parameter.
type OptionalPayload = func(map[string]any) map[string]any

// WithOptionalPayload lets you add optional parameters when calling some client methods and you need to add
// optional parameters to your payload.
//
// Example
//
//	import (
//		p "github.com/gray-adeyi/paystack"
//		"context"
//	)
//
//	client := p.NewAPIClient(p.WithSecretKey("<your-paystack-secret-key>"))
//	resp, err := client.DedicatedVirtualAccounts.Create(context.TODO(),"481193", p.WithOptionalPayload("preferred_bank","wema-bank"))
//
// WithOptionalPayload is used to pass the `preferred_bank` optional parameter in the client method call
func WithOptionalPayload(key string, value any) OptionalPayload {
	return func(m map[string]any) map[string]any {
		m[key] = value
		return m
	}
}
