package paystack

import (
	"context"
	"fmt"
	"net/http"
)

// TransferClient interacts with endpoints related to paystack transfer resource that lets you
// automate sending money to your Customers.
type TransferClient struct {
	*restClient
}

// NewTransferClient creates a TransferClient
func NewTransferClient(options ...ClientOptions) *TransferClient {
	client := NewClient(options...)
	return client.Transfers
}

// Initiate lets you send money to your Customers.
// Status of a transfer object returned will be pending if OTP is disabled.
// In the event that an OTP is required, status will read otp.
//
// Default response: models.Response[models.Transfer]
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
//		var response models.Response[models.Transfer]
//		if err := client.Transfers.Initiate(context.TODO(),"balance",500000,"RCP_gx2wn530m0i3w3m", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Transfers.Initiate(context.TODO(),"balance",500000,"RCP_gx2wn530m0i3w3m", &response, p.WithOptionalPayload("reason","Discount Refund"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/transfer/
func (t *TransferClient) Initiate(ctx context.Context, source string, amount int, recipient string, response any,
	optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"source":    source,
		"amount":    amount,
		"recipient": recipient,
	}

	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transfer", payload, response)
}

// Finalize lets you finalize an initiated transfer
//
// Default response: models.Response[models.Transfer]
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
//		var response models.Response[models.Transfer]
//		if err := client.Transfers.Finalize(context.TODO(),"TRF_vsyqdmlzble3uii","928783", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferClient) Finalize(ctx context.Context, transferCode string, otp string, response any) error {
	payload := map[string]any{
		"transfer_code": transferCode,
		"otp":           otp,
	}

	return t.APICall(ctx, http.MethodPost, "/transfer/finalize_transfer", payload, response)
}

// BulkInitiate lets you initiate multiple Transfers in a single request.
// You need to disable the Transfers OTP requirement to use this endpoint.
//
// Default response: models.Response[[]models.BulkTransferItem]
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
// 		transfers := []map[string]any{
//			{"amount": 20000,"reference": "588YtfftReF355894J","reason": "Why not?","recipient":"RCP_2tn9clt23s7qr28"},
//			{"amount": 30000,"reference": "YunoTReF35e0r4J","reason": "Because I can","recipient":"RCP_1a25w1h3n0xctjg"},
//			{"amount": 40000,"reason": "Coming right up","recipient": "RCP_aps2aibr69caua7"},
//		}
//		var response models.Response[[]models.BulkTransferItem]
//		if err := client.Transfers.BulkInitiate(context.TODO(),"balance",transfers, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferClient) BulkInitiate(ctx context.Context, source string, transfers, response any) error {
	payload := map[string]any{
		"source":    source,
		"transfers": transfers,
	}

	return t.APICall(ctx, http.MethodPost, "/transfer/bulk", payload, response)
}

// All lets you retrieve all the Transfers made on your Integration.
//
// Default response: models.Response[[]models.Transfer]
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
//		var response models.Response[[]models.Transfer]
//		if err := client.Transfers.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Transfers.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/transfer/
func (t *TransferClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/transfer", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve the details of a transfer on your Integration.
//
// Default response: models.Response[models.Transfer]
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
//		var response models.Response[models.Transfer]
//		if err := client.Transfers.FetchOne(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transfer/%s", idOrCode), nil, response)
}

// Verify lets you verify the status of a transfer on your Integration.
//
// Default response: models.Response[models.Transfer]
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
//		var response models.Response[models.Transfer]
//		if err := client.Transfers.Verify(context.TODO(),"<reference>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferClient) Verify(ctx context.Context, reference string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transfer/verify/%s", reference), nil, response)
}
