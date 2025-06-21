package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// SubscriptionClient interacts with endpoints related to paystack subscription resource that lets you
// create and manage recurring payment on your Integration.
type SubscriptionClient struct {
	*restClient
}

// NewSubscriptionClient create a SubscriptionClient
func NewSubscriptionClient(options ...ClientOptions) *SubscriptionClient {
	client := NewClient(options...)

	return client.Subscriptions
}

// Create lets you create a subscription on your Integration
//
// Default response: models.Response[models.Subscription]
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
//		var response models.Response[models.Subscription]
//		if err := client.Subscriptions.Create(context.TODO(),"CUS_xnxdt6s1zg1f4nx", "PLN_gx2wn530m0i3w3m", "AUTH_xxx", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Subscriptions.Create(context.TODO(),"CUS_xnxdt6s1zg1f4nx", "PLN_gx2wn530m0i3w3m", "AUTH_xxx", &response, p.WithOptionalParameter("start_date","2023-10-16T00:30:13+01:00"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/subscription/
func (s *SubscriptionClient) Create(ctx context.Context, customer string, plan string, authorization string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"customer":      customer,
		"plan":          plan,
		"authorization": authorization,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(ctx, http.MethodPost, "/subscription", payload, response)
}

// All lets you retrieve Subscriptions available on your Integration
//
// Default response: models.Response[[]models.Subscription]
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
//		var response models.Response[[]models.Subscription]
//		if err := client.Subscriptions.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Subscriptions.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/subscription/
func (s *SubscriptionClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/subscription", queries...)
	return s.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve details of a subscription on your Integration
//
// Default response: models.Response[models.Subscription]
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
//		var response models.Response[models.Subscription]
//		if err := client.Subscriptions.FetchOne(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (s *SubscriptionClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return s.APICall(ctx, http.MethodGet, fmt.Sprintf("/subscription/%s", idOrCode), nil, response)
}

// Enable lets you enable a subscription on your Integration
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
//		if err := client.Subscriptions.Enable(context.TODO(),"SUB_vsyqdmlzble3uii", "d7gofp6yppn3qz7", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (s *SubscriptionClient) Enable(ctx context.Context, code string, token string, response any) error {
	payload := map[string]any{
		"code":  code,
		"token": token,
	}

	return s.APICall(ctx, http.MethodPost, "/subscription/enable", payload, response)
}

// Disable lets you disable a subscription on your Integration
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
//		if err := client.Subscriptions.Disable(context.TODO(),"SUB_vsyqdmlzble3uii", "d7gofp6yppn3qz7", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (s *SubscriptionClient) Disable(ctx context.Context, code string, token string, response any) error {
	payload := map[string]any{
		"code":  code,
		"token": token,
	}
	return s.APICall(ctx, http.MethodPost, "/subscription/disable", payload, response)
}

// GenerateLink lets you generate a link for updating the card on a subscription
//
// Default response: models.Response[models.SubscriptionLink]
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
//		var response models.Response[models.SubscriptionLink]
//		if err := client.Subscriptions.GenerateLink(context.TODO(),"SUB_vsyqdmlzble3uii", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (s *SubscriptionClient) GenerateLink(ctx context.Context, code string, response any) error {
	return s.APICall(ctx, http.MethodGet, fmt.Sprintf("/subscription/%s/manage/link/", code), nil, response)
}

// SendLink lets you email a customer a link for updating the card on their subscription
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
//		if err := client.Subscriptions.SendLink(context.TODO(),"SUB_vsyqdmlzble3uii", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (s *SubscriptionClient) SendLink(ctx context.Context, code string, response any) error {
	return s.APICall(ctx, http.MethodPost, fmt.Sprintf("/subscription/%s/manage/email/", code), nil, response)
}
