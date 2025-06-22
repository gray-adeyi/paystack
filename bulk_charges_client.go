package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// BulkChargeClient interacts with endpoints related to paystack bulk Charges resource that lets
// you create and manage multiple recurring payments from your Customers.
type BulkChargeClient struct {
	*restClient
}

// NewBulkChargeClient creates a BulkChargeClient
func NewBulkChargeClient(options ...ClientOptions) *BulkChargeClient {
	client := NewClient(options...)
	return client.BulkCharges
}

// Initiate lets you send an array of map with authorization codes and amount, using the
// supported currency format, so paystack can process Transactions as a batch.
//
// Default response: models.Response[models.BulkCharge]
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
//		charges := []map[string]any{
//			{"authorization": "AUTH_ncx8hews93", "amount": 2500, "reference": "dam1266638dhhd"},
//			{"authorization": "AUTH_xfuz7dy4b9", "amount": 1500, "reference": "dam1266638dhhe"},
//		}
//		var response models.Response[models.BulkCharge]
//		if err := client.BulkCharges.Initiate(context.TODO(),charges, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (b *BulkChargeClient) Initiate(ctx context.Context, charges any, response any) error {
	return b.APICall(ctx, http.MethodPost, "/bulkcharge", charges, response)
}

// All lets you retrieve all bulk charge batches created by the Integration.
//
// Default response: models.Response[[]models.BulkCharge]
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
//		var response models.Response[[]models.BulkCharge]
//		if err := client.BulkCharges.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.BulkCharges.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/bulk-charge/
func (b *BulkChargeClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/bulkcharge", queries...)
	return b.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve a specific batch code. It also returns useful information
// on its progress by way of the `total_charges` and `pending_charges` attributes.
//
// Default response: models.Response[models.BulkCharge]
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
//		var response models.Response[models.BulkCharge]
//		if err := client.BulkCharges.FetchOne(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (b *BulkChargeClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return b.APICall(ctx, http.MethodGet, fmt.Sprintf("/bulkcharge/%s", idOrCode), nil, response)
}

// Charges lets you retrieve the Charges associated with a specified batch code.
//
// Default response: models.Response[[]models.BulkChargeUnitCharge]
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
//		var response models.Response[[]models.BulkChargeUnitCharge]
//		if err := client.BulkCharges.Charges(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.BulkCharges.Charges(context.TODO(),"<idOrCode>", &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/bulk-charge/
func (b *BulkChargeClient) Charges(ctx context.Context, idOrCode string, response any, queries ...Query) error {
	url := AddQueryParamsToUrl(fmt.Sprintf("/bulkcharge/%s/Charges", idOrCode), queries...)
	return b.APICall(ctx, http.MethodGet, url, nil, response)
}

// Pause lets you pause a processing a batch
//
// Default response: response models.Response[struct{}]
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
//		if err := client.BulkCharges.Pause(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (b *BulkChargeClient) Pause(ctx context.Context, idOrCode string, response any) error {
	return b.APICall(ctx, http.MethodGet, fmt.Sprintf("/bulkcharge/pause/%s", idOrCode), nil, response)
}

// Resume lets you resume a paused batch
//
// Default response: response models.Response[struct{}]
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
//		if err := client.BulkCharges.Resume(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (b *BulkChargeClient) Resume(ctx context.Context, idOrCode string, response any) error {
	return b.APICall(ctx, http.MethodGet, fmt.Sprintf("/bulkcharge/resume/%s", idOrCode), nil, response)
}
