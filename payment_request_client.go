package paystack

import (
	"fmt"
	"net/http"
)

// PaymentRequestClient interacts with endpoints related to paystack payment request resource that lets you manage requests for payment of goods and services.
type PaymentRequestClient struct {
	*baseAPIClient
}

// NewPaymentRequestClient creates a PaymentRequestClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	prClient := p.NewPaymentPageClient(p.WithSecretKey("<paystack-secret-key>"))
func NewPaymentRequestClient(options ...ClientOptions) *PaymentRequestClient {
	client := NewAPIClient(options...)
	return client.paymentRequests
}

// Create lets you create a payment request for a transaction on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.Create("CUS_xwaj0txjryg393b", 500000)
//
//	// you can pass in optional parameters to the `paymentRequests.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `due_date`.
//	// resp, err := prClient.Create("CUS_xwaj0txjryg393b", 500000, p.WithOptionalParameter("due_date","2023-12-25"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/payment-request/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
// resp, err := prClient.Create("CUS_xwaj0txjryg393b", 500000)
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
func (p *PaymentRequestClient) Create(customerIdOrCode string, amount int,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = customerIdOrCode
	payload["amount"] = amount

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPost, "/paymentrequest", payload)
}

// All lets you retrieve the payment requests available on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.All()
//
//	// All also accepts queries, so say you want to customize how many payment pages to retrieve
//	// and which page to retrieve, you can write it like so.
//	// resp, err := prClient.All(p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//
// // see https://paystack.com/docs/api/payment-request/#list for supported query parameters
//
//	resp, err := prClient.All()
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
func (p *PaymentRequestClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/paymentrequest", queries...)
	return p.APICall(http.MethodGet, url, nil)
}

// FetchOne lets you retrieve details of a payment request on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.FetchOne("<idOrCode>")
//
//	resp, err := prClient.FetchOne("<idOrCode>")
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
func (p *PaymentRequestClient) FetchOne(idOrCode string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/paymentrequest/%s", idOrCode), nil)
}

// Verify lets you verify the details of a payment request on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.Verify("<code>")
//
//	resp, err := prClient.Verify("<code>")
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
func (p *PaymentRequestClient) Verify(code string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/paymentrequest/verify/%s", code), nil)
}

// SendNotification lets you send notification of a payment request to your customers
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.SendNotification("<code>")
//
//	resp, err := prClient.SendNotification("<code>")
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
func (p *PaymentRequestClient) SendNotification(code string) (*Response, error) {
	return p.APICall(http.MethodPost, fmt.Sprintf("/paymentrequest/notify/%s", code), nil)
}

// Total lets you retrieve payment requests metric
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.Total()
//
//	resp, err := prClient.Total()
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
func (p *PaymentRequestClient) Total() (*Response, error) {
	return p.APICall(http.MethodGet, "/paymentrequest/totals", nil)
}

// Finalize lets you finalize a draft payment request
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.Finalize("<code>", true)
//
//	resp, err := prClient.Finalize("<code>", true)
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
func (p *PaymentRequestClient) Finalize(code string, sendNotification bool) (*Response, error) {
	payload := make(map[string]interface{})
	payload["send_notification"] = sendNotification
	return p.APICall(http.MethodPost, fmt.Sprintf("/paymentrequest/finalize/%s", code), nil)
}

// Update lets you update a payment request details on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.Update("<idOrCode>", "CUS_XXX", "description")
//
//	// you can pass in optional parameters to the `subAccounts.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `amount`.
//	// resp, err := saClient.Create("<idOrSlug>", "Buttercup Brunch", "description",
//	//	p.WithOptionalParameter("amount",500000))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/page/#update
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := prClient.Update("<idOrSlug>", "Buttercup Brunch", "description")
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
func (p *PaymentRequestClient) Update(idOrCode string, customerIdOrCode string,
	amount int, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = customerIdOrCode
	payload["amount"] = amount

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPut, fmt.Sprintf("/paymentrequest/%s", idOrCode), nil)
}

// Archive lets you archive a payment request. A payment request will no longer be fetched on list or returned on verify
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	prClient := p.NewPaymentRequestClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment request client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.paymentRequests field is a `PaymentRequestClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.paymentRequests.Archive("<idOrCode>")
//
//	resp, err := prClient.Archive("<idOrCode>")
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
func (p *PaymentRequestClient) Archive(idOrCode string) (*Response, error) {
	return p.APICall(http.MethodPost, fmt.Sprintf("/paymentrequest/archive/%s", idOrCode), nil)
}
