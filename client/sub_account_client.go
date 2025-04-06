package client

import (
	"context"
	"fmt"
	"net/http"
)

// SubAccountClient interacts with endpoints related to paystack subaccount resource that lets you
// create and manage subaccounts on your Integration. Subaccounts can be used to split payment
// between two accounts (your main account and a subaccount).
type SubAccountClient struct {
	*restClient
}

// NewSubAccountClient creates a SubAccountClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	saClient := p.NewSubAccountClient(p.WithSecretKey("<paystack-secret-key>"))
func NewSubAccountClient(options ...ClientOptions) *SubAccountClient {
	client := NewPaystackClient(options...)

	return client.SubAccounts
}

// Create lets you create a dedicated virtual account for an existing customer
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	saClient := p.NewSubAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subaccount client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.SubAccounts field is a `SubAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.Create("Sunshine Studios", "044", "0193274682", 18.2,"")
//
//	// you can pass in optional parameters to the `SubAccounts.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `preferred_bank`.
//	// resp, err := saClient.Create("Sunshine Studios", "044", "0193274682", 18.2,"",
//	//	p.WithOptionalParameter("primary_contact_email","johndoe@example.com"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/subaccount/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := saClient.Create("Sunshine Studios", "044", "0193274682", 18.2,"")
//
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
func (s *SubAccountClient) Create(ctx context.Context, businessName string, settlementBank string,
	accountNumber string, percentageCharge float32, description string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"business_name":     businessName,
		"settlement_bank":   settlementBank,
		"account_number":    accountNumber,
		"percentage_charge": percentageCharge,
		"description":       description,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(ctx, http.MethodPost, "/subaccount", payload)
}

// All lets you retrieve subaccounts available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	saClient := p.NewSubAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subaccount client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.SubAccounts field is a `SubAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.All()
//
//	// All also accepts queries, so say you want to customize how many subaccounts to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := saClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/subaccount/#list for supported query parameters
//
//	resp, err := saClient.All()
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
func (s *SubAccountClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/subaccount", queries...)
	return s.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you retrieve details of a subaccount on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	saClient := p.NewSubAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subaccount client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.SubAccounts field is a `SubAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.FetchOne("<idOrCode>")
//
//	resp, err := saClient.FetchOne("<idOrCode>")
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
func (s *SubAccountClient) FetchOne(ctx context.Context, idOrCode string) (*Response, error) {
	return s.APICall(ctx, http.MethodGet, fmt.Sprintf("/subaccount/%s", idOrCode), nil)
}

// Update lets you update a subaccount details on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	saClient := p.NewSubAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a terminal client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.SubAccounts field is a `SubAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.Update("<idOrCode>", "Sunshine Studios", "044")
//
//	// you can pass in optional parameters to the `SubAccounts.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `preferred_bank`.
//	// resp, err := saClient.Create("<idOrCode>","Sunshine Studios", "044",
//	//	p.WithOptionalParameter("primary_contact_email","johndoe@example.com"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/subaccount/#update
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := saClient.Update("<idOrCode>", "Sunshine Studios", "044")
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
func (s *SubAccountClient) Update(ctx context.Context, idOrCode string, businessName string, settlementBank string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"business_name":   businessName,
		"settlement_bank": settlementBank,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(ctx, http.MethodPut, fmt.Sprintf("/subaccount/%s", idOrCode), payload)
}
