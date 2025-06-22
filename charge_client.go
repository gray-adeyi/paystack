package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// ChargeClient interacts with endpoints related to paystack charge resource that
// lets you configure a payment channel of your choice when initiating a payment.
type ChargeClient struct {
	*restClient
}

// NewChargeClient creates a ChargeClient
func NewChargeClient(options ...ClientOptions) *ChargeClient {
	client := NewClient(options...)
	return client.Charges
}

// Create lets you initiate a payment by integrating the payment channel of your choice.
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Charges.Create(context.TODO(),"johndoe@example.com",100000, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Charges.Create(context.TODO(),"johndoe@example.com",100000, &response,p.WithOptionalPayload("authorization_code","AUTH_xxx"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/charge/
func (c *ChargeClient) Create(ctx context.Context, email string, amount string, response any, optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"email":  email,
		"amount": amount,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(ctx, http.MethodPost, "/charge", payload, response)
}

// SubmitPin lets you submit pin to continue a charge
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Charges.SubmitPin(context.TODO(),"1234", "5bwib5v6anhe9xa", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (c *ChargeClient) SubmitPin(ctx context.Context, pin string, reference string, response any) error {
	payload := map[string]any{
		"pin":       pin,
		"reference": reference,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_pin", payload, response)
}

// SubmitOtp lets you submit otp to complete a charge
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Charges.SubmitOpt(context.TODO(),"123456", "5bwib5v6anhe9xa", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (c *ChargeClient) SubmitOtp(ctx context.Context, otp string, reference string, response any) error {
	payload := map[string]any{
		"otp":       otp,
		"reference": reference,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_otp", payload, response)
}

// SubmitPhone lets you submit phone number when requested
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Charges.SubmitPhone(context.TODO(),"08012345678", "5bwib5v6anhe9xa", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (c *ChargeClient) SubmitPhone(ctx context.Context, phone string, reference string, response any) error {
	payload := map[string]any{
		"phone":     phone,
		"reference": reference,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_phone", payload, response)
}

// SubmitBirthday lets you submit a birthday when requested
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Charges.SubmitBirthday(context.TODO(),"2016-09-21", "5bwib5v6anhe9xa", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (c *ChargeClient) SubmitBirthday(ctx context.Context, birthday string, reference string, response any) error {
	payload := map[string]any{
		"birthday":  birthday,
		"reference": reference,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_birthday", payload, response)
}

// SubmitAddress lets you submit address to continue a charge
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Charges.SubmitAddress(context.TODO(),"140 N 2ND ST", "7c7rpkqpc0tijs8", "Stroudsburg", "PA", "18360", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (c *ChargeClient) SubmitAddress(ctx context.Context, address string, reference string, city string,
	state string, zipCode string, response any) error {
	payload := map[string]any{
		"address":   address,
		"reference": reference,
		"city":      city,
		"state":     state,
		"zipcode":   zipCode,
	}

	return c.APICall(ctx, http.MethodPost, "/charge/submit_address", payload, response)
}

// PendingCharge lets you check the status of a charge. When you get pending as a charge status or if there
// was an exception when calling any of the ChargeClient methods, wait 10 seconds or more, then make a check
// to see if its status has changed. Don't call too early as you may get a lot more pending than you should.
//
// Default response: models.Response[models.Transaction]
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
//		var response models.Response[models.Transaction]
//		if err := client.Charges.PendingCharge(context.TODO(), "5bwib5v6anhe9xa", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (c *ChargeClient) PendingCharge(ctx context.Context, reference string, response any) error {
	return c.APICall(ctx, http.MethodGet, fmt.Sprintf("/charge/%s", reference), nil, response)
}
