package paystack

import (
	"fmt"
	"net/http"
)

// PaymentPageClient interacts with endpoints
// related to paystack payment page resource
// that lets you provide a quick and secure way to collect payment for Products.
type PaymentPageClient struct {
	*baseAPIClient
}

// NewPaymentPageClient creates a PaymentPageClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	ppClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
func NewPaymentPageClient(options ...ClientOptions) *PaymentPageClient {
	client := NewAPIClient(options...)
	return client.paymentPages
}

// Create lets you create a payment page on your integration
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
//	// Alternatively, you can access the payment pages client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentPages.Create("Buttercup Brunch")
//
//	// you can pass in optional parameters to the `paymentPages.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `amount`.
//	// resp, err := ppClient.Create("Buttercup Brunch", p.WithOptionalParameter("amount",500000))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/page/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
// resp, err := ppClient.Create("Buttercup Brunch")
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
func (p *PaymentPageClient) Create(name string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPost, "/page", payload)
}

// All lets you retrieve payment pages available on your integration
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
//	// Alternatively, you can access a payment pages client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentPages.All()
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := ppClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
// // see https://paystack.com/docs/api/page/#list for supported query parameters
//
//	resp, err := ppClient.All()
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
func (p *PaymentPageClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/page", queries...)
	return p.APICall(http.MethodGet, url, nil)
}

// FetchOne lets you retrieve details of a payment page on your integration
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
//	// paystackClient.paymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentPages.FetchOne("<idOrSlug>")
//
//	resp, err := ppClient.FetchOne("<idOrSlug>")
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
func (p *PaymentPageClient) FetchOne(idOrSlug string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/page/%s", idOrSlug), nil)
}

// Update lets you update a payment page details on your integration
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
//	// paystackClient.paymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentPages.Update("<idOrSlug>", "Buttercup Brunch", "description")
//
//	// you can pass in optional parameters to the `SubAccounts.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `amount`.
//	// resp, err := trClient.Create("<idOrSlug>", "Buttercup Brunch", "description",
//	//	p.WithOptionalParameter("amount",500000))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/page/#update
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := ppClient.Update("<idOrSlug>", "Buttercup Brunch", "description")
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
func (p *PaymentPageClient) Update(idOrSlug string, name string, description string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["description"] = description

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPut, fmt.Sprintf("/page/%s", idOrSlug), payload)
}

// CheckSlug lets you check the availability of a slug for a payment page
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
//	// paystackClient.paymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentPages.CheckSlug("<slug>")
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
func (p *PaymentPageClient) CheckSlug(slug string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/page/check_slug_availability/%s", slug), nil)
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
//	// paystackClient.paymentPages field is a `PaymentPageClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentPages.AddProducts("<id>", []string{4"73", "292"})
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
func (p *PaymentPageClient) AddProducts(id string, products []string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["product"] = products
	return p.APICall(http.MethodGet, fmt.Sprintf("/page/%s/product", id), nil)
}
