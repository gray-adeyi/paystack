package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// ChargeClient interacts with endpoints related to paystack charge resource that
// lets you configure a payment channel of your choice when initiating a payment.
type ChargeClient struct {
	*baseAPIClient
}

// NewChargeClient creates a ChargeClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	chargeClient := p.NewChargeClient(p.WithSecretKey("<paystack-secret-key>"))
func NewChargeClient(options ...ClientOptions) *ChargeClient {
	client := NewAPIClient(options...)
	return client.Charges
}

// Create lets you initiate a payment by integrating the payment channel of your choice.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	chargeClient := p.NewChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Charges field is a `ChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Charges.Create("johndoe@example.com", 100000)
//
//	// you can pass in optional parameters to the `Charges.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `amount`.
//	// resp, err := ppClient.Create("johndoe@example.com", 100000, p.WithOptionalParameter("authorization_code","AUTH_xxx"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/charge/#create
//	// Multiple optional parameters can be passed into `Create` each with it's `p.WithOptionalParameter`
//
// resp, err := chargeClient.Create("johndoe@example.com", 100000)
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
func (c *ChargeClient) Create(ctx context.Context, email string, amount string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"email":  email,
		"amount": amount,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPost, "/charge", payload)
}

// SubmitPin lets you submit pin to continue a charge
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	chargeClient := p.NewChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Charges field is a `ChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Charges.SubmitPin("1234", "5bwib5v6anhe9xa")
//
// resp, err := chargeClient.SubmitPin("1234", "5bwib5v6anhe9xa")
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
func (c *ChargeClient) SubmitPin(ctx context.Context, pin string, reference string) (*Response, error) {
	payload := map[string]any{
		"pin":       pin,
		"reference": reference,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_pin", payload)
}

// SubmitPhone lets you submit phone number when requested
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	chargeClient := p.NewChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Charges field is a `ChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Charges.SubmitPhone("08012345678", "5bwib5v6anhe9xa")
//
// resp, err := chargeClient.SubmitPhone("08012345678", "5bwib5v6anhe9xa")
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
func (c *ChargeClient) SubmitPhone(ctx context.Context, phone string, reference string) (*Response, error) {
	payload := map[string]any{
		"phone":     phone,
		"reference": reference,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_phone", payload)
}

// SubmitBirthday lets you submit a birthday when requested
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	chargeClient := p.NewChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Charges field is a `ChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Charges.SubmitBirthday("2016-09-21", "5bwib5v6anhe9xa")
//
// resp, err := chargeClient.SubmitBirthday("2016-09-21", "5bwib5v6anhe9xa")
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
func (c *ChargeClient) SubmitBirthday(ctx context.Context, birthday string, reference string) (*Response, error) {
	payload := map[string]any{
		"birthday":  birthday,
		"reference": reference,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_birthday", payload)
}

// SubmitAddress lets you submit address to continue a charge
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	chargeClient := p.NewChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Charges field is a `ChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Charges.SubmitAddress("140 N 2ND ST",
//	//	"7c7rpkqpc0tijs8", "Stroudsburg", "PA", "18360")
//
// resp, err := chargeClient.SubmitAddress("140 N 2ND ST", "7c7rpkqpc0tijs8", "Stroudsburg", "PA", "18360")
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
func (c *ChargeClient) SubmitAddress(ctx context.Context, address string, reference string, city string,
	state string, zipCode string) (*Response, error) {
	payload := map[string]any{
		"address":   address,
		"reference": reference,
		"city":      city,
		"state":     state,
		"zipcode":   zipCode,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_address", payload)
}

// PendingCharge lets you check the status of a charge. When you get pending as a charge status or if there
// was an exception when calling any of the /charge endpoints, wait 10 seconds or more, then make a check
// to see if its status has changed. Don't call too early as you may get a lot more pending than you should.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	chargeClient := p.NewChargeClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access the charge client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Charges field is a `ChargeClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Charges.PendingCharge("5bwib5v6anhe9xa")
//
// resp, err := chargeClient.PendingCharge("5bwib5v6anhe9xa")
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
func (c *ChargeClient) PendingCharge(ctx context.Context, reference string) (*Response, error) {
	return c.APICall(ctx, http.MethodGet, fmt.Sprintf("/charge/%s", reference), nil)
}
