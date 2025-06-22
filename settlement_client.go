package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// SettlementClient interacts with endpoints related to paystack settlement resource that lets you
// gain insights into payouts made by Paystack to your bank account.
type SettlementClient struct {
	*restClient
}

// NewSettlementClient creates a SettlementClient
func NewSettlementClient(options ...ClientOptions) *SettlementClient {
	client := NewClient(options...)
	return client.Settlements
}

// All lets you retrieve Settlements made to your settlement accounts
//
// Default response: models.Response[[]models.Settlement]
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
//		var response models.Response[[]models.Settlement]
//		if err := client.Settlements.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Settlements.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/settlement/
func (s *SettlementClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/settlement", queries...)
	return s.APICall(ctx, http.MethodGet, url, nil, response)
}

// AllTransactions lets you retrieve the Transactions that make up a particular settlement
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
//		if err := client.Settlements.AllTransactions(context.TODO(), "<settlementId>",&response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Settlements.AllTransactions(context.TODO(),"<settlementId>", &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/settlement/
func (s *SettlementClient) AllTransactions(ctx context.Context, settlementId string, response any, queries ...Query) error {
	url := AddQueryParamsToUrl(fmt.Sprintf("/settlement/%s", settlementId), queries...)
	return s.APICall(ctx, http.MethodGet, url, nil, response)
}
