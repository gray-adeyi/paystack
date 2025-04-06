package client

import (
	"context"
	"fmt"
	"net/http"
)

// TransactionClient interacts with endpoints related to paystack Transaction resource
// that allows you to create and manage payments on your Integration.
type TransactionClient struct {
	*restClient
}

// NewTransactionClient creates a TransactionClient
//
// Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTransactionClient(options ...ClientOptions) *TransactionClient {
	client := NewPaystackClient(options...)

	return client.Transactions

}

// Initialize lets you initialize a transaction from your backend
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.Initialize(200000, "johndoe@example.com")
//
//	// you can pass in optional parameters to the `txnClient.Initialize` with `p.WithOptionalParameter`
//	// for example say you want to specify the currency.
//	// resp, err := txnClient.Initialize(200000, "johndoe@example.com", p.WithOptionalParameter("currency", "NGN"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/transaction/#initialize
//	// Multiple optional parameters can be passed into `Initialize` each with it's `p.WithOptionalParameter`
//	resp, err := txnClient.Initialize(200000, "johndoe@example.com")
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
func (t *TransactionClient) Initialize(ctx context.Context, amount int, email string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"amount": amount,
		"email":  email,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transaction/initialize", payload)
}

// Verify lets you confirm the status of a transaction
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.Verify("<reference>")
//
//	resp, err := txnClient.Verify("<reference>")
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
func (t *TransactionClient) Verify(ctx context.Context, reference string) (*Response, error) {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transaction/verify/%s", reference), nil)
}

// All lets you list Transactions carried out on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.All()
//
//	// All also accepts queries, so say you want to customize how many Transactions to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := txnClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/transaction/#list for supported query parameters
//
//	resp, err := txnClient.All()
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
func (t *TransactionClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you get the details of a transaction carried out on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.FetchOne("<id>")
//
//	resp, err := txnClient.FetchOne("<id>")
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
func (t *TransactionClient) FetchOne(ctx context.Context, id string) (*Response, error) {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transaction/%s", id), nil)
}

// ChargeAuthorization lets you charge authorizations that are marked as reusable
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.ChargeAuthorization(200000,"johndoe@example.com","AUTH_xxx")
//
//	// you can pass in optional parameters to the `txnClient.ChargeAuthorization` with `p.WithOptionalParameter`
//	// for example say you want to specify the currency and channel
//	// resp, err := txnClient.ChargeAuthorization(200000, "johndoe@example.com",
//	p.WithOptionalParameter("currency", "NGN"), p.WithOptionalParameter("channel","bank"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/transaction/#charge-authorization
//	// Multiple optional parameters can be passed into `ChargeAuthorization` each with it's `p.WithOptionalParameter`
//	resp, err := txnClient.ChargeAuthorization(200000,"johndoe@example.com","AUTH_xxx")
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
func (t *TransactionClient) ChargeAuthorization(ctx context.Context, amount int, email string, authorizationCode string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"amount":             amount,
		"email":              email,
		"authorization_code": authorizationCode,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transaction/charge_authorization", payload)
}

// Timeline lets you view the timeline of a transaction
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.Timeline("<idOrReference>")
//
//	resp, err := txnClient.Timeline("<idOrReference>")
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
func (t *TransactionClient) Timeline(ctx context.Context, idOrReference string) (*Response, error) {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transaction/timeline/%s", idOrReference), nil)
}

// Total lets you retrieve the total amount received on your account
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.Total()
//
//	// Total also accepts queries, so say you want to customize how many Transactions to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := txnClient.Total(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/transaction/#totals for supported query parameters
//
//	resp, err := txnClient.Total()
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
func (t *TransactionClient) Total(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction/totals", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil)
}

// Export lets you export a list of Transactions carried out on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.Export()
//
//	// Export also accepts queries, so say you want to customize how many Transactions to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := txnClient.Export(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/transaction/#export for supported query parameters
//
//	resp, err := txnClient.Export()
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
func (t *TransactionClient) Export(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction/export", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil)
}

// PartialDebit lets you retrieve part of a payment from a customer
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	txnClient := p.NewTransactionClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transactions field is a `TransactionClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.PartialDebit("AUTH_xxx","NGN","200000", "johndoe@example.com")
//
//	// you can pass in optional parameters to the `txnClient.PartialDebit` with `p.WithOptionalParameter`
//	// for example say you want to specify the `at_least` optional parameter.
//	// resp, err := txnClient.PartialDebit("AUTH_xxx","NGN","200000", "johndoe@example.com",
//		p.WithOptionalParameter("at_least",100000))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/transaction/#partial-debit
//	// Multiple optional parameters can be passed into `Initialize` each with it's `p.WithOptionalParameter`
//	resp, err := txnClient.PartialDebit("AUTH_xxx","NGN","200000", "johndoe@example.com")
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
func (t *TransactionClient) PartialDebit(ctx context.Context, authorizationCode string, currency string, amount string, email string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"authorization_code": authorizationCode,
		"currency":           currency,
		"amount":             amount,
		"email":              email,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transaction/partial_debit", payload)
}
