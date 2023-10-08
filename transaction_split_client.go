package paystack

import (
	"fmt"
	"net/http"
)

// TransactionSplitClient interacts with endpoints related to paystack Transaction Split resource
// that allows you to split the settlement for a transaction across a payout account, and one or
// more subaccounts.
type TransactionSplitClient struct {
	*baseAPIClient
}

// NewTransactionSplitClient creates a TransactionSplitClient
//
// Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	txnSplitClient := p.NewTransactionSplitClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTransactionSplitClient(options ...ClientOptions) *TransactionSplitClient {
	client := NewAPIClient(options...)
	return client.transactionSplits
}

// Create lets you create a split payment on your integration
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
//	// paystackClient.transactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// subaccounts := []map[string]interface{}{
//	// {"subaccount": "ACCT_z3x6z3nbo14xsil", "share": 20},
//	// {"subaccount": "ACCT_pwwualwty4nhq9d", "share": 80},
//	// }
//	// resp, err := paystackClient.transactionSplits.Create("co-founders account","percentage","NGN",
//	//	subaccounts,"subaccount","ACCT_hdl8abxl8drhrl3")
//
// subaccounts := []map[string]interface{}{
// {"subaccount": "ACCT_z3x6z3nbo14xsil", "share": 20},
// {"subaccount": "ACCT_pwwualwty4nhq9d", "share": 80},
// }
// resp, err := paystackClient.transactionSplits.Create("co-founders account","percentage","NGN",
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
func (t *TransactionSplitClient) Create(name string, transactionSplitType string, currency string, subaccounts interface{}, bearerType string, bearerSubaccount string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
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
	return t.APICall(http.MethodPost, "/split", payload)
}

// All let you list the transaction splits available on your integration
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
//	// paystackClient.transactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transactionSplits.All()
//
//	// All also accepts queries, so say you want to filter by the name of the split
//	// and if it is active, you can write it like so.
//	// resp, err := txnSplitClient.All(p.WithQuery("name","co-founders account"), p.WithQuery("active", true))
//
// // see https://paystack.com/docs/api/split/#list for supported query parameters
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
func (t *TransactionSplitClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/split", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

// FetchOne lets you get the details of a split on your integration
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
//	// paystackClient.transactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transactionSplits.FetchOne("<id>")
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
func (t *TransactionSplitClient) FetchOne(id string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/split/%s", id), nil)
}

// Update lets you update a transaction split details on your integration
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
//	// paystackClient.transactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transactionSplits.Update("143", "co-authors account", true)
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
func (t *TransactionSplitClient) Update(id string, name string, active bool, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"name":   name,
		"active": active,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return t.APICall(http.MethodPut, fmt.Sprintf("/split/%s", id), payload)
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
//	// paystackClient.transactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transactionSplits.Add("ACCT_hdl8abxl8drhrl3", 15)
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
func (t *TransactionSplitClient) Add(id string, subAccount string, share int) (*Response, error) {
	payload := map[string]interface{}{
		"subaccount": subAccount,
		"share":      share,
	}
	return t.APICall(http.MethodPost, fmt.Sprintf("/split/%s/add", id), payload)
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
//	// paystackClient.transactionSplits field is a `TransactionSplitClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transactionSplits.Remove("143","ACCT_hdl8abxl8drhrl3")
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
func (t *TransactionSplitClient) Remove(id string, subAccount string) (*Response, error) {
	payload := map[string]interface{}{
		"subaccount": subAccount,
	}

	return t.APICall(http.MethodPost, fmt.Sprintf("/split/%s/remove", id), payload)
}
