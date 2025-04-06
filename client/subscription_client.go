package client

import (
	"context"
	"fmt"
	"net/http"
)

// SubscriptionClient interacts with endpoints related to paystack subscription resource that lets you
// create and manage recurring payment on your Integration.
type SubscriptionClient struct {
	*restClient
}

// NewSubscriptionClient create a SubscriptionClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	subClient := p.NewSubscriptionClient(p.WithSecretKey("<paystack-secret-key>")
func NewSubscriptionClient(options ...ClientOptions) *SubscriptionClient {
	client := NewPaystackClient(options...)

	return client.Subscriptions
}

// Create lets you create a subscription on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	subClient := p.NewSubAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subscription client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Subscriptions.Create("CUS_xnxdt6s1zg1f4nx", "PLN_gx2wn530m0i3w3m", "AUTH_xxx")
//
//	// you can pass in optional parameters to the `SubAccounts.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `start_date`.
//	// resp, err := subClient.Create("CUS_xnxdt6s1zg1f4nx", "PLN_gx2wn530m0i3w3m", "AUTH_xxx",
//	//	p.WithOptionalParameter("start_date","2023-10-16T00:30:13+01:00"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/subscription/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := subClient.CreateCreate("CUS_xnxdt6s1zg1f4nx", "PLN_gx2wn530m0i3w3m", "AUTH_xxx")
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
func (s *SubscriptionClient) Create(ctx context.Context, customer string, plan string, authorization string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"customer":      customer,
		"plan":          plan,
		"authorization": authorization,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(ctx, http.MethodPost, "/subscription", payload)
}

// All lets you retrieve Subscriptions available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	subClient := p.NewSubscriptionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subscription client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Subscriptions.All()
//
//	// All also accepts queries, so say you want to customize how many Subscriptions to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := subClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/subscription/#list for supported query parameters
//
//	resp, err := subClient.All()
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
func (s *SubscriptionClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/subscription", queries...)
	return s.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you retrieve details of a subscription on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	subClient := p.NewSubscriptionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subscription client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.FetchOne("<idOrCode>")
//
//	resp, err := subClient.FetchOne("<idOrCode>")
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
func (s *SubscriptionClient) FetchOne(ctx context.Context, idOrCode string) (*Response, error) {
	return s.APICall(ctx, http.MethodGet, fmt.Sprintf("/subscription/%s", idOrCode), nil)
}

// Enable lets you enable a subscription on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	subClient := p.NewSubscriptionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subscription client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.Enable("SUB_vsyqdmlzble3uii", "d7gofp6yppn3qz7")
//
//	resp, err := subClient.Enable("SUB_vsyqdmlzble3uii", "d7gofp6yppn3qz7")
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
func (s *SubscriptionClient) Enable(ctx context.Context, code string, token string) (*Response, error) {
	payload := map[string]any{
		"code":  code,
		"token": token,
	}

	return s.APICall(ctx, http.MethodPost, "/subscription/enable", payload)
}

// Disable lets you disable a subscription on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	subClient := p.NewSubscriptionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subscription client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.Disable("SUB_vsyqdmlzble3uii", "d7gofp6yppn3qz7")
//
//	resp, err := subClient.Disable("SUB_vsyqdmlzble3uii", "d7gofp6yppn3qz7")
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
func (s *SubscriptionClient) Disable(ctx context.Context, code string, token string) (*Response, error) {
	payload := map[string]any{
		"code":  code,
		"token": token,
	}
	return s.APICall(ctx, http.MethodPost, "/subscription/disable", payload)
}

// GenerateLink lets you generate a link for updating the card on a subscription
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	subClient := p.NewSubscriptionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subscription client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.GenerateLink("SUB_vsyqdmlzble3uii")
//
//	resp, err := subClient.GenerateLink("SUB_vsyqdmlzble3uii")
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
func (s *SubscriptionClient) GenerateLink(ctx context.Context, code string) (*Response, error) {
	return s.APICall(ctx, http.MethodGet, fmt.Sprintf("/subscription/%s/manage/link/", code), nil)
}

// SendLink lets you email a customer a link for updating the card on their subscription
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	subClient := p.NewSubscriptionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a subscription client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.SubAccounts.SendLink("SUB_vsyqdmlzble3uii")
//
//	resp, err := subClient.SendLink("SUB_vsyqdmlzble3uii")
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
func (s *SubscriptionClient) SendLink(ctx context.Context, code string) (*Response, error) {
	return s.APICall(ctx, http.MethodPost, fmt.Sprintf("/subscription/%s/manage/email/", code), nil)
}
