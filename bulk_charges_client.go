package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// BulkChargeClient interacts with endpoints related to paystack bulk Charges resource that lets
// you create and manage multiple recurring payments from your Customers.
type BulkChargeClient struct {
	*baseAPIClient
}

// NewBulkChargeClient creates a BulkChargeClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
// bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
func NewBulkChargeClient(options ...ClientOptions) *BulkChargeClient {
	client := NewAPIClient(options...)
	return client.BulkCharges
}

// Initiate lets you send an array of map with authorization codes and amount, using the
// supported currency format, so paystack can process Transactions as a batch.
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
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.BulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	//	batch := []map[string]any{
//	//	{"authorization": "AUTH_ncx8hews93", "amount": 2500, "reference": "dam1266638dhhd"},
//	//	{"authorization": "AUTH_xfuz7dy4b9", "amount": 1500, "reference": "dam1266638dhhe"},
//	//	}
//	//	resp, err := paystackClient.BulkCharges.Initiate(context.TODO(),batch)
//
//	batch := []map[string]any{
//	{"authorization": "AUTH_ncx8hews93", "amount": 2500, "reference": "dam1266638dhhd"},
//	{"authorization": "AUTH_xfuz7dy4b9", "amount": 1500, "reference": "dam1266638dhhe"},
//	}
//	resp, err := bcClient.Initiate(context.TODO(),batch)
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
func (b *BulkChargeClient) Initiate(ctx context.Context, charges interface{}) (*Response, error) {
	return b.APICall(ctx, http.MethodPost, "/bulkcharge", charges)
}

// All lets you retrieve all bulk charge batches created by the Integration.
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
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.BulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.BulkCharges.All(context.TODO())
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := bcClient.All(context.TODO(),p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/bulk-charge/#list for supported query parameters
//
//	resp, err := bcClient.All(context.TODO())
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
func (b *BulkChargeClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/bulkcharge", queries...)
	return b.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you retrieve a specific batch code. It also returns useful information
// on its progress by way of the `total_charges` and `pending_charges` attributes.
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
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.BulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.BulkCharges.FetchOne(context.TODO(),"<idOrCode>")
//
//	resp, err := bcClient.FetchOne(context.TODO(),"<idOrSlug>")
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
func (b *BulkChargeClient) FetchOne(ctx context.Context, idOrCode string) (*Response, error) {
	return b.APICall(ctx, http.MethodGet, fmt.Sprintf("/bulkcharge/%s", idOrCode), nil)
}

// Charges lets you retrieve the Charges associated with a specified batch code.
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
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.BulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.BulkCharges.Charges(context.TODO(),"<idOrCode>")
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := bcClient.Charges(context.TODO(),"<idOrSlug>",p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/bulk-charge/#fetch-charge for supported query parameters
//
//	resp, err := bcClient.Charges(context.TODO(),"<idOrSlug>")
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
func (b *BulkChargeClient) Charges(ctx context.Context, idOrCode string, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl(fmt.Sprintf("/bulkcharge/%s/Charges", idOrCode), queries...)
	return b.APICall(ctx, http.MethodGet, url, nil)
}

// Pause lets you pause a processing a batch
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
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.BulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.BulkCharges.Pause(context.TODO(),"<idOrCode>")
//
//	resp, err := bcClient.Pause(context.TOOD(),"<idOrSlug>")
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
func (b *BulkChargeClient) Pause(ctx context.Context, idOrCode string) (*Response, error) {
	return b.APICall(ctx, http.MethodGet, fmt.Sprintf("/bulkcharge/pause/%s", idOrCode), nil)
}

// Resume lets you resume a paused batch
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
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.BulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.BulkCharges.Resume(context.TODO(),"<idOrCode>")
//
//	resp, err := bcClient.Resume(context.TODO(),"<idOrSlug>")
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
func (b *BulkChargeClient) Resume(ctx context.Context, idOrCode string) (*Response, error) {
	return b.APICall(ctx, http.MethodGet, fmt.Sprintf("/bulkcharge/resume/%s", idOrCode), nil)
}
