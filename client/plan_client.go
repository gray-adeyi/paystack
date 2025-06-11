package client

import (
	"context"
	"fmt"
	"net/http"
)

// PlanClient interacts with endpoints related to paystack plan resource that lets you
// create and manage installment payment options on your Integration.
type PlanClient struct {
	*restClient
}

// NewPlanClient creates a PlanClient
//
//	Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	planClient := p.NewPlanClient(p.WithSecretKey("<paystack-secret-key>"))
func NewPlanClient(options ...ClientOptions) *PlanClient {
	client := NewPaystackClient(options...)

	return client.Plans
}

// Create lets you create a plan on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	planClient := p.NewPlanClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a plan client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Plans field is a `PlanClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Plans.Create("Monthly retainer", 500000, "monthly")
//
//	// you can pass in optional parameters to the `Plans.Create` with `p.WithOptionalParameter`
//	// for example, you want to specify the `description`.
//	// resp, err := planClient.Create("Monthly retainer", 500000, "monthly",
//	//	p.WithOptionalParameter("description","a test description"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/plan/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := planClient.Create("Monthly retainer", 500000, "monthly")
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
func (p *PlanClient) Create(ctx context.Context, name string, amount int, interval string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":     name,
		"amount":   amount,
		"interval": interval,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPost, "/plan", payload, response)
}

// All lets you retrieve Plans available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	planClient := p.NewPlanClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a plan client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Plans field is a `PlanClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Plans.All()
//
//	// All also accepts queries, so say you want to customize how many Plans to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := planClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/plan/#list for supported query parameters
//
//	resp, err := planClient.All()
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
func (p *PlanClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/plan", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve details of a plan on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	planClient := p.NewPlanClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a plan client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Plans field is a `PlanClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Plans.FetchOne("<idOrCode>")
//
//	resp, err := planClient.FetchOne("<idOrCode>")
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
func (p *PlanClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/plan/%s", idOrCode), nil, response)
}

// Update lets you update a plan details on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	planClient := p.NewPlanClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a plan client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Plans field is a `PlanClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Plans.Update("<idOrCode>","Monthly retainer", 500000, "monthly")
//
//	// you can pass in optional parameters to the `SubAccounts.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `description`.
//	// resp, err := saClient.Update("<idOrCode>","Monthly retainer", 500000, "monthly",
//	//	p.WithOptionalParameter("description","test description"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/subaccount/#update
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := planClient.Update("<idOrCode>","Monthly retainer", 500000, "monthly")
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
func (p *PlanClient) Update(ctx context.Context, idOrCode string, name string, amount int, interval string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":     name,
		"amount":   amount,
		"interval": interval,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPut, fmt.Sprintf("/plan/%s", idOrCode), payload, response)
}
