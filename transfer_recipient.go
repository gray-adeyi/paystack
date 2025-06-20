package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// TransferRecipientClient interacts with endpoints related to paystack transfer recipient resource
// that lets you create and manage beneficiaries that you send money to.
type TransferRecipientClient struct {
	*restClient
}

// NewTransferRecipientClient creates a TransferRecipientClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	trClient := p.NewTransferRecipientClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTransferRecipientClient(options ...ClientOptions) *TransferRecipientClient {
	client := NewClient(options...)
	return client.TransferRecipients
}

// Create lets you create a new recipient. A duplicate account number will lead to the retrieval of the existing record.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	trClient := p.NewTransferRecipientClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the transfer recipient client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferRecipients field is a `TransferRecipientClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferRecipients.Create("nuban","Tolu Robert","01000000010", "058")
//
//	// you can pass in optional parameters to the `PaymentPages.Create` with `p.WithOptionalParameter`
//	// for example, say you want to specify the `currency`.
//	// resp, err := ppClient.Create("nuban","Tolu Robert","01000000010", "058", p.WithOptionalParameter("currency","NGN"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/transfer-recipient/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := trClient.Create("nuban","Tolu Robert","01000000010", "058")
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
func (t *TransferRecipientClient) Create(ctx context.Context, recipientType string, name string, accountNumber string,
	bankCode string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"type":           recipientType,
		"name":           name,
		"account_number": accountNumber,
		"bank_code":      bankCode,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transferrecipient", payload, response)
}

// BulkCreate lets you create multiple transfer recipients in batches. A duplicate account number will lead to the retrieval of the existing record.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	trClient := p.NewTransferRecipientClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the transfer recipient client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferRecipients field is a `TransferRecipientClient`
//	// Therefore, this is possible
//	//	batch := []map[string]interface{}{
//	//	{"type":"nuban", "name" : "Habenero Mundane", "account_number": "0123456789","bank_code":"033","currency": "NGN"},
//	//	{"type":"nuban","name" : "Soft Merry","account_number": "98765432310","bank_code": "50211","currency": "NGN"},
//	//	}
//	//	resp, err := paystackClient.TransferRecipients.BulkCreate(batch)
//
//	batch := []map[string]interface{}{
//		{"type":"nuban", "name" : "Habenero Mundane", "account_number": "0123456789","bank_code":"033","currency": "NGN"},
//		{"type":"nuban","name" : "Soft Merry","account_number": "98765432310","bank_code": "50211","currency": "NGN"},
//		}
//
//	resp, err := trClient.BulkCreate(batch)
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
func (t *TransferRecipientClient) BulkCreate(ctx context.Context, batch, response any) error {
	payload := map[string]any{
		"batch": batch,
	}
	payload["batch"] = batch

	return t.APICall(ctx, http.MethodPost, "/transferrecipient/bulk", payload, response)
}

// All lets you retrieve transfer recipients available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	trClient := p.NewTransferRecipientClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer recipient client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferRecipients field is a `TransferRecipientClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferRecipients.All()
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := trClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/transfer-recipient/#list for supported query parameters
//
//	resp, err := trClient.All()
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
func (t *TransferRecipientClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/transferrecipient", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve the details of a transfer recipient
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	trClient := p.NewTransferRecipientClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer recipient client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferRecipients field is a `TransferRecipientClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferRecipients.FetchOne("<idOrSlug>")
//
//	resp, err := trClient.FetchOne("<idOrCode>")
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
func (t *TransferRecipientClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transferrecipient/%s", idOrCode), nil, response)
}

// Update lets you update transfer recipients available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	trClient := p.NewTransferRecipientClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer recipient client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferRecipients field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferRecipients.Update("<idOrCode>", "Rick Sanchez")
//
//	// you can pass in optional parameters to the `TransferRecipients.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `email`.
//	// resp, err := trClient.Create("<idOrCode>", "Rick Sanchez", p.WithOptionalParameter("email","johndoe@example.com"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/transfer-recipient/#update
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := trClient.Update("<idOrCode>", "Rick Sanchez")
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
func (t *TransferRecipientClient) Update(ctx context.Context, idOrCode string, name string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name": name,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPut, fmt.Sprintf("/transferrecipient/%s", idOrCode), payload, response)
}

// Delete lets you delete a transfer recipient (sets the transfer recipient to inactive)
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	trClient := p.NewTransferRecipientClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer recipient client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferRecipients field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.PaymentPages.Delete("<idOrCode>")
//
//	resp, err := trClient.Delete("<idOrCode>")
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
func (t *TransferRecipientClient) Delete(ctx context.Context, idOrCode string, response any) error {
	return t.APICall(ctx, http.MethodDelete, fmt.Sprintf("/transferrecipient/%s", idOrCode), nil, response)
}
