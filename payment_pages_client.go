package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// PaymentPageClient interacts with endpoints
// related to paystack payment page resource
// that lets you provide a quick and secure way to collect payment for Products.
type PaymentPageClient struct {
	*restClient
}

// NewPaymentPageClient creates a PaymentPageClient
func NewPaymentPageClient(options ...ClientOptions) *PaymentPageClient {
	client := NewClient(options...)
	return client.PaymentPages
}

// Create lets you create a payment page on your Integration
//
// Default response: models.Response[models.PaymentPage]
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
//		var response models.Response[models.PaymentPage]
//		if err := client.PaymentPages.Create(context.TODO(),"Buttercup Brunch", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.PaymentPages.Create(context.TODO(),"Buttercup Brunch", &response, p.WithOptionalParameter("amount",500000))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/page/
func (p *PaymentPageClient) Create(ctx context.Context, name string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name": name,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPost, "/page", payload, response)
}

// All lets you retrieve payment pages available on your Integration
//
// Default response: models.Response[[]models.PaymentPage]
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
//		var response models.Response[[]models.PaymentPage]
//		if err := client.PaymentPages.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.PaymentPages.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/page/
func (p *PaymentPageClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/page", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve details of a payment page on your Integration
//
// Default response: models.Response[models.PaymentPage]
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
//		var response models.Response[models.PaymentPage]
//		if err := client.PaymentPages.FetchOne(context.TODO(),"<idOrSlug>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentPageClient) FetchOne(ctx context.Context, idOrSlug string, response any) error {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/page/%s", idOrSlug), nil, response)
}

// Update lets you update a payment page details on your Integration
//
// Default response: models.Response[models.PaymentPage]
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
//		var response models.Response[models.PaymentPage]
//		if err := client.PaymentPages.Update(context.TODO(),"<idOrSlug>", "Buttercup Brunch", "description", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.PaymentPages.Update(context.TODO(), "Buttercup Brunch", "description", &response, p.WithOptionalParameter("amount",500000))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/page/
func (p *PaymentPageClient) Update(ctx context.Context, idOrSlug string, name string, description string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":        name,
		"description": description,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPut, fmt.Sprintf("/page/%s", idOrSlug), payload, response)
}

// CheckSlug lets you check the availability of a slug for a payment page
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
//		if err := client.PaymentPages.CheckSlug(context.TODO(),"<slug>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentPageClient) CheckSlug(ctx context.Context, slug string, response any) error {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/page/check_slug_availability/%s", slug), nil, response)
}

// AddProducts lets you add Products to a payment page
//
// Default response: models.Response[models.PaymentPage]
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
//		var response models.Response[models.PaymentPage]
//		if err := client.PaymentPages.AddProducts(context.TODO(),"<id>", []string{4"73", "292"}, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PaymentPageClient) AddProducts(ctx context.Context, id string, products []string, response any) error {
	payload := map[string][]string{
		"product": products,
	}
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/page/%s/product", id), payload, response)
}
