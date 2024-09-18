package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// DedicatedVirtualAccountClient interacts with endpoints related to paystack dedicated virtual account
// resource that enables Nigerian merchants to manage unique payment accounts of their Customers.
type DedicatedVirtualAccountClient struct {
	*baseAPIClient
}

// NewDedicatedVirtualAccountClient creates a DedicatedVirtualAccountClient
//
// Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
func NewDedicatedVirtualAccountClient(options ...ClientOptions) *DedicatedVirtualAccountClient {
	client := NewAPIClient(options...)

	return client.DedicatedVirtualAccounts
}

// Create lets you create a dedicated virtual account for an existing customer
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dedicated virtual account client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.Create("481193")
//
//	// you can pass in optional parameters to the `DedicatedVirtualAccounts.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `preferred_bank`.
//	// resp, err := dvaClient.Create("481193", p.WithOptionalParameter("preferred_bank","wema-bank"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/dedicated-virtual-account/#create
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//
//	resp, err := dvaClient.Create("481193")
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
func (d *DedicatedVirtualAccountClient) Create(ctx context.Context, customerIdOrCode string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"customer": customerIdOrCode,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(ctx, http.MethodPost, "/dedicated_account", payload)
}

// Assign lets you can create a customer, validate the customer, and assign a DVA to the customer.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.Assign("janedoe@test.com","Jane",
//	//	"Doe","Karen", "+2348100000000", "test-bank", "NG")
//
//	// you can pass in optional parameters to the `DedicatedVirtualAccounts.Assign` with `p.WithOptionalParameter`
//	// for example say you want to specify the `account_number`.
//	// resp, err := dvaClient.Assign("janedoe@test.com","Jane", "Doe","Karen", "+2348100000000", "test-bank", "NG",
//	//	p.WithOptionalParameter("account_number","5273681014"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/dedicated-virtual-account/#create
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//
//	resp, err := dvaClient.Assign("janedoe@test.com","Jane", "Doe","Karen", "+2348100000000", "test-bank", "NG")
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
func (d *DedicatedVirtualAccountClient) Assign(ctx context.Context, email string, firstName string, lastName string,
	phone string, preferredBank string, country string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"email":          email,
		"first_name":     firstName,
		"last_name":      lastName,
		"phone":          phone,
		"preferred_bank": preferredBank,
		"country":        country,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(ctx, http.MethodPost, "/dedicated_account/assign", payload)
}

// All lets you retrieve dedicated virtual accounts available on your Integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dedicated virtual account client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.All()
//
//	// All also accepts queries, so say you want to customize how many Transactions to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := dvaClient.All(p.WithQuery("active","true"), p.WithQuery("bank_id","035"))
//
//	// see https://paystack.com/docs/api/dedicated-virtual-account/#list for supported query parameters
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
func (d *DedicatedVirtualAccountClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/dedicated_account", queries...)
	return d.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you retrieve details of a dedicated virtual account on your Integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dedicated virtual account client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.FetchOne("<dedicatedAccountId>")
//
//	resp, err := dvaClient.FetchOne("<dedicatedAccountId>")
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
func (d *DedicatedVirtualAccountClient) FetchOne(ctx context.Context, dedicatedAccountId string) (*Response, error) {
	return d.APICall(ctx, http.MethodGet, fmt.Sprintf("/dedicated_account/%s", dedicatedAccountId), nil)
}

// Requery lets you requery Dedicated Virtual Account for new Transactions
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dedicated virtual account client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.Requery()
//
//	// All also accepts queries, so say you want to customize how many Transactions to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := dvaClient.Requery(p.WithQuery("account_number","1234567890"), p.WithQuery("provider_slug","example-provider"))
//
// // see https://paystack.com/docs/api/dedicated-virtual-account/#requery for supported query parameters
//
//	resp, err := txnClient.Requery()
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
func (d *DedicatedVirtualAccountClient) Requery(ctx context.Context, queries ...Query) (*Response, error) {
	return d.All(ctx, queries...)
}

// Deactivate lets you deactivate a dedicated virtual account on your Integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dedicated virtual account client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.Deactivate("<dedicatedAccountId>")
//
//	resp, err := dvaClient.Deactivate("<dedicatedAccountId>")
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
func (d *DedicatedVirtualAccountClient) Deactivate(ctx context.Context, id string) (*Response, error) {
	return d.APICall(ctx, http.MethodDelete, fmt.Sprintf("/dedicated_account/%s", id), nil)
}

// Split lets you split a dedicated virtual account transaction with one or more accounts
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.Split("<customerIdOrCode>")
//
//	// you can pass in optional parameters to the `DedicatedVirtualAccounts.Split` with `p.WithOptionalParameter`
//	// for example say you want to specify the `preferred_bank`.
//	// resp, err := dvaClient.Split("<customerIdOrCode>", p.WithOptionalParameter("preferred_bank","wema-bank"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/dedicated-virtual-account/#add-split
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//
//	resp, err := dvaClient.Split("<customerIdOrCode>")
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
func (d *DedicatedVirtualAccountClient) Split(ctx context.Context, customerIdOrCode string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"customer": customerIdOrCode,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return d.APICall(ctx, http.MethodPost, "/dedicated_account/split", payload)
}

// RemoveSplit lets you remove a split payment for Transactions. If you've previously set up split payment
// for Transactions on a dedicated virtual account
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dedicated virtual account client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.RemoveSplit("<accountNumber>")
//
//	resp, err := dvaClient.RemoveSplit("<accountNumber>")
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
func (d *DedicatedVirtualAccountClient) RemoveSplit(ctx context.Context, accountNumber string) (*Response, error) {
	payload := map[string]any{
		"account_number": accountNumber,
	}
	return d.APICall(ctx, http.MethodDelete, "/dedicated_account/split", payload)
}

// BankProviders lets you retrieve available bank providers for a dedicated virtual account
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dvaClient := p.NewDedicatedVirtualAccountClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dedicated virtual account client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.DedicatedVirtualAccounts field is a `DedicatedVirtualAccountClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.DedicatedVirtualAccounts.BankProviders()
//
//	resp, err := dvaClient.BankProviders()
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
func (d *DedicatedVirtualAccountClient) BankProviders(ctx context.Context) (*Response, error) {
	return d.APICall(ctx, http.MethodPost, "/dedicated_account/available_providers", nil)
}
