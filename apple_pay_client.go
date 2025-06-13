package paystack

import (
	"context"
	"net/http"
)

// ApplePayClient interacts with endpoints related to paystack Apple Pay resource that
// lets you register your application's top-level domain or subdomain.
type ApplePayClient struct {
	*restClient
}

// NewApplePayClient creates a ApplePayClient
//
// Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	applePayClient := p.NewApplePayClient(p.WithSecretKey("<paystack-secret-key>"))
func NewApplePayClient(options ...ClientOptions) *ApplePayClient {
	client := NewPaystackClient(options...)

	return client.ApplePay
}

// Register lets you register a top-level domain or subdomain for your Apple Pay Integration.
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
//	applePayClient := p.ApplePayClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an Apple Pay client from APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.ApplePay field is a `ApplePayClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.ApplePay.Register(context.TODO(),"<domainName>")
//
//	resp, err := applePayClient.Register(context.TODO(),"<domainName>")
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
func (a *ApplePayClient) Register(ctx context.Context, domainName string, response any) error {
	payload := map[string]string{"domainName": domainName}
	return a.APICall(ctx, http.MethodPost, "/apple-pay/domain", payload, response)
}

// All lets you retrieve all registered domains on your Integration.
// Returns an empty array if no domains have been added.
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
//	applePayClient := p.NewApplePayClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an Apple Pay client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.ApplePay field is a `ApplePayClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.ApplePay.All(context.TODO())
//
//	// All also accepts queries, so say you want to use cursor, you can write it like so.
//	// resp, err := applePayClient.All(context.TODO(),p.WithQuery("use_cursor","true"))
//
//	// see https://paystack.com/docs/api/apple-pay/#list-domains for supported query parameters
//
//	resp, err := applePayClient.All(context.TODO())
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
func (a *ApplePayClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/apple-pay/domain", queries...)
	return a.APICall(ctx, http.MethodGet, url, nil, response)
}

// Unregister lets you unregister a top-level domain or subdomain previously used for your Apple Pay Integration.
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
//	applePayClient := p.ApplePayClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an Apple Pay client from APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.ApplePay field is a `ApplePayClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.ApplePay.Unregister(context.TODO(),"<domainName>")
//
//	resp, err := applePayClient.Unregister(context.TODO(),"<domainName>")
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
func (a *ApplePayClient) Unregister(ctx context.Context, domainName string, response any) error {
	payload := map[string]string{"domainName": domainName}
	return a.APICall(ctx, http.MethodDelete, "/apple-pay/domain", payload, response)
}
