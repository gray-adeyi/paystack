package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// TransactionClient interacts with endpoints related to paystack Transaction resource
// that allows you to create and manage payments on your Integration.
type TransactionClient struct {
	*restClient
}

// NewTransactionClient creates a TransactionClient
func NewTransactionClient(options ...ClientOptions) *TransactionClient {
	client := NewClient(options...)

	return client.Transactions

}

// Initialize lets you initialize a transaction from your backend
//
// Default response: models.Response[models.InitTransaction]
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
//		var response models.Response[models.InitTransaction]
//		if err := client.Transactions.Initialize(context.TODO(),200000, "johndoe@example.com", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Transactions.Initialize(context.TODO(),200000, "johndoe@example.com", &response, p.WithOptionalParameter("currency", string(enum.CurrencyNgn)))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/transaction/
func (t *TransactionClient) Initialize(ctx context.Context, amount int, email string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"amount": amount,
		"email":  email,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transaction/initialize", payload, response)
}

// Verify lets you confirm the status of a transaction
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Transactions.Verify(context.TODO(),"<reference>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransactionClient) Verify(ctx context.Context, reference string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transaction/verify/%s", reference), nil, response)
}

// All lets you list Transactions carried out on your Integration
//
// Default response: models.Response[[]models.Transaction]
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
//		var response models.Response[[]models.Transaction]
//		if err := client.Transactions.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Transactions.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/transaction/
func (t *TransactionClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/transaction", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you get the details of a transaction carried out on your Integration
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Transactions.FetchOne(context.TODO(),"<id>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransactionClient) FetchOne(ctx context.Context, id string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transaction/%s", id), nil, response)
}

// ChargeAuthorization lets you charge authorizations that are marked as reusable
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Transactions.ChargeAuthorization(context.TODO(),200000,"johndoe@example.com","AUTH_xxx", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Transactions.ChargeAuthorization(context.TODO(),200000,"johndoe@example.com","AUTH_xxx", &response, p.WithOptionalParameter("channel","bank"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/transaction/
func (t *TransactionClient) ChargeAuthorization(ctx context.Context, amount int, email string, authorizationCode string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"amount":             amount,
		"email":              email,
		"authorization_code": authorizationCode,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transaction/charge_authorization", payload, response)
}

// Timeline lets you view the timeline of a transaction
//
// Default response: models.Response[models.TransactionLog]
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
//		var response models.Response[models.TransactionLog]
//		if err := client.Transactions.Timeline(context.TODO(),"<idOrReference>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransactionClient) Timeline(ctx context.Context, idOrReference string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transaction/timeline/%s", idOrReference), nil, response)
}

// Total lets you retrieve the total amount received on your account
//
// Default response: models.Response[[]models.TransactionTotal]
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
//		var response models.Response[[]models.TransactionTotal]
//		if err := client.Transactions.Total(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Transactions.Total(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/transaction/
func (t *TransactionClient) Total(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/transaction/totals", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// Export lets you export a list of Transactions carried out on your Integration
//
// Default response: models.Response[[]models.TransactionExport]
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
//		var response models.Response[[]models.TransactionExport]
//		if err := client.Transactions.Export(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Transactions.Export(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/transaction/
func (t *TransactionClient) Export(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/transaction/export", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// PartialDebit lets you retrieve part of a payment from a customer
//
// Default response: models.Response[[]models.Transaction]
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
//		var response models.Response[[]models.Transaction]
//		if err := client.Transactions.PartialDebit(context.TODO(),"AUTH_xxx",enum.CurrencyNgn,"200000", "johndoe@example.com", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Transactions.PartialDebit(context.TODO(), "AUTH_xxx",enum.CurrencyNgn,"200000", "johndoe@example.com", &response,p.WithOptionalParameter("at_least",100000))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/transaction/
func (t *TransactionClient) PartialDebit(ctx context.Context, authorizationCode string, currency enum.Currency, amount string, email string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"authorization_code": authorizationCode,
		"currency":           currency,
		"amount":             amount,
		"email":              email,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transaction/partial_debit", payload, response)
}
