package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// TransactionSplitClient interacts with endpoints related to paystack Transaction Split resource
// that allows you to split the settlement for a transaction across a payout account, and one or
// more subaccounts.
type TransactionSplitClient struct {
	*restClient
}

// NewTransactionSplitClient creates a TransactionSplitClient
//
// Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	txnSplitClient := p.NewTransactionSplitClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTransactionSplitClient(options ...ClientOptions) *TransactionSplitClient {
	client := NewPaystackClient(options...)
	return client.TransactionSplits
}

// Create lets you create a split payment on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnSplitClient := p.NewTransactionSplitClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction split client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// subaccounts := []map[string]interface{}{
//	// {"subaccount": "ACCT_z3x6z3nbo14xsil", "share": 20},
//	// {"subaccount": "ACCT_pwwualwty4nhq9d", "share": 80},
//	// }
//	// resp, err := paystackClient.TransactionSplits.Create("co-founders account","percentage","NGN",
//	//	subaccounts,"subaccount","ACCT_hdl8abxl8drhrl3")
//
// subaccounts := []map[string]interface{}{
// {"subaccount": "ACCT_z3x6z3nbo14xsil", "share": 20},
// {"subaccount": "ACCT_pwwualwty4nhq9d", "share": 80},
// }
// resp, err := transactionSplitClient.Create("co-founders account","percentage","NGN",
//
//	subaccounts,"subaccount","ACCT_hdl8abxl8drhrl3")
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
func (t *TransactionSplitClient) Create(ctx context.Context, name string, transactionSplitType string, currency string, subaccounts interface{}, bearerType string, bearerSubaccount string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":              name,
		"type":              transactionSplitType,
		"currency":          currency,
		"subaccounts":       subaccounts,
		"bearer_type":       bearerType,
		"bearer_subaccount": bearerSubaccount,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/split", payload, response)
}

// All let you list the transaction splits available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnSplitClient := p.NewTransactionSplitClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction split client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransactionSplits.All()
//
//	// All also accepts queries, so say you want to filter by the name of the split
//	// and if it is active, you can write it like so.
//	// resp, err := txnSplitClient.All(p.WithQuery("name","co-founders account"), p.WithQuery("active", true))
//
//	// see https://paystack.com/docs/api/split/#list for supported query parameters
//
//	resp, err := txnSplitClient.All()
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
func (t *TransactionSplitClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/split", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you get the details of a split on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnSplitClient := p.NewTransactionSplitClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction split client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransactionSplits.FetchOne("<id>")
//
//	resp, err := txnSplitClient.FetchOne("<id>")
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
func (t *TransactionSplitClient) FetchOne(ctx context.Context, id string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/split/%s", id), nil, response)
}

// Update lets you update a transaction split details on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnSplitClient := p.NewTransactionSplitClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction split client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransactionSplits.Update("143", "co-authors account", true)
//
//	// you can pass in optional parameters to the `txnSplitClient.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `bearer_type`.
//	// resp, err := txnSplitClient.Update("143", "co-authors account", true, p.WithOptionalParameter("bearer_type","all"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/split/#update
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//	resp, err := txnSplitClient.Update("143", "co-authors account", true)
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
func (t *TransactionSplitClient) Update(ctx context.Context, id string, name string, active bool, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":   name,
		"active": active,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return t.APICall(ctx, http.MethodPut, fmt.Sprintf("/split/%s", id), payload, response)
}

// Add lets you add a Subaccount to a Transaction Split, or update the share of an existing
// Subaccount in a Transaction Split
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnSplitClient := p.NewTransactionSplitClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction split client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransactionSplits.Add("ACCT_hdl8abxl8drhrl3", 15)
//
//	resp, err := txnSplitClient.Add("ACCT_hdl8abxl8drhrl3", 15)
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
func (t *TransactionSplitClient) Add(ctx context.Context, id string, subAccount string, share int, response any) error {
	payload := map[string]interface{}{
		"subaccount": subAccount,
		"share":      share,
	}
	return t.APICall(ctx, http.MethodPost, fmt.Sprintf("/split/%s/add", id), payload, response)
}

// Remove lets you remove a subaccount from a transaction split
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnSplitClient := p.NewTransactionSplitClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction split client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransactionSplits.Remove("143","ACCT_hdl8abxl8drhrl3")
//
//	resp, err := txnSplitClient.Remove("143","ACCT_hdl8abxl8drhrl3")
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
func (t *TransactionSplitClient) Remove(ctx context.Context, id string, subAccount string, response any) error {
	payload := map[string]any{
		"subaccount": subAccount,
	}

	return t.APICall(ctx, http.MethodPost, fmt.Sprintf("/split/%s/remove", id), payload, response)
}
