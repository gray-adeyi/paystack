package client

// OptionalPayloadParameter is a type for storing optional parameters used by some APIClient methods that needs
// to accept optional parameter.
type OptionalPayloadParameter = func(map[string]any) map[string]any

// WithOptionalParameter lets you add optional parameters when calling some client methods and you need to add
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
//	resp, err := client.DedicatedVirtualAccounts.Create(context.TODO(),"481193", p.WithOptionalParameter("preferred_bank","wema-bank"))
//
// WithOptionalParameter is used to pass the `preferred_bank` optional parameter in the client method call
func WithOptionalParameter(key string, value any) OptionalPayloadParameter {
	return func(m map[string]any) map[string]any {
		m[key] = value
		return m
	}
}