package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// PaymentRequestClient interacts with endpoints related to paystack payment request resource that
// lets you manage requests for payment of goods and services.
type PaymentRequestClient struct {
	*restClient
}

// NewPaymentRequestClient creates a PaymentRequestClient
func NewPaymentRequestClient(options ...ClientOptions) *PaymentRequestClient {
	client := NewClient(options...)
	return client.PaymentRequests
}

// Create lets you create a payment request for a transaction on your Integration
//
// Default response: models.Response[models.PaymentRequest]
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
//		if err := client.PaymentRequests.Create(context.TODO(),"CUS_xwaj0txjryg393b", 500000, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.PaymentRequests.Create(context.TODO(),"CUS_xwaj0txjryg393b", 500000, &response, p.WithOptionalPayload("due_date","2023-12-25"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/payment-request/
func (p *PaymentRequestClient) Create(ctx context.Context, customerIdOrCode string, amount int, response any,
	optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"customer": customerIdOrCode,
		"amount":   amount,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPost, "/paymentrequest", payload, response)
}

// All lets you retrieve the payment requests available on your Integration
//

// Default response: models.Response[[]models.PaymentRequest]
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
//		var response models.Response[[]models.PaymentRequest]
//		if err := client.Customers.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.PaymentRequests.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/payment-request/
func (p *PaymentRequestClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/paymentrequest", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve details of a payment request on your Integration
//
// Default response: models.Response[models.PaymentRequest]
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
//		var response models.Response[models.PaymentRequest]
//		if err := client.PaymentRequests.FetchOne(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentRequestClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/paymentrequest/%s", idOrCode), nil, response)
}

// Verify lets you verify the details of a payment request on your Integration
//
// Default response: models.Response[models.PaymentRequest]
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
//		var response models.Response[models.PaymentRequest]
//		if err := client.PaymentRequests.Verify(context.TODO(),"<code>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentRequestClient) Verify(ctx context.Context, code string, response any) error {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/paymentrequest/verify/%s", code), nil, response)
}

// SendNotification lets you send notification of a payment request to your Customers
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
//		if err := client.PaymentRequests.SendNotification(context.TODO(),"<code>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentRequestClient) SendNotification(ctx context.Context, idOrCode string, response any) error {
	return p.APICall(ctx, http.MethodPost, fmt.Sprintf("/paymentrequest/notify/%s", idOrCode), nil, response)
}

// Total lets you retrieve payment requests metric
//
// Default response: models.Response[models.PaymentRequestStat]
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
//		var response models.Response[models.PaymentRequestStat]
//		if err := client.PaymentRequests.Total(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentRequestClient) Total(ctx context.Context, response any) error {
	return p.APICall(ctx, http.MethodGet, "/paymentrequest/totals", nil, response)
}

// Finalize lets you finalize a draft payment request
//
// Default response: models.Response[models.PaymentRequest]
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
//		var response models.Response[models.PaymentRequest]
//		if err := client.PaymentRequests.Finalize(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentRequestClient) Finalize(ctx context.Context, idOrCode string, sendNotification bool, response any) error {
	payload := map[string]any{
		"send_notification": sendNotification,
	}
	return p.APICall(ctx, http.MethodPost, fmt.Sprintf("/paymentrequest/finalize/%s", idOrCode), payload, response)
}

// Update lets you update a payment request details on your Integration
//
// Default response: models.Response[models.PaymentRequest]
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
//		var response models.Response[models.PaymentRequest]
//		if err := client.PaymentRequests.Update(context.TODO(),"<idOrCode>", 50000, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.PaymentRequests.Update(context.TODO(),"<idOrCode>", 50000, &response, p.WithOptionalPayload("description","update description"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/payment-request/
func (p *PaymentRequestClient) Update(ctx context.Context, idOrCode string, customerIdOrCode string,
	amount int, response any, optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"customer": customerIdOrCode,
		"amount":   amount,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPut, fmt.Sprintf("/paymentrequest/%s", idOrCode), payload, response)
}

// Archive lets you archive a payment request. A payment request will no longer be fetched on list or returned on verify
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
//		if err := client.PaymentRequests.Archive(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentRequestClient) Archive(ctx context.Context, idOrCode string, response any) error {
	return p.APICall(ctx, http.MethodPost, fmt.Sprintf("/paymentrequest/archive/%s", idOrCode), nil, response)
}
