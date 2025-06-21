package paystack

import (
	"context"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// TransferControlClient interacts with endpoints related to paystack transfer control resource that lets
// you manage settings of your Transfers.
type TransferControlClient struct {
	*restClient
}

// NewTransferControlClient creates a TransferControlClient
func NewTransferControlClient(options ...ClientOptions) *TransferControlClient {
	client := NewClient(options...)
	return client.TransferControl
}

// Balance lets you retrieve the available balance on your Integration
//
// Default response: models.Response[[]models.IntegrationBalance]
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
//		var response models.Response[[]models.IntegrationBalance]
//		if err := client.TransferControl.Balance(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferControlClient) Balance(ctx context.Context, response any) error {
	return t.APICall(ctx, http.MethodGet, "/balance", nil, response)
}

// BalanceLedger lets you retrieve all pay-ins and pay-outs that occurred on your Integration
//
// Default response: models.Response[[]models.BalanceLedgerItem]
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
//		var response models.Response[[]models.BalanceLedgerItem]
//		if err := client.TransferControl.BalanceLedger(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferControlClient) BalanceLedger(ctx context.Context, response any) error {
	return t.APICall(ctx, http.MethodGet, "/balance/ledger", nil, response)
}

// ResendOtp lets you generate a new OTP and sends to customer in the event they are having trouble receiving one.
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
//		"github.com/gray-adeyi/paystack/enum"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[struct{}]
//		if err := client.TransferControl.ResendOtp(context.TODO(),"TRF_vsyqdmlzble3uii",enum.ReasonResendOtp, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferControlClient) ResendOtp(ctx context.Context, transferCode string, reason enum.Reason, response any) error {
	payload := map[string]any{
		"transfer_code": transferCode,
		"reason":        reason,
	}

	return t.APICall(ctx, http.MethodPost, "/transfer/resend_otp", payload, response)
}

// DisableOtp lets you complete Transfers without use of OTPs.
// You will get an OTP to complete the request.
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
//		if err := client.TransferControl.DisableOtp(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferControlClient) DisableOtp(ctx context.Context, response any) error {
	return t.APICall(ctx, http.MethodPost, "/transfer/disable_otp", nil, response)
}

// FinalizeDisableOtp lets you finalize the request to disable OTP on your Transfers.
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
//		if err := client.TransferControl.FinalizeDisableOtp(context.TODO(),"<otp>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferControlClient) FinalizeDisableOtp(ctx context.Context, otp string, response any) error {
	payload := map[string]any{"otp": otp}
	return t.APICall(ctx, http.MethodPost, "/transfer/disable_otp", payload, response)
}

// EnableOtp lets you turn OTP requirement back on.
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
//		if err := client.TransferControl.EnableOtp(context.TODO(),"<otp>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferControlClient) EnableOtp(ctx context.Context, response any) error {
	return t.APICall(ctx, http.MethodPost, "/transfer/enable_otp", nil, response)
}
