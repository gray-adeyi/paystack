package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// SettlementClient interacts with endpoints related to paystack settlement resource that lets you
// gain insights into payouts made by Paystack to your bank account.
type SettlementClient struct {
	*restClient
}

// NewSettlementClient creates a SettlementClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	sClient := p.NewSettlementClient(p.WithSecretKey("<paystack-secret-key>"))
func NewSettlementClient(options ...ClientOptions) *SettlementClient {
	client := NewPaystackClient(options...)
	return client.Settlements
}

// All lets you retrieve Settlements made to your settlement accounts
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	sClient := p.NewSettlementClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a settlement client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Settlements field is a `SettlementClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Settlements.All()
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := sClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/settlement/#list for supported query parameters
//
//	resp, err := sClient.All()
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
func (s *SettlementClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/settlement", queries...)
	return s.APICall(ctx, http.MethodGet, url, nil, response)
}

// AllTransactions lets you retrieve the Transactions that make up a particular settlement
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	sClient := p.NewSettlementClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a settlement client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Settlements field is a `SettlementClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Settlements.AllTransactions("<settlementId>")
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := sClient.AllTransactions("<settlementId>", p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/settlement/#transactions for supported query parameters
//
//	resp, err := sClient.AllTransactions("<settlementId>")
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
func (s *SettlementClient) AllTransactions(ctx context.Context, settlementId string, response any, queries ...Query) error {
	url := AddQueryParamsToUrl(fmt.Sprintf("/settlement/%s", settlementId), queries...)
	return s.APICall(ctx, http.MethodGet, url, nil, response)
}
