package paystack

import (
	"fmt"
	"net/http"
)

// DisputeClient interacts with endpoint related to paystack dispute resource that lets you
// manage transaction disputes on your integration.
type DisputeClient struct {
	*baseAPIClient
}

// NewDisputeClient creates a DisputeClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
func NewDisputeClient(options ...ClientOptions) *DisputeClient {
	client := NewAPIClient(options...)
	return client.disputes
}

// All lets you retrieve disputes filed against you
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dispute client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.disputes field is a `DisputeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.disputes.All()
//
//	// All also accepts queries, so say you want to specify the date range, you can write it like so.
//	// resp, err := dClient.All(p.WithQuery("from","2023-01-01"), p.WithQuery("to","2023-12-31"))
//
// // see https://paystack.com/docs/api/dispute/#list for supported query parameters
//
//	resp, err := dClient.All()
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
func (d *DisputeClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/dispute", queries...)
	return d.APICall(http.MethodGet, url, nil)
}

// FetchOne lets you retrieve more details about a dispute.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a payment page client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.disputes field is a `DisputeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.disputes.FetchOne("<id>")
//
//	resp, err := dClient.FetchOne("<id>")
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
func (d *DisputeClient) FetchOne(id string) (*Response, error) {
	return d.APICall(http.MethodGet, fmt.Sprintf("/dispute/%s", id), nil)
}

// AllTransactionDisputes lets you retrieve disputes for a particular transaction
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dispute client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.disputes field is a `DisputeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.disputes.AllTransactionDisputes("transactionId")
//
//	resp, err := dClient.AllTransactionDisputes("transactionId")
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
func (d *DisputeClient) AllTransactionDisputes(transactionId string) (*Response, error) {
	return d.APICall(http.MethodGet, fmt.Sprintf("/dispute/transaction/%s", transactionId), nil)
}

// Update lets you update the details of a dispute on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dispute client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.disputes field is a `DisputeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.disputes.Update("<id>", 1002)
//
//	// you can pass in optional parameters to the `disputes.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `uploaded_filename`.
//	// resp, err := dClient.Update("<id>", 1002, "description",
//	//	p.WithOptionalParameter("uploaded_filename","disputes.pdf"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/dispute/#update
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//
//	resp, err := dClient.Update("<id>", 1002)
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
func (d *DisputeClient) Update(id string, referenceAmount int,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"reference_amount": referenceAmount,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(http.MethodPut, fmt.Sprintf("/dispute/%s", id), payload)
}

// AddEvidence lets you provide evidence for a dispute
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dispute client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.disputes field is a `DisputeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.disputes.AddEvidence("<id>", "johndoe@example.com",
//	//	"John Doe", "5085072209", "claim for buying product")
//
//	// you can pass in optional parameters to the `disputes.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `delivery_address`.
//	// resp, err := dClient.AddEvidence("<id>", "johndoe@example.com", "John Doe", "5085072209", "claim for buying product",
//	//	p.WithOptionalParameter("delivery_address", "3a ladoke street ogbomoso"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/dispute/#evidence
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//
//	resp, err := dClient.AddEvidence("<id>", "johndoe@example.com", "John Doe", "5085072209", "claim for buying product")
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
func (d *DisputeClient) AddEvidence(id string, customerEmail string,
	customerName string, customerPhone string, serviceDetails string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"customer_email":  customerEmail,
		"customer_name":   customerName,
		"customer_phone":  customerPhone,
		"service_details": serviceDetails,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(http.MethodPost, fmt.Sprintf("/dispute/%s/evidence", id), payload)
}

// UploadURL lets you retrieve disputes for a particular transaction
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dispute client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.disputes field is a `DisputeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.disputes.UploadURL()
//
//	// All also accepts queries, so say you want to specify the `upload_filename`, you can write it like so.
//	// resp, err := dClient.UploadURL("disputeId", p.WithQuery("upload_filename","filename.txt"))
//
// // see https://paystack.com/docs/api/dispute/#upload-url for supported query parameters
//
//	resp, err := dClient.UploadURL()
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
func (d *DisputeClient) UploadURL(id string, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl(fmt.Sprintf("/dispute/%s/upload_url", id), queries...)
	return d.APICall(http.MethodPost, url, nil)
}

// Resolve lets you resolve a dispute on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dispute client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.disputes field is a `DisputeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.disputes.AddEvidence("<id>", "johndoe@example.com",
//	//	"John Doe", "5085072209", "claim for buying product")
//
//	// you can pass in optional parameters to the `disputes.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `evidence`.
//	// resp, err := dClient.AddEvidence("<id>", "johndoe@example.com","John Doe", "5085072209", "claim for buying product",
//	//	p.WithOptionalParameter("evidence", "evidenceId"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/dispute/#evidence
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//
//	resp, err := dClient.AddEvidence("<id>", "merchant-accepted", "Merchant accepted", 10000, "resolve.pdf")
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
func (d *DisputeClient) Resolve(id string, resolution string, message string,
	refundAmount int, uploadedFilename string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"resolution":        resolution,
		"message":           message,
		"refund_amount":     refundAmount,
		"uploaded_filename": uploadedFilename,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(http.MethodPut, fmt.Sprintf("/dispute/%s/resolve", id), payload)
}

// Export lets you export disputes available on your integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	dClient := p.NewDisputeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a dispute client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.disputes field is a `DisputeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.disputes.Export()
//
//	// All also accepts queries, so say you want to specify the date range, you can write it like so.
//	// resp, err := dClient.Export(p.WithQuery("from","2023-01-01"), p.WithQuery("to","2023-12-31"))
//
// // see https://paystack.com/docs/api/dispute/#export for supported query parameters
//
//	resp, err := dClient.Export()
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
func (d *DisputeClient) Export(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/dispute/export", queries...)
	return d.APICall(http.MethodGet, url, nil)
}
