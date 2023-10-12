package paystack

import (
	"fmt"
	"net/http"
)

// SubscriptionClient interacts with endpoints related to paystack subscription resource that lets you
// create and manage recurring payment on your integration.
type SubscriptionClient struct {
	*baseAPIClient
}

// NewSubscriptionClient create a SubscriptionClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	subClient := p.NewSubscriptionClient(p.WithSecretKey("<paystack-secret-key>")
func NewSubscriptionClient(options ...ClientOptions) *SubscriptionClient {
	client := NewAPIClient(options...)

	return client.subscriptions
}

// Create lets you create a subscription on your integration
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
//	// paystackClient.subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.subscriptions.Create("CUS_xnxdt6s1zg1f4nx", "PLN_gx2wn530m0i3w3m", "AUTH_xxx")
//
//	// you can pass in optional parameters to the `subAccounts.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `start_date`.
//	// resp, err := subClient.Create("CUS_xnxdt6s1zg1f4nx", "PLN_gx2wn530m0i3w3m", "AUTH_xxx",
//	//	p.WithOptionalParameter("start_date","2023-10-16T00:30:13+01:00"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/subscription/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
// resp, err := subClient.CreateCreate("CUS_xnxdt6s1zg1f4nx", "PLN_gx2wn530m0i3w3m", "AUTH_xxx")
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
func (s *SubscriptionClient) Create(customer string, plan string, authorization string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = customer
	payload["plan"] = plan
	payload["authorization"] = authorization

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(http.MethodPost, "/subscription", payload)
}

// All lets you retrieve subscriptions available on your integration
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
//	// paystackClient.subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.subscriptions.All()
//
//	// All also accepts queries, so say you want to customize how many subscriptions to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := subClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
// // see https://paystack.com/docs/api/subscription/#list for supported query parameters
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
func (s *SubscriptionClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/subscription", queries...)
	return s.APICall(http.MethodGet, url, nil)
}

// FetchOne lets you retrieve details of a subscription on your integration
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
//	// paystackClient.subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.subAccounts.FetchOne("<idOrCode>")
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
func (s *SubscriptionClient) FetchOne(idOrCode string) (*Response, error) {
	return s.APICall(http.MethodGet, fmt.Sprintf("/subscription/%s", idOrCode), nil)
}

// Enable lets you enable a subscription on your integration
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
//	// paystackClient.subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.subAccounts.Enable("SUB_vsyqdmlzble3uii", "d7gofp6yppn3qz7")
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
func (s *SubscriptionClient) Enable(code string, token string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["code"] = code
	payload["token"] = token
	return s.APICall(http.MethodPost, "/subscription/enable", payload)
}

// Disable lets you disable a subscription on your integration
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
//	// paystackClient.subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.subAccounts.Disable("SUB_vsyqdmlzble3uii", "d7gofp6yppn3qz7")
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
func (s *SubscriptionClient) Disable(code string, token string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["code"] = code
	payload["token"] = token
	return s.APICall(http.MethodPost, "/subscription/disable", payload)
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
//	// paystackClient.subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.subAccounts.GenerateLink("SUB_vsyqdmlzble3uii")
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
func (s *SubscriptionClient) GenerateLink(code string) (*Response, error) {
	return s.APICall(http.MethodGet, fmt.Sprintf("/subscription/%s/manage/link/", code), nil)
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
//	// paystackClient.subscriptions field is a `SubscriptionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.subAccounts.SendLink("SUB_vsyqdmlzble3uii")
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
func (s *SubscriptionClient) SendLink(code string) (*Response, error) {
	return s.APICall(http.MethodPost, fmt.Sprintf("/subscription/%s/manage/email/", code), nil)
}
