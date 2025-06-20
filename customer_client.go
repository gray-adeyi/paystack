package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// CustomerClient interacts with endpoints related to paystack Customer resource
// that allows you to create and manage Customers on your Integration.
type CustomerClient struct {
	*restClient
}

// NewCustomerClient creates a CustomerClient
func NewCustomerClient(options ...ClientOptions) *CustomerClient {
	client := NewClient(options...)

	return client.Customers
}

// Create lets you create a customer on your Integration
//
// Default response: models.Response[models.Customer]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.Customer]
//		if err := client.Customers.Create(context.TODO(),"johndoe@example.com","John","Doe", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Customers.Create(context.TODO(),"johndoe@example.com","John","Doe", &response, p.WithOptionalParameter("phone","+2348123456789"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/customer/
func (c *CustomerClient) Create(ctx context.Context, email string, firstName string, lastName string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"email":      email,
		"first_name": firstName,
		"last_name":  lastName,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPost, "/customer", payload, response)
}

// All lets you retrieve Customers available on your Integration.
//
// Default response: models.Response[[]models.Customer]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[[]models.Customer]
//		if err := client.Customers.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Customers.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/customer/
func (c *CustomerClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/terminal", queries...)
	return c.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve the details of a customer on your Integration
//
// Default response: models.Response[models.Customer]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.Customer]
//		if err := client.Customers.FetchOne(context.TODO(),"<emailOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (c *CustomerClient) FetchOne(ctx context.Context, emailOrCode string, response any) error {
	return c.APICall(ctx, http.MethodGet, fmt.Sprintf("/customer/%s", emailOrCode), nil, response)
}

// Update lets you update a customer's details on your Integration
//
// Default response: models.Response[models.Customer]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.Customer]
//		if err := client.Customers.Update(context.TODO(),"<code>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Customers.Update(context.TODO(),"<code>", &response, p.WithOptionalParameter("first_name","John"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/customer/
func (c *CustomerClient) Update(ctx context.Context, code string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := make(map[string]any)

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPut, fmt.Sprintf("/customer/%s", code), payload, response)
}

// Validate lets you validate a customer's identity
//
// Default response: models.Response[struct{}]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//		"github.com/gray-adeyi/paystack/enum"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[struct{}]
//		if err := client.Customers.Validate(context.TODO(),"<code>", "Asta","Lavista","bank_account","",enum.CountryNigeria,"20012345677","007","0123456789", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Customers.Validate(context.TODO(),"<code>", "Asta","Lavista","bank_account","",enum.CountryNigeria,"20012345677","007","0123456789", &response, , p.WithOptionalParameter("middle_name","Doe"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/customer/
func (c *CustomerClient) Validate(ctx context.Context, code string, firstName string, lastName string, identificationType string,
	value string, country enum.Country, bvn string, bankCode string, accountNumber string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
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

	return c.APICall(ctx, http.MethodPost, fmt.Sprintf("/customer/%s/identification", code), payload, response)
}

// Flag lets you whitelist or blacklist a customer on your Integration
//
// Default response: models.Response[models.Customer]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.Customer]
//		if err := client.Customers.Flag(context.TODO(),"CUS_xr58yrr2ujlft9k", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Customers.Flag(context.TODO(),"CUS_xr58yrr2ujlft9k", &response, p.WithOptionalParameter("risk_action","allow"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/customer/
func (c *CustomerClient) Flag(ctx context.Context, emailOrCode string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"cutomer": emailOrCode,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPost, "/customer/set_risk_action", payload, response)
}

// Deactivate lets you deactivate an authorization when the card needs to be forgotten
//
// Default response: models.Response[struct{}]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[struct{}]
//		if err := client.Customers.Deactivate(context.TODO(),"AUTH_72btv547", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (c *CustomerClient) Deactivate(ctx context.Context, authorizationCode string, response any) error {
	payload := map[string]any{
		"authorization_code": authorizationCode,
	}

	return c.APICall(ctx, http.MethodPost, "/customer/deactivate_authorization", payload, response)
}
