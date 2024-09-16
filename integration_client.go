package paystack

import "net/http"

// IntegrationClient interacts with endpoints related to paystack Integration resource
// that lets you manage some settings on your Integration.
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
	return client.Integration
}

// Timeout lets you retrieve the payment session timeout on your Integration
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
//	// Alternatively, you can access an Integration client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Integration field is a `IntegrationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Integration.Timeout()
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
	return i.APICall(http.MethodGet, "/Integration/payment_session_timeout", nil)
}

// UpdateTimeout lets you update the payment session timeout on your Integration
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
//	// Alternatively, you can access an Integration client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Integration field is a `IntegrationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Integration.UpdateTimeout(5)
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
	return i.APICall(http.MethodPut, "/Integration/payment_session_timeout", payload)
}
