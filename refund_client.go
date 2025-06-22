package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// RefundClient interacts with endpoints related to paystack refund resource that lets you
// create and manage transaction Refunds.
type RefundClient struct {
	*restClient
}

// NewRefundClient creates a RefundClient
func NewRefundClient(options ...ClientOptions) *RefundClient {
	client := NewClient(options...)
	return client.Refunds
}

// Create lets you create and manage transaction Refunds.
// 
// Default response: models.Response[models.Refund]
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
//		var response models.Response[models.Refund]
//		if err := client.Refunds.Create(context.TODO(),"1641", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Refunds.Create(context.TODO(),"1641", &response, p.WithOptionalPayload("amount",500000))
//	}
// For supported optional parameters, see:
// https://paystack.com/docs/api/refund/
func (r *RefundClient) Create(ctx context.Context, transaction string, response any,
	optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"transaction": transaction,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return r.APICall(ctx, http.MethodPost, "/refund", payload, response)
}

// All lets you retrieve Refunds available on your Integration
//
// Default response: models.Response[[]models.Refund]
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
//		var response models.Response[[]models.Refund]
//		if err := client.Refunds.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Refunds.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/refund/
func (r *RefundClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/refund", queries...)
	return r.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve the details of a refund on your Integration
//
// Default response: models.Response[models.Refund]
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
//		var response models.Response[models.Refund]
//		if err := client.Plans.FetchOne(context.TODO(),"<reference>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (r *RefundClient) FetchOne(ctx context.Context, reference string, response any) error {
	return r.APICall(ctx, http.MethodGet, fmt.Sprintf("/refund/%s", reference), nil, response)
}
