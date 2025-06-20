package paystack

import (
	"context"
	"net/http"
)

// IntegrationClient interacts with endpoints related to paystack Integration resource
// that lets you manage some settings on your Integration.
type IntegrationClient struct {
	*restClient
}

// NewIntegrationClient creates an IntegrationClient
func NewIntegrationClient(options ...ClientOptions) *IntegrationClient {
	client := NewClient(options...)
	return client.Integration
}

// Timeout lets you retrieve the payment session timeout on your Integration
//
// Default response: models.Response[models.IntegrationTimeout]
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
//		var response models.Response[models.IntegrationTimeout]
//		if err := client.Integration.Timeout(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (i *IntegrationClient) Timeout(ctx context.Context, response any) error {
	return i.APICall(ctx, http.MethodGet, "/integration/payment_session_timeout", nil, response)
}

// UpdateTimeout lets you update the payment session timeout on your Integration
//
// Default response: models.Response[models.IntegrationTimeout]
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
//		var response models.Response[models.IntegrationTimeout]
//		if err := client.Integration.UpdateTimeout(context.TODO(),5, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (i *IntegrationClient) UpdateTimeout(ctx context.Context, timeout int, response any) error {
	payload := map[string]any{
		"timeout": timeout,
	}
	return i.APICall(ctx, http.MethodPut, "/integration/payment_session_timeout", payload, response)
}
