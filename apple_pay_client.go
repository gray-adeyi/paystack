package paystack

import "net/http"

// ApplePayClient interacts with endpoints related to paystack Apple Pay resource that
// lets you register your application's top-level domain or subdomain.
type ApplePayClient struct {
	*baseAPIClient
}

// NewApplePayClient creates a ApplePayClient
//
// Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	applePayClient := p.NewApplePayClient(p.WithSecretKey("<paystack-secret-key>"))
func NewApplePayClient(options ...ClientOptions) *ApplePayClient {
	client := NewAPIClient(options...)

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
//	)
//
//	applePayClient := p.ApplePayClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an Apple Pay client from APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.ApplePay field is a `ApplePayClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.ApplePay.Register("<domainName>")
//
//	resp, err := applePayClient.Register("<domainName>")
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
func (a *ApplePayClient) Register(domainName string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["domainName"] = domainName
	return a.APICall(http.MethodPost, "/apple-pay/domain", payload)
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
//	)
//
//	applePayClient := p.NewApplePayClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an Apple Pay client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.ApplePay field is a `ApplePayClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.ApplePay.All()
//
//	// All also accepts queries, so say you want to use cursor, you can write it like so.
//	// resp, err := applePayClient.All(p.WithQuery("use_cursor","true"))
//
// // see https://paystack.com/docs/api/apple-pay/#list-domains for supported query parameters
//
//	resp, err := applePayClient.All()
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
func (a *ApplePayClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/apple-pay/domain", queries...)
	return a.APICall(http.MethodGet, url, nil)
}

// Unregister lets you unregister a top-level domain or subdomain previously used for your Apple Pay Integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	applePayClient := p.ApplePayClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access an Apple Pay client from APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.ApplePay field is a `ApplePayClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.ApplePay.Unregister("<domainName>")
//
//	resp, err := applePayClient.Unregister("<domainName>")
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
func (a *ApplePayClient) Unregister(domainName string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["domainName"] = domainName
	return a.APICall(http.MethodDelete, "/apple-pay/domain", payload)
}
