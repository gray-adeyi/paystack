package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// CustomerClient interacts with endpoints related to paystack Customer resource
// that allows you to create and manage Customers on your Integration.
type CustomerClient struct {
	*baseAPIClient
}

// NewCustomerClient creates a CustomerClient
//
// Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	customerClient := p.NewCustomerClient(p.WithSecretKey("<paystack-secret-key>"))
func NewCustomerClient(options ...ClientOptions) *CustomerClient {
	client := NewAPIClient(options...)

	return client.Customers
}

// Create lets you create a customer on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//		"context"
//	)
//
//	customerClient := p.NewCustomerClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Customers field is a `CustomerClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Customers.Create(context.TODO(),"johndoe@example.com","John","Doe")
//
//	// you can pass in optional parameters to the `customerClient.Create` with `p.WithOptionalParameter`
//	// for example say you want to specify the `phone`.
//	// resp, err := customerClient.Create("johndoe@example.com","John","Doe", p.WithOptionalParameter("phone","+2348123456789"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/customer/#create
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//
//	resp, err := customerClient.Create(context.TODO(),"johndoe@example.com","John","Doe")
//
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (c *CustomerClient) Create(ctx context.Context, email string, firstName string, lastName string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"email":      email,
		"first_name": firstName,
		"last_name":  lastName,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPost, "/customer", payload)
}

// All lets you retrieve Customers available on your Integration.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//		"context"
//	)
//
//	customerClient := p.NewCustomerClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Customers field is a `CustomerClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Customers.All(context.TODO())
//
//	// All also accepts queries, so say you want to specify the page
//	// to retrieve, you can write it like so.
//	// resp, err := customerClient.All(context.TODO(),p.WithQuery("page",2))
//
//	// see https://paystack.com/docs/api/customer/#list for supported query parameters
//
//	resp, err := customerClient.All(context.TODO())
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (c *CustomerClient) All(ctx context.Context, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/terminal", queries...)
	return c.APICall(ctx, http.MethodGet, url, nil)
}

// FetchOne lets you retrieve the details of a customer on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//		"context"
//	)
//
//	customerClient := p.NewCustomerClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Customers field is a `CustomerClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Customers.FetchOne(context.TODO(),"<emailOrCode>")
//
//	resp, err := customerClient.FetchOne(context.TODO(),"<emailOrCode>")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (c *CustomerClient) FetchOne(ctx context.Context, emailOrCode string) (*Response, error) {
	return c.APICall(ctx, http.MethodGet, fmt.Sprintf("/customer/%s", emailOrCode), nil)
}

// Update lets you update a customer's details on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//		"context"
//	)
//
//	customerClient := p.NewCustomerClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Customers field is a `CustomerClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Customers.Update(context.TODO(),"143", p.WithOptionalParameter("first_name","John"))
//
//	// you can pass in optional parameters to the `customerClient.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `first_name`.
//	// resp, err := customerClient.Update(context.TODO(),"143", p.WithOptionalParameter("first_name","John"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/customer/#update
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//	resp, err := customerClient.Update(context.TODO(),"143", p.WithOptionalParameter("first_name","John"))
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (c *CustomerClient) Update(ctx context.Context, code string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]any)

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPut, fmt.Sprintf("/customer/%s", code), payload)
}

// Validate lets you validate a customer's identity
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//		"context"
//	)
//
//	customerClient := p.NewCustomerClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Customers field is a `CustomerClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Customers.Validate(context.TODO(),"143", "Asta","Lavista","bank_account","","NG","20012345677","007","0123456789")
//
//	// you can pass in optional parameters to the `customerClient.Validate` with `p.WithOptionalParameter`
//	// for example say you want to specify the `middle_name`.
//	// resp, err := customerClient.Validate(context.TODO(),"143", "Asta","Lavista","bank_account","","NG","20012345677","007",
//	//	"0123456789", p.WithOptionalParameter("middle_name","Doe"))
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/customer/#validate
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//	resp, err := customerClient.Validate(context.TODO(),"143", "Asta","Lavista","bank_account","","NG","20012345677","007","0123456789")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (c *CustomerClient) Validate(ctx context.Context, code string, firstName string, lastName string, identificationType string,
	value string, country string, bvn string, bankCode string, accountNumber string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"first_name":     firstName,
		"last_name":      lastName,
		"type":           identificationType,
		"value":          value,
		"country":        country,
		"bvn":            bvn,
		"bank_code":      bankCode,
		"account_number": accountNumber,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPost, fmt.Sprintf("/customer/%s/identification", code), payload)
}

// Flag lets you whitelist or blacklist a customer on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//		"context"
//	)
//
//	customerClient := p.NewCustomerClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Customers field is a `CustomerClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Customers.Flag(context.TODO(),"CUS_xr58yrr2ujlft9k")
//
//	// you can pass in optional parameters to the `customerClient.Update` with `p.WithOptionalParameter`
//	// for example say you want to specify the `risk_action`.
//	// resp, err := customerClient.Flag(context.TODO(),"CUS_xr58yrr2ujlft9k", p.WithOptionalParameter("risk_action", "allow")
//	// the `p.WithOptionalParameter` takes in a key and value parameter, the key should match the optional parameter
//	// from paystack documentation see https://paystack.com/docs/api/customer/#whitelist-blacklist
//	// Multiple optional parameters can be passed into `Update` each with it's `p.WithOptionalParameter`
//
//	resp, err := customerClient.Flag(context.TODO(),"CUS_xr58yrr2ujlft9k")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (c *CustomerClient) Flag(ctx context.Context, emailOrCode string,
	optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]any{
		"cutomer": emailOrCode,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPost, "/customer/set_risk_action", payload)
}

// Deactivate lets you deactivate an authorization when the card needs to be forgotten
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//		"context"
//	)
//
//	customerClient := p.NewCustomerClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a customer client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Customers field is a `CustomerClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Customers.Deactivate(context.TODO(),"AUTH_72btv547")
//
//	resp, err := customerClient.Deactivate(context.TODO(),"AUTH_72btv547")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]any` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]any)
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (c *CustomerClient) Deactivate(ctx context.Context, authorizationCode string) (*Response, error) {
	payload := map[string]any{
		"authorization_code": authorizationCode,
	}

	return c.APICall(ctx, http.MethodPost, "/customer/deactivate_authorization", payload)
}
