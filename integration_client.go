package paystack

import "net/http"

// IntegrationClient interacts with endpoints related to paystack integration resource
// that lets you manage some settings on your integration.
type IntegrationClient struct {
	*baseAPIClient
}

// NewIntegrationClient creates an IntegrationClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	intClient := p.NewIntegrationClient(p.WithSecretKey("<paystack-secret-key>"))
func NewIntegrationClient(options ...ClientOptions) *IntegrationClient {
	client := NewAPIClient(options...)
	return client.integration
}

// Timeout lets you retrieve the payment session timeout on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	intClient := p.NewIntegrationClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an integration client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.integration field is a `IntegrationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.integration.Timeout()
//
//	resp, err := intClient.Timeout()
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (i *IntegrationClient) Timeout() (*Response, error) {
	return i.APICall(http.MethodGet, "/integration/payment_session_timeout", nil)
}

// UpdateTimeout lets you update the payment session timeout on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tcClient := p.NewIntegrationClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an integration client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.integration field is a `IntegrationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.integration.UpdateTimeout(5)
//
//	resp, err := tcClient.UpdateTimeout(5)
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (i *IntegrationClient) UpdateTimeout(timeout int) (*Response, error) {
	payload := map[string]interface{}{
		"timeout": timeout,
	}
	return i.APICall(http.MethodPut, "/integration/payment_session_timeout", payload)
}
