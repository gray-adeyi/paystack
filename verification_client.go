package paystack

import (
	"fmt"
	"net/http"
)

// VerificationClient interacts with endpoints related to paystack Verification resource
// that allows you to perform KYC processes.
type VerificationClient struct {
	*baseAPIClient
}

// NewVerificationClient creates a VerificationClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	vClient := p.NewVerificationClient(p.WithSecretKey("<paystack-secret-key>"))
func NewVerificationClient(options ...ClientOptions) *VerificationClient {
	client := NewAPIClient(options...)
	return client.Verification
}

// ResolveAccount lets you confirm an account belongs to the right customer
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	vClient := p.NewVerificationClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a Verification client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Verification field is a `VerificationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Verification.ResolveAccount(p.WithQuery("account_number","0022728151"),
//	//	p.WithQuery("bank_code","063"))
//
//	// All also accepts queries, you can write it like so.
//	// resp, err := paystackClient.Verification.ResolveAccount(p.WithQuery("account_number","0022728151"),
//	//	p.WithQuery("bank_code","063"))
//
// // see https://paystack.com/docs/api/verification/#resolve-account for supported query parameters
//
//	resp, err := vClient.ResolveAccount(p.WithQuery("account_number","0022728151"),
//	p.WithQuery("bank_code","063"))
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
func (v *VerificationClient) ResolveAccount(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/bank/resolve", queries...)
	return v.APICall(http.MethodGet, url, nil)
}

// ValidateAccount lets you confirm the authenticity of a customer's account number before sending money
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	vClient := p.NewVerificationClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a Verification client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Verification field is a `VerificationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Verification.ValidateAccount("Ann Bron","0123456789","personal",
//	//	"632005","ZA","identityNumber")
//
//	// you can pass in optional parameters to the `Verification.ValidateAccount` with `p.WithOptionalParameter`
//	// for example say you want to specify the `document_number`.
//	// resp, err := vClient.CreateValidateAccount("Ann Bron","0123456789","personal", "632005","ZA",
//	//	"identityNumber", p.WithOptionalParameter("document_number","1234567890123"))
//
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/verification/#validate-account
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
//	resp, err := vClient.ValidateAccount("Ann Bron","0123456789","personal","632005","ZA","identityNumber")
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
func (v *VerificationClient) ValidateAccount(accountName string, accountNumber string, accountType string, bankCode string, countryCode string, documentType string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"account_name":   accountName,
		"account_number": accountNumber,
		"account_type":   accountType,
		"bank_code":      bankCode,
		"country_code":   countryCode,
		"document_type":  documentType,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return v.APICall(http.MethodGet, "/bank/validate", payload)
}

// ResolveBIN lets you retrieve more information about a customer's card
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	vClient := p.NewVerificationClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a Verification client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Verification field is a `VerificationClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Verification.ResolveBIN("539983")
//
//	resp, err := vClient.ResolveAccount("539983")
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
func (v *VerificationClient) ResolveBIN(bin string) (*Response, error) {
	return v.APICall(http.MethodGet, fmt.Sprintf("/decision/bin/%s", bin), nil)
}
