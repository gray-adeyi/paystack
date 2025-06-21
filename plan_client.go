package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// PlanClient interacts with endpoints related to paystack plan resource that lets you
// create and manage installment payment options on your Integration.
type PlanClient struct {
	*restClient
}

// NewPlanClient creates a PlanClient
func NewPlanClient(options ...ClientOptions) *PlanClient {
	client := NewClient(options...)

	return client.Plans
}

// Create lets you create a plan on your Integration
//
// Default response: models.Response[models.Plan]
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
//		if err := client.Plans.Create(context.TODO(),"Monthly retainer", 500000, enum.IntervalMonthly, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Plans.Create(context.TODO(),"Monthly retainer", 500000, enum.IntervalMonthly, &response, p.WithOptionalParameter("description","a test description"))
//	}
// For supported optional parameters, see:
// https://paystack.com/docs/api/plan/
func (p *PlanClient) Create(ctx context.Context, name string, amount int, interval enum.Interval, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":     name,
		"amount":   amount,
		"interval": interval,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPost, "/plan", payload, response)
}

// All lets you retrieve Plans available on your Integration
//
// Default response: models.Response[[]models.Plan]
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
//		var response models.Response[[]models.Plan]
//		if err := client.Plans.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Plans.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/plan/
func (p *PlanClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/plan", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve details of a plan on your Integration
//
// Default response: models.Response[models.Plan]
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
//		var response models.Response[models.Plan]
//		if err := client.Plans.FetchOne(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *PlanClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return p.APICall(ctx, http.MethodGet, fmt.Sprintf("/plan/%s", idOrCode), nil, response)
}

// Update lets you update a plan details on your Integration
//
// Default response: models.Response[models.Plan]
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
//		var response models.Response[models.Plan]
//		if err := client.Plans.Update(context.TODO(),"<idOrCode>","Monthly retainer", 500000, enum.IntervalMonthly, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Plans.Update(context.TODO(),"<idOrCode>","Monthly retainer", 500000, enum.IntervalMonthly, &response, p.WithOptionalParameter("description","test description"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/plan/
func (p *PlanClient) Update(ctx context.Context, idOrCode string, name string, amount int, interval enum.Interval, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name":     name,
		"amount":   amount,
		"interval": interval,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(ctx, http.MethodPut, fmt.Sprintf("/plan/%s", idOrCode), payload, response)
}
