package client

import (
	"context"
	"fmt"
	"net/http"
)

// TransferClient interacts with endpoints related to paystack transfer resource that lets you
// automate sending money to your Customers.
type TransferClient struct {
	*restClient
}

// NewTransferClient creates a TransferClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
// tfClient := p.NewTransferClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTransferClient(options ...ClientOptions) *TransferClient {
	client := NewPaystackClient(options...)
	return client.Transfers
}

// Initiate lets you send money to your Customers.
// Status of a transfer object returned will be pending if OTP is disabled.
// In the event that an OTP is required, status will read otp.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tfClient := p.NewTransferClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the transfer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transfers.Initiate("balance",500000,"RCP_gx2wn530m0i3w3m")
//
//	// you can pass in optional parameters to the `Transfers.Initiate` with `p.WithOptionalParameter`
//	// for example say you want to specify the `reason`.
//	// resp, err := tfClient.Initiate("balance",500000,"RCP_gx2wn530m0i3w3m", p.WithOptionalParameter("reason","Discount Refund"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/transfer/#initiate
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
// resp, err := tfClient.Initiate("balance",500000,"RCP_gx2wn530m0i3w3m")
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
func (t *TransferClient) Initiate(ctx context.Context, source string, amount int, recipient string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"source":    source,
		"amount":    amount,
		"recipient": recipient,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transfer", payload)
}

// Finalize lets you finalize an initiated transfer
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tfClient := p.NewTransferClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the transfer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transfers.Finalize("TRF_vsyqdmlzble3uii","928783")
//
//	resp, err := tfClient.Finalize("TRF_vsyqdmlzble3uii","928783")
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
func (t *TransferClient) Finalize(ctx context.Context, transferCode string, otp string) (*Response, error) {
	payload := map[string]any{
		"transfer_code": transferCode,
		"otp":           otp,
	}

	return t.APICall(ctx, http.MethodPost, "/transfer/finalize_transfer", payload)
}

// BulkInitiate lets you initiate multiple Transfers in a single request.
// You need to disable the Transfers OTP requirement to use this endpoint.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tfClient := p.NewTransferClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the transfer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// Transfers := []map[string]interface{
//	//	{"amount": 20000,"reference": "588YtfftReF355894J","reason": "Why not?","recipient":"RCP_2tn9clt23s7qr28"},
//	//	{"amount": 30000,"reference": "YunoTReF35e0r4J","reason": "Because I can","recipient":"RCP_1a25w1h3n0xctjg"},
//	//	{"amount": 40000,"reason": "Coming right up","recipient": "RCP_aps2aibr69caua7"},
//	}
//	// resp, err := paystackClient.Transfers.BulkInitiate("balance", Transfers)
//
//	Transfers := []map[string]interface{
//		{"amount": 20000,"reference": "588YtfftReF355894J","reason": "Why not?","recipient":"RCP_2tn9clt23s7qr28"},
//		{"amount": 30000,"reference": "YunoTReF35e0r4J","reason": "Because I can","recipient":"RCP_1a25w1h3n0xctjg"},
//		{"amount": 40000,"reason": "Coming right up","recipient": "RCP_aps2aibr69caua7"},
//	}
//
//	resp, err := tfClient.BulkInitiate("balance", Transfers)
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
func (t *TransferClient) BulkInitiate(ctx context.Context, source string, transfers interface{}) (*Response, error) {
	payload := map[string]any{
		"source":    source,
		"Transfers": transfers,
	}

	return t.APICall(ctx, http.MethodPost, "/transfer/bulk", payload)
}

// All lets you retrieve all the Transfers made on your Integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tfClient := p.NewTransferClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transfers.All()
//
//	// All also accepts queries, so say you want to customize how many Transfers to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := tfClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
//	// see https://paystack.com/docs/api/transfer/#list for supported query parameters
//
//	resp, err := tfClient.All()
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
func (t *TransferClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transfer", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you retrieve the details of a transfer on your Integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tfClient := p.NewTransferClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transfers.FetchOne("<idOrCode>")
//
//	resp, err := tfClient.FetchOne("<idOrCode>")
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
func (t *TransferClient) FetchOne(ctx context.Context, idOrCode string) (*Response, error) {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transfer/%s", idOrCode), nil)
}

// Verify lets you verify the status of a transfer on your Integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tfClient := p.NewTransferClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transfers.Verify("<reference>")
//
//	resp, err := tfClient.Verify("<reference>")
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
func (t *TransferClient) Verify(ctx context.Context, reference string) (*Response, error) {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transfer/verify/%s", reference), nil)
}
