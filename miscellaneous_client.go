package paystack

import (
	"context"
	"net/http"
)

// MiscellaneousClient interacts with endpoints related to paystack Miscellaneous resource that
// provides information that is relevant to other client methods
type MiscellaneousClient struct {
	*baseAPIClient
}

// NewMiscellaneousClient creates a MiscellaneousClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	miscClient := p.NewMiscellaneousClient(p.WithSecretKey("<paystack-secret-key>"))
func NewMiscellaneousClient(options ...ClientOptions) *MiscellaneousClient {
	client := NewAPIClient(options...)
	return client.Miscellaneous
}

// Banks lets you retrieve a list of all supported banks and their properties
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	miscClient := p.NewMiscellaneousClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment pages client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Miscellaneous field is a `MiscellaneousClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Miscellaneous.Banks()
//
//	// Banks also accepts queries, so say you want to customize how many banks to retrieve
//	// and which country the banks to retrieve are from, you can write it like so.
//	// resp, err := miscClient.Banks(p.WithQuery("perPage","50"), p.WithQuery("country","nigeria"))
//
//	// see https://paystack.com/docs/api/miscellaneous/#bank for supported query parameters
//
//	resp, err := miscClient.Banks()
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
func (p *MiscellaneousClient) Banks(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/bank", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil)
}

// Countries let you retrieve a list of countries that Paystack currently supports
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	miscClient := p.NewMiscellaneousClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment pages client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Miscellaneous field is a `MiscellaneousClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Miscellaneous.Countries()
//
//	resp, err := miscClient.Countries()
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
func (p *MiscellaneousClient) Countries(ctx context.Context) (*Response, error) {
	return p.APICall(ctx, http.MethodGet, "/country", nil)
}

// States lets you retrieve a list of states for a country for address Verification
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	miscClient := p.NewMiscellaneousClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment pages client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Miscellaneous field is a `MiscellaneousClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Miscellaneous.States()
//
//	// States also accepts queries, so say you want to specify country the states
//	// to retrieve are from, you can write it like so.
//	// resp, err := miscClient.States(p.WithQuery("country","nigeria"))
//
//	// see https://paystack.com/docs/api/miscellaneous/#avs-states for supported query parameters
//
//	resp, err := miscClient.States()
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
func (p *MiscellaneousClient) States(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/address_verification/states", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil)
}
