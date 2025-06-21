package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// TransactionSplitClient interacts with endpoints related to paystack Transaction Split resource
// that allows you to split the settlement for a transaction across a payout account, and one or
// more subaccounts.
type TransactionSplitClient struct {
	*restClient
}

// NewTransactionSplitClient creates a TransactionSplitClient
func NewTransactionSplitClient(options ...ClientOptions) *TransactionSplitClient {
	client := NewClient(options...)
	return client.TransactionSplits
}

// Create lets you create a split payment on your Integration
// 
// Default response: models.Response[models.TransactionSplit]
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
// 		subaccounts := []map[string]any{
//			{"subaccount": "ACCT_z3x6z3nbo14xsil", "share": 20},
// 			{"subaccount": "ACCT_pwwualwty4nhq9d", "share": 80},
// 		}
//		var response models.Response[models.TransactionSplit]
//		if err := client.TransactionSplits.Create(context.TODO(),"co-founders account",enum.SplitPercentage,enum.CurrencyNgn,subaccounts, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Customers.Create(context.TODO(),"co-founders account",enum.SplitPercentage,enum.CurrencyNgn,subaccounts, &response, p.WithOptionalParameter("bearer_type","all"))
//	}
// 
// For supported optional parameters, see:
// https://paystack.com/docs/api/split/
func (t *TransactionSplitClient) Create(ctx context.Context, name string, transactionSplitType enum.Split, currency enum.Currency, subaccounts, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":              name,
		"type":              transactionSplitType,
		"currency":          currency,
		"subaccounts":       subaccounts,
		// "bearer_type":       bearerType, // These fields seem to be optional and should be passed as optional parameters
		// "bearer_subaccount": bearerSubaccount,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/split", payload, response)
}

// All let you list the transaction splits available on your Integration
//
// Default response: models.Response[[]models.TransactionSplit]
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
//		var response models.Response[[]models.TransactionSplit]
//		if err := client.TransactionSplits.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.TransfersSplits.All(context.TODO(), &response,p.WithQuery("name","co-founders account"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/split/
func (t *TransactionSplitClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/split", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you get the details of a split on your Integration
//
// Default response: models.Response[models.TransactionSplit]
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
//		var response models.Response[models.TransactionSplit]
//		if err := client.TransactionSplits.FetchOne(context.TODO(),"<id>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransactionSplitClient) FetchOne(ctx context.Context, id string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/split/%s", id), nil, response)
}

// Update lets you update a transaction split details on your Integration
//
// Default response: models.Response[models.TransactionSplit]
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
//		var response models.Response[models.TransactionSplit]
//		if err := client.TransactionSplits.Update(context.TODO(), "143", "co-authors account", true,&response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.TransfersSplits.Update(context.TODO(),"143", "co-authors account", true, &response,p.WithOptionalParameter("bearer_type","all"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/split/
func (t *TransactionSplitClient) Update(ctx context.Context, id string, name string, active bool, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":   name,
		"active": active,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return t.APICall(ctx, http.MethodPut, fmt.Sprintf("/split/%s", id), payload, response)
}

// Add lets you add a Subaccount to a Transaction Split, or update the share of an existing
// Subaccount in a Transaction Split
//
// Default response: models.Response[models.TransactionSplit]
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
//		var response models.Response[models.TransactionSplit]
//		if err := client.TransactionSplits.Add(context.TODO(),"ACCT_hdl8abxl8drhrl3", 15, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransactionSplitClient) Add(ctx context.Context, id string, subAccount string, share int, response any) error {
	payload := map[string]any{
		"subaccount": subAccount,
		"share":      share,
	}
	return t.APICall(ctx, http.MethodPost, fmt.Sprintf("/split/%s/add", id), payload, response)
}

// Remove lets you remove a subaccount from a transaction split
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
//		if err := client.TransactionSplits.Remove(context.TODO(),"143","ACCT_hdl8abxl8drhrl3", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransactionSplitClient) Remove(ctx context.Context, id string, subAccount string, response any) error {
	payload := map[string]any{
		"subaccount": subAccount,
	}

	return t.APICall(ctx, http.MethodPost, fmt.Sprintf("/split/%s/remove", id), payload, response)
}
