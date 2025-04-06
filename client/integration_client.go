package client

import (
	"context"
	"net/http"
)

// IntegrationClient interacts with endpoints related to paystack Integration resource
// that lets you manage some settings on your Integration.
type IntegrationClient struct {
	*restClient
}

// NewIntegrationClient creates an IntegrationClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	intClient := p.NewIntegrationClient(p.WithSecretKey("<paystack-secret-key>"))
func NewIntegrationClient(options ...ClientOptions) *IntegrationClient {
	client := NewPaystackClient(options...)
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
//		"context"
//	)
//
//	intClient := p.NewIntegrationClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an Integration client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Integration field is a `IntegrationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Integration.Timeout(context.TODO())
//
//	resp, err := intClient.Timeout(context.TODO())
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (i *IntegrationClient) Timeout(ctx context.Context) (*Response, error) {
	return i.APICall(ctx, http.MethodGet, "/integration/payment_session_timeout", nil)
}

// UpdateTimeout lets you update the payment session timeout on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//		"context"
//	)
//
//	tcClient := p.NewIntegrationClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an Integration client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Integration field is a `IntegrationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Integration.UpdateTimeout(context.TODO(),5)
//
//	resp, err := tcClient.UpdateTimeout(context.TODO(),5)
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (i *IntegrationClient) UpdateTimeout(ctx context.Context, timeout int) (*Response, error) {
	payload := map[string]any{
		"timeout": timeout,
	}
	return i.APICall(ctx, http.MethodPut, "/integration/payment_session_timeout", payload)
}
