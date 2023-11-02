package paystack

import (
	"fmt"
	"net/http"
)

// BulkChargeClient interacts with endpoints related to paystack bulk charges resource that lets
// you create and manage multiple recurring payments from your customers.
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
	return client.bulkCharges
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
//	)
//
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.bulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	//	batch := []map[string]interface{}{
//	//	{"authorization": "AUTH_ncx8hews93", "amount": 2500, "reference": "dam1266638dhhd"},
//	//	{"authorization": "AUTH_xfuz7dy4b9", "amount": 1500, "reference": "dam1266638dhhe"},
//	//	}
//	//	resp, err := paystackClient.bulkCharges.Initiate(batch)
//
//	batch := []map[string]interface{}{
//	{"authorization": "AUTH_ncx8hews93", "amount": 2500, "reference": "dam1266638dhhd"},
//	{"authorization": "AUTH_xfuz7dy4b9", "amount": 1500, "reference": "dam1266638dhhe"},
//	}
//	resp, err := bcClient.Initiate(batch)
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
func (b *BulkChargeClient) Initiate(charges interface{}) (*Response, error) {
	return b.APICall(http.MethodPost, "/bulkcharge", charges)
}

// All lets you retrieve all bulk charge batches created by the integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.bulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.bulkCharges.All()
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := bcClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
// // see https://paystack.com/docs/api/bulk-charge/#list for supported query parameters
//
//	resp, err := bcClient.All()
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
func (b *BulkChargeClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/bulkcharge", queries...)
	return b.APICall(http.MethodGet, url, nil)
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
//	)
//
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.bulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.bulkCharges.FetchOne("<idOrCode>")
//
//	resp, err := bcClient.FetchOne("<idOrSlug>")
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
func (b *BulkChargeClient) FetchOne(idOrCode string) (*Response, error) {
	return b.APICall(http.MethodGet, fmt.Sprintf("/bulkcharge/%s", idOrCode), nil)
}

// Charges lets you retrieve the charges associated with a specified batch code.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.bulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.bulkCharges.Charges("<idOrCode>")
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := bcClient.Charges("<idOrSlug>",p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
// // see https://paystack.com/docs/api/bulk-charge/#fetch-charge for supported query parameters
//
//	resp, err := bcClient.Charges("<idOrSlug>")
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
func (b *BulkChargeClient) Charges(idOrCode string, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl(fmt.Sprintf("/bulkcharge/%s/charges", idOrCode), queries...)
	return b.APICall(http.MethodGet, url, nil)
}

// Pause lets you pause a processing a batch
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.bulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.bulkCharges.Pause("<idOrCode>")
//
//	resp, err := bcClient.Pause("<idOrSlug>")
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
func (b *BulkChargeClient) Pause(idOrCode string) (*Response, error) {
	return b.APICall(http.MethodGet, fmt.Sprintf("/bulkcharge/pause/%s", idOrCode), nil)
}

// Resume lets you resume a paused batch
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	bcClient := p.NewBulkChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a bulk charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.bulkCharges field is a `BulkChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.bulkCharges.Resume("<idOrCode>")
//
//	resp, err := bcClient.Resume("<idOrSlug>")
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
func (b *BulkChargeClient) Resume(idOrCode string) (*Response, error) {
	return b.APICall(http.MethodGet, fmt.Sprintf("/bulkcharge/resume/%s", idOrCode), nil)
}
