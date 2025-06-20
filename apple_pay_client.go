package paystack

import (
	"context"
	"net/http"
)

// ApplePayClient interacts with endpoints related to paystack Apple Pay resource that
// lets you register your application's top-level domain or subdomain.
type ApplePayClient struct {
	*restClient
}

// NewApplePayClient creates a ApplePayClient
func NewApplePayClient(options ...ClientOptions) *ApplePayClient {
	client := NewClient(options...)

	return client.ApplePay
}

// Register lets you register a top-level domain or subdomain for your Apple Pay Integration.
//
// Default: response models.Response[struct{}]
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
//		if err := client.ApplePay.Register(context.TODO(),"<domainName>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (a *ApplePayClient) Register(ctx context.Context, domainName string, response any) error {
	payload := map[string]string{"domainName": domainName}
	return a.APICall(ctx, http.MethodPost, "/apple-pay/domain", payload, response)
}

// All retrieves all registered Apple Pay domains on your integration.
//
// It returns an empty array if no domains have been added.
//
// Default response: models.Response[models.ApplePayDomains]
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
//		var response models.Response[models.ApplePayDomains]
//		if err := client.ApplePay.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.ApplePay.All(context.TODO(), &response, p.WithQuery("use_cursor", "true"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/apple-pay/
func (a *ApplePayClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/apple-pay/domain", queries...)
	return a.APICall(ctx, http.MethodGet, url, nil, response)
}

// Unregister lets you unregister a top-level domain or subdomain previously used for your Apple Pay Integration.
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
//		if err := client.ApplePay.Unregister(context.TODO(),"<domainName>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (a *ApplePayClient) Unregister(ctx context.Context, domainName string, response any) error {
	payload := map[string]string{"domainName": domainName}
	return a.APICall(ctx, http.MethodDelete, "/apple-pay/domain", payload, response)
}
