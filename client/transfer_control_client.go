package client

import (
	"context"
	"net/http"
)

// TransferControlClient interacts with endpoints related to paystack transfer control resource that lets
// you manage settings of your Transfers.
type TransferControlClient struct {
	*restClient
}

// NewTransferControlClient creates a TransferControlClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	tcClient := p.NewTransferControlClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTransferControlClient(options ...ClientOptions) *TransferControlClient {
	client := NewPaystackClient(options...)
	return client.TransferControl
}

// Balance lets you retrieve the available balance on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tcClient := p.NewTransferControlClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer control client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferControl.Balance()
//
//	resp, err := tcClient.Balance()
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TransferControlClient) Balance(ctx context.Context) (*Response, error) {
	return t.APICall(ctx, http.MethodGet, "/balance", nil)
}

// BalanceLedger lets you retrieve all pay-ins and pay-outs that occurred on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tcClient := p.NewTransferControlClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer control client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferControl.BalanceLedger()
//
//	resp, err := tcClient.BalanceLedger()
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TransferControlClient) BalanceLedger(ctx context.Context) (*Response, error) {
	return t.APICall(ctx, http.MethodGet, "/balance/ledger", nil)
}

// ResendOTP lets you generate a new OTP and sends to customer in the event they are having trouble receiving one.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tcClient := p.NewTransferControlClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer control client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferControl.ResendOTP("TRF_vsyqdmlzble3uii","resend_otp")
//
//	resp, err := tcClient.ResendOTP("TRF_vsyqdmlzble3uii","resend_otp")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TransferControlClient) ResendOTP(ctx context.Context, transferCode string, reason string) (*Response, error) {
	payload := map[string]any{
		"transfer_code": transferCode,
		"reason":        reason,
	}

	return t.APICall(ctx, http.MethodPost, "/transfer/resend_otp", payload)
}

// DisableOTP lets you complete Transfers without use of OTPs.
// You will get an OTP to complete the request.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tcClient := p.NewTransferControlClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer control client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferControl.DisableOTP()
//
//	resp, err := tcClient.DisableOTP()
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TransferControlClient) DisableOTP(ctx context.Context) (*Response, error) {
	return t.APICall(ctx, http.MethodPost, "/transfer/disable_otp", nil)
}

// FinalizeDisableOTP lets you finalize the request to disable OTP on your Transfers.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tcClient := p.NewTransferControlClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer control client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferControl.FinalizeDisableOTP("<otp>")
//
//	resp, err := tcClient.FinalizeDisableOTP("<otp>")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TransferControlClient) FinalizeDisableOTP(ctx context.Context, otp string) (*Response, error) {
	payload := map[string]any{"otp": otp}
	return t.APICall(ctx, http.MethodPost, "/transfer/disable_otp", payload)
}

// EnableOTP lets you turn OTP requirement back on.
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	tcClient := p.NewTransferControlClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transfer control client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.TransferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.TransferControl.EnableOTP()
//
//	resp, err := tcClient.EnableOTP()
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TransferControlClient) EnableOTP(ctx context.Context) (*Response, error) {
	return t.APICall(ctx, http.MethodPost, "/transfer/enable_otp", nil)
}
