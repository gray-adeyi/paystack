package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// SubAccountClient interacts with endpoints related to paystack subaccount resource that lets you
// create and manage subaccounts on your Integration. Subaccounts can be used to split payment
// between two accounts (your main account and a subaccount).
type SubAccountClient struct {
	*restClient
}

// NewSubAccountClient creates a SubAccountClient
func NewSubAccountClient(options ...ClientOptions) *SubAccountClient {
	client := NewClient(options...)

	return client.SubAccounts
}

// Create lets you create a dedicated virtual account for an existing customer
//
// Default response: models.Response[[]models.SubAccount]
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
//		var response models.Response[[]models.SubAccount]
//		if err := client.SubAccounts.Create(context.TODO(), "Sunshine Studios", "044", "0193274682", 18.2,"",&response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.SubAccounts.Create(context.TODO(), "Sunshine Studios", "044", "0193274682", 18.2,"",&response,p.WithOptionalParameter("primary_contact_email","johndoe@example.com"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/subaccount/
func (s *SubAccountClient) Create(ctx context.Context, businessName string, settlementBank string,
	accountNumber string, percentageCharge float32, description string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"business_name":     businessName,
		"settlement_bank":   settlementBank,
		"account_number":    accountNumber,
		"percentage_charge": percentageCharge,
		"description":       description,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(ctx, http.MethodPost, "/subaccount", payload, response)
}

// All lets you retrieve subaccounts available on your Integration
//
// Default response: models.Response[[]models.SubAccount]
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
//		var response models.Response[[]models.SubAccount]
//		if err := client.SubAccounts.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.SubAccounts.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/subaccount/
func (s *SubAccountClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/subaccount", queries...)
	return s.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve details of a subaccount on your Integration
//
// Default response: models.Response[models.SubAccount]
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
//		var response models.Response[models.SubAccount]
//		if err := client.SubAccounts.FetchOne(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (s *SubAccountClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return s.APICall(ctx, http.MethodGet, fmt.Sprintf("/subaccount/%s", idOrCode), nil, response)
}

// Update lets you update a subaccount details on your Integration
//
// Default response: models.Response[models.SubAccount]
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
//		if err := client.SubAccounts.Update(context.TODO(),"<idOrCode>", "Sunshine Studios", "044", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.SubAccounts.Update(context.TODO(),"<idOrCode>", "Sunshine Studios", "044", &response, p.WithOptionalParameter("primary_contact_email","johndoe@example.com"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/subaccount/
func (s *SubAccountClient) Update(ctx context.Context, idOrCode string, businessName string, settlementBank string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"business_name":   businessName,
		"settlement_bank": settlementBank,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(ctx, http.MethodPut, fmt.Sprintf("/subaccount/%s", idOrCode), payload, response)
}
