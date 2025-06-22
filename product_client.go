package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// ProductClient interacts with endpoints related to paystack product resource that allows you to create and
// manage inventories on your Integration.
type ProductClient struct {
	*restClient
}

// NewProductClient creates a ProductClient
func NewProductClient(options ...ClientOptions) *ProductClient {
	client := NewClient(options...)
	return client.Products
}

// Create lets you create a product on your Integration
//
// Default response: models.Response[models.Product]
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
//		var response models.Response[models.Product]
//		if err := client.Products.Create(context.TODO(),"Puff Puff", "Crispy flour ball with fluffy interior", 5000, enum.CurrencyNgn, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Products.Create(context.TODO(),"Puff Puff", "Crispy flour ball with fluffy interior", 5000, enum.CurrencyNgn, &response, p.WithOptionalPayload("unlimited","true"))
//	}
// For supported optional parameters, see:
// https://paystack.com/docs/api/product/
func (p *ProductClient) Create(ctx context.Context, name string, description string, price int, currency enum.Currency, response any,
	optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"name":        name,
		"description": description,
		"price":       price,
		"currency":    currency,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPost, "/product", payload, response)
}

// All lets you retrieve Products available on your Integration
//
// Default response: models.Response[[]models.Product]
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
//		var response models.Response[[]models.Product]
//		if err := client.Products.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Products.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/product/
func (p *ProductClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/product", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you Get details of a product on your Integration
//
// Default response: models.Response[models.Product]
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
//		var response models.Response[models.Product]
//		if err := client.Plans.FetchOne(context.TODO(),"<id>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *ProductClient) FetchOne(ctx context.Context, id string, response any) error {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/product/%s", id), nil, response)
}

// Update lets you update a product details on your Integration
//
// Default response: models.Response[models.Product]
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
//		var response models.Response[models.Plan]
//		if err := client.Products.Update(context.TODO(),"<id>", "Product Six", "Product Six Description",500000, enum.CurrencyUsd, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Products.Update(context.TODO(),"<id>", "Product Six", "Product Six Description",500000, enum.CurrencyUsd, &response, p.WithOptionalPayload("unlimited","true"))
//	}
func (p *ProductClient) Update(ctx context.Context, id string, name string, description string, price int, currency enum.Currency, response any, optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"name":        name,
		"description": description,
		"price":       price,
		"currency":    currency,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPut, fmt.Sprintf("/product/%s", id), payload, response)
}
