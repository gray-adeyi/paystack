package client

import (
	"context"
	"fmt"
	"net/http"
)

// RefundClient interacts with endpoints related to paystack refund resource that lets you
// create and manage transaction Refunds.
type RefundClient struct {
	*restClient
}

// NewRefundClient creates a RefundClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	refundClient := p.NewRefundClient(p.WithSecretKey("<paystack-secret-key>"))
func NewRefundClient(options ...ClientOptions) *RefundClient {
	client := NewPaystackClient(options...)
	return client.Refunds
}

// Create lets you create and manage transaction Refunds.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	refundClient := p.NewRefundClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the refund client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.refund field is a `RefundClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Refunds.Create("1641")
//
//	// you can pass in optional parameters to the `Refunds.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `amount`.
//	// resp, err := refundClient.Create("1641", p.WithOptionalParameter("amount",500000))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/refund/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := refundClient.Create("1641")
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
func (r *RefundClient) Create(ctx context.Context, transaction string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"transaction": transaction,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return r.APICall(ctx, http.MethodPost, "/refund", payload, response)
}

// All lets you retrieve Refunds available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	refundClient := p.NewRefundClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a refund client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Refunds field is a `RefundClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Refunds.All()
//
//	// All also accepts queries, so say you want to customize how many Refunds to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := refundClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/refund/#list for supported query parameters
//
//	resp, err := refundClient.All()
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
func (r *RefundClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/refund", queries...)
	return r.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve the details of a refund on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	refundClient := p.NewRefundClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a refund client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Refunds field is a `RefundClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Refunds.FetchOne("<reference>")
//
//	resp, err := ppClient.FetchOne("<reference>")
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
func (r *RefundClient) FetchOne(ctx context.Context, reference string, response any) error {
	return r.APICall(ctx, http.MethodGet, fmt.Sprintf("/refund/%s", reference), nil, response)
}
