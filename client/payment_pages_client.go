package client

import (
	"context"
	"fmt"
	"net/http"
)

// PaymentPageClient interacts with endpoints
// related to paystack payment page resource
// that lets you provide a quick and secure way to collect payment for Products.
type PaymentPageClient struct {
	*restClient
}

// NewPaymentPageClient creates a PaymentPageClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	ppClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
func NewPaymentPageClient(options ...ClientOptions) *PaymentPageClient {
	client := NewPaystackClient(options...)
	return client.PaymentPages
}

// Create lets you create a payment page on your Integration
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
//	ppClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the payment pages client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.PaymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.PaymentPages.Create(context.TODO(),"Buttercup Brunch")
//
//	// you can pass in optional parameters to the `PaymentPages.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `amount`.
//	// resp, err := ppClient.Create(context.TODO(),"Buttercup Brunch", p.WithOptionalParameter("amount",500000))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/page/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := ppClient.Create(context.TODO(),"Buttercup Brunch")
//
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
func (p *PaymentPageClient) Create(ctx context.Context, name string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"name": name,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPost, "/page", payload)
}

// All lets you retrieve payment pages available on your Integration
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
//	ppClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment pages client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.PaymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.PaymentPages.All(context.TODO())
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := ppClient.All(context.TODO(),p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/page/#list for supported query parameters
//
//	resp, err := ppClient.All(context.TODO())
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
func (p *PaymentPageClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/page", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you retrieve details of a payment page on your Integration
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
//	ppClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment page client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.PaymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.PaymentPages.FetchOne(context.TODO(),"<idOrSlug>")
//
//	resp, err := ppClient.FetchOne(context.TODO(),"<idOrSlug>")
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
func (p *PaymentPageClient) FetchOne(ctx context.Context, idOrSlug string) (*Response, error) {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/page/%s", idOrSlug), nil)
}

// Update lets you update a payment page details on your Integration
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
//	ppClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment page client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.PaymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.PaymentPages.Update(context.TODO(),"<idOrSlug>", "Buttercup Brunch", "description")
//
//	// you can pass in optional parameters to the `SubAccounts.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `amount`.
//	// resp, err := trClient.Create(context.TODO(),"<idOrSlug>", "Buttercup Brunch", "description",
//	//	p.WithOptionalParameter("amount",500000))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/page/#update
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := ppClient.Update(context.TODO(),"<idOrSlug>", "Buttercup Brunch", "description")
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
func (p *PaymentPageClient) Update(ctx context.Context, idOrSlug string, name string, description string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"name":        name,
		"description": description,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPut, fmt.Sprintf("/page/%s", idOrSlug), payload)
}

// CheckSlug lets you check the availability of a slug for a payment page
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
//	ppClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment page client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.PaymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.PaymentPages.CheckSlug(context.TODO(),"<slug>")
//
//	resp, err := ppClient.CheckSlug("<slug>")
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
func (p *PaymentPageClient) CheckSlug(ctx context.Context, slug string) (*Response, error) {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/page/check_slug_availability/%s", slug), nil)
}

// AddProducts lets you add Products to a payment page
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	ppClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment page client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.PaymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.PaymentPages.AddProducts("<id>", []string{4"73", "292"})
//
//	resp, err := ppClient.AddProducts("<id>", []string{4"73", "292"})
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
func (p *PaymentPageClient) AddProducts(ctx context.Context, id string, products []string) (*Response, error) {
	payload := map[string][]string{
		"product": products,
	}
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/page/%s/product", id), payload)
}
