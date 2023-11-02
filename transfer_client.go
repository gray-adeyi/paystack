package paystack

import (
	"fmt"
	"net/http"
)

// TransferClient interacts with endpoints related to paystack transfer resource that lets you
// automate sending money to your Customers.
type TransferClient struct {
	*baseAPIClient
}

// NewTransferClient creates a TransferClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
// tfClient := p.NewTransferClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTransferClient(options ...ClientOptions) *TransferClient {
	client := NewAPIClient(options...)
	return client.transfers
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
//	// paystackClient.transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transfers.Initiate("balance",500000,"RCP_gx2wn530m0i3w3m")
//
//	// you can pass in optional parameters to the `transfers.Initiate` with `p.WithOptionalParameter`
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
func (t *TransferClient) Initiate(source string, amount int, recipient string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["source"] = source
	payload["amount"] = amount
	payload["recipient"] = recipient

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPost, "/transfer", payload)
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
//	// paystackClient.transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transfers.Finalize("TRF_vsyqdmlzble3uii","928783")
//
// resp, err := tfClient.Finalize("TRF_vsyqdmlzble3uii","928783")
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
func (t *TransferClient) Finalize(transferCode string, otp string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["transfer_code"] = transferCode
	payload["otp"] = otp
	return t.APICall(http.MethodPost, "/transfer/finalize_transfer", payload)
}

// BulkInitiate lets you initiate multiple transfers in a single request.
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
//	// paystackClient.transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// transfers := []map[string]interface{
//	//	{"amount": 20000,"reference": "588YtfftReF355894J","reason": "Why not?","recipient":"RCP_2tn9clt23s7qr28"},
//	//	{"amount": 30000,"reference": "YunoTReF35e0r4J","reason": "Because I can","recipient":"RCP_1a25w1h3n0xctjg"},
//	//	{"amount": 40000,"reason": "Coming right up","recipient": "RCP_aps2aibr69caua7"},
//	}
//	// resp, err := paystackClient.transfers.BulkInitiate("balance", transfers)
//
//	transfers := []map[string]interface{
//		{"amount": 20000,"reference": "588YtfftReF355894J","reason": "Why not?","recipient":"RCP_2tn9clt23s7qr28"},
//		{"amount": 30000,"reference": "YunoTReF35e0r4J","reason": "Because I can","recipient":"RCP_1a25w1h3n0xctjg"},
//		{"amount": 40000,"reason": "Coming right up","recipient": "RCP_aps2aibr69caua7"},
//	}
//
//	resp, err := tfClient.BulkInitiate("balance", transfers)
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
func (t *TransferClient) BulkInitiate(source string, transfers interface{}) (*Response, error) {
	payload := make(map[string]interface{})
	payload["source"] = source
	payload["transfers"] = transfers

	return t.APICall(http.MethodPost, "/transfer/bulk", payload)
}

// All lets you retrieve all the transfers made on your integration.
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
//	// paystackClient.transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transfers.All()
//
//	// All also accepts queries, so say you want to customize how many transfers to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := tfClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
// // see https://paystack.com/docs/api/transfer/#list for supported query parameters
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
func (t *TransferClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transfer", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

// FetchOne lets you retrieve the details of a transfer on your integration.
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
//	// paystackClient.transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transfers.FetchOne("<idOrCode>")
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
func (t *TransferClient) FetchOne(idOrCode string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transfer/%s", idOrCode), nil)
}

// Verify lets you verify the status of a transfer on your integration.
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
//	// paystackClient.transfers field is a `TransferClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transfers.Verify("<reference>")
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
func (t *TransferClient) Verify(reference string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transfer/verify/%s", reference), nil)
}
