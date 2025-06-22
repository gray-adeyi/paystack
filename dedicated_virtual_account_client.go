package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// DedicatedVirtualAccountClient interacts with endpoints related to paystack dedicated virtual account
// resource that enables Nigerian merchants to manage unique payment accounts of their Customers.
type DedicatedVirtualAccountClient struct {
	*restClient
}

// NewDedicatedVirtualAccountClient creates a DedicatedVirtualAccountClient
func NewDedicatedVirtualAccountClient(options ...ClientOptions) *DedicatedVirtualAccountClient {
	client := NewClient(options...)

	return client.DedicatedVirtualAccounts
}

// Create lets you create a dedicated virtual account for an existing customer
//
// Default response: models.Response[models.DedicatedAccount]
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
//		var response models.Response[models.DedicatedAccount]
//		if err := client.DedicatedVirtualAccounts.Create(context.TODO(),"CUS_xr58yrr2ujlft9k", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.DedicatedVirtualAccounts.Create(context.TODO(),"CUS_xr58yrr2ujlft9k", &response, p.WithOptionalPayload("preferred_bank","wema-bank"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/dedicated-virtual-account/
func (d *DedicatedVirtualAccountClient) Create(ctx context.Context, customerIdOrCode string, response any,
	optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"customer": customerIdOrCode,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(ctx, http.MethodPost, "/dedicated_account", payload, response)
}

// Assign lets you can create a customer, validate the customer, and assign a DVA to the customer.
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
//		var response models.Response[models.DedicatedAccount]
//		if err := client.DedicatedVirtualAccounts.Assign(context.TODO(),"janedoe@test.com","Jane", "Doe","Karen", "+2348100000000", "test-bank", enum.CountryNigeria, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.DedicatedVirtualAccounts.Assign(context.TODO(),"janedoe@test.com","Jane", "Doe","Karen", "+2348100000000", "test-bank", enum.CountryNigeria, &response, p.WithOptionalPayload("account_number","5273681014"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/dedicated-virtual-account/
func (d *DedicatedVirtualAccountClient) Assign(ctx context.Context, email string, firstName string, lastName string,
	phone string, preferredBank string, country enum.Country, response any,
	optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"email":          email,
		"first_name":     firstName,
		"last_name":      lastName,
		"phone":          phone,
		"preferred_bank": preferredBank,
		"country":        country,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(ctx, http.MethodPost, "/dedicated_account/assign", payload, response)
}

// All lets you retrieve dedicated virtual accounts available on your Integration.
//
// Default response: models.Response[[]models.DedicatedAccount]
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
//		var response models.Response[[]models.DedicatedAccount]
//		if err := client.DedicatedVirtualAccounts.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.DedicatedVirtualAccounts.All(context.TODO(), &response,p.WithQuery("active","true"), p.WithQuery("bank_id","035"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/dedicated-virtual-account/
func (d *DedicatedVirtualAccountClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/dedicated_account", queries...)
	return d.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve details of a dedicated virtual account on your Integration.
//
// Default response: models.Response[models.DedicatedAccount]
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
//		var response models.Response[models.DedicatedAccount]
//		if err := client.DedicatedVirtualAccounts.FetchOne(context.TODO(),"<dedicatedAccountId>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (d *DedicatedVirtualAccountClient) FetchOne(ctx context.Context, dedicatedAccountId string, response any) error {
	return d.APICall(ctx, http.MethodGet, fmt.Sprintf("/dedicated_account/%s", dedicatedAccountId), nil, response)
}

// Requery lets you requery Dedicated Virtual Account for new Transactions
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
//		var response models.Response[[]models.DedicatedAccount]
//		if err := client.DedicatedVirtualAccounts.Requery(context.TODO(),"1234567890","wema-bank", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.DedicatedVirtualAccounts.Requery(context.TODO(),"1234567890","wema-bank", &response,p.WithQuery("date","2025-06-20"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/dedicated-virtual-account/
func (d *DedicatedVirtualAccountClient) Requery(ctx context.Context, accountNumber, providerSlug string, response any, queries ...Query) error {
	return d.APICall(ctx, http.MethodGet, fmt.Sprintf("/dedicated_account/requery?account_number=%s&prodiver_slug=%s", accountNumber, providerSlug), nil, response)
}

// Deactivate lets you deactivate a dedicated virtual account on your Integration.
//
// Default response: models.Response[models.DedicatedAccount]
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
//		var response models.Response[models.DedicatedAccount]
//		if err := client.DedicatedVirtualAccounts.Deactivate(context.TODO(),"<dedicatedAccountId>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (d *DedicatedVirtualAccountClient) Deactivate(ctx context.Context, id string, response any) error {
	return d.APICall(ctx, http.MethodDelete, fmt.Sprintf("/dedicated_account/%s", id), nil, response)
}

// Split lets you split a dedicated virtual account transaction with one or more accounts
//
// Default response: models.Response[models.DedicatedAccount]
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
//		var response models.Response[models.DedicatedAccount]
//		if err := client.DedicatedVirtualAccounts.Split(context.TODO(),"<customerIdOrCode>",&response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.DedicatedVirtualAccounts.Split(context.TODO(),"<customerIdOrCode>",p.WithOptionalPayload("preferred_bank","wema-bank"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/dedicated-virtual-account/
func (d *DedicatedVirtualAccountClient) Split(ctx context.Context, customerIdOrCode string, response any, optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"customer": customerIdOrCode,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}

	return d.APICall(ctx, http.MethodPost, "/dedicated_account/split", payload, response)
}

// RemoveSplit lets you remove a split payment for Transactions. If you've previously set up split payment
// for Transactions on a dedicated virtual account
//
// Default response: models.Response[models.DedicatedAccount]
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
//		var response models.Response[models.DedicatedAccount]
//		if err := client.DedicatedVirtualAccounts.RemoveSplit(context.TODO(),"<accountNumber>",&response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (d *DedicatedVirtualAccountClient) RemoveSplit(ctx context.Context, accountNumber string, response any) error {
	payload := map[string]any{
		"account_number": accountNumber,
	}
	return d.APICall(ctx, http.MethodDelete, "/dedicated_account/split", payload, response)
}

// BankProviders lets you retrieve available bank providers for a dedicated virtual account
//
// Default response: models.Response[[]models.DedicatedAccountProvider]
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
//		var response models.Response[[]models.DedicatedAccountProvider]
//		if err := client.DedicatedVirtualAccounts.BankProviders(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (d *DedicatedVirtualAccountClient) BankProviders(ctx context.Context, response any) error {
	return d.APICall(ctx, http.MethodPost, "/dedicated_account/available_providers", nil, response)
}
