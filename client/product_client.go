package client

import (
	"context"
	"fmt"
	"net/http"
)

// ProductClient interacts with endpoints related to paystack product resource that allows you to create and
// manage inventories on your Integration.
type ProductClient struct {
	*restClient
}

// NewProductClient creates a ProductClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
// prodClient := p.NewProductClient(p.WithSeretKey("<paystack-secret-key>"))
func NewProductClient(options ...ClientOptions) *ProductClient {
	client := NewPaystackClient(options...)
	return client.Products
}

// Create lets you create a product on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prodClient := p.NewProductClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a product client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Products field is a `ProductClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Products.Create("Puff Puff", "Crispy flour ball with fluffy interior", 5000, "NGN")
//
//	// you can pass in optional parameters to the `Products.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `unlimited`.
//	// resp, err := prodClient.Create("Puff Puff", "Crispy flour ball with fluffy interior", 5000, "NGN",
//	//	p.WithOptionalParameter("unlimited","true"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/product/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := prodClient.Create("Puff Puff", "Crispy flour ball with fluffy interior", 5000, "NGN")
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
func (p *ProductClient) Create(ctx context.Context, name string, description string, price int, currency string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"name":        name,
		"description": description,
		"price":       price,
		"currency":    currency,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPost, "/product", payload)
}

// All lets you retrieve Products available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prodClient := p.NewProductClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a product client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Products field is a `ProductClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Products.All()
//
//	// All also accepts queries, so say you want to customize how many Products to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := prodClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/product/#list for supported query parameters
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
func (p *ProductClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/product", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you Get details of a product on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prodClient := p.NewSubAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a product client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Products field is a `ProductClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Products.FetchOne("<id>")
//
//	resp, err := prodClient.FetchOne("<id>")
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
func (p *ProductClient) FetchOne(ctx context.Context, id string) (*Response, error) {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/product/%s", id), nil)
}

// Update lets you update a product details on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prodClient := p.NewProductClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a product client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Products field is a `ProductClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Products.Update("<id>", "Product Six", "Product Six Description",500000, "USD")
//
//	// you can pass in optional parameters to the `Products.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `unlimited`.
//	// resp, err := prodClient.Update("<id>", "Product Six", "Product Six Description",500000, "USD",
//	//	p.WithOptionalParameter("unlimited","true"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/product/#update
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := prodClient.Update("<id>", "Product Six", "Product Six Description",500000, "USD")
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
func (p *ProductClient) Update(ctx context.Context, id string, name string, description string, price int, currency string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"name":        name,
		"description": description,
		"price":       price,
		"currency":    currency,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPut, fmt.Sprintf("/product/%s", id), payload)
}
