package paystack

import "net/http"

// TransferControlClient interacts with endpoints related to paystack transfer control resource that lets
// you manage settings of your transfers.
type TransferControlClient struct {
	*baseAPIClient
}

// NewTransferControlClient creates a TransferControlClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	tcClient := p.NewTransferControlClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTransferControlClient(options ...ClientOptions) *TransferControlClient {
	client := NewAPIClient(options...)
	return client.transferControl
}

// Balance lets you retrieve the available balance on your integration
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
//	// paystackClient.transferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transferControl.Balance()
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
func (t *TransferControlClient) Balance() (*Response, error) {
	return t.APICall(http.MethodGet, "/balance", nil)
}

// BalanceLedger lets you retrieve all pay-ins and pay-outs that occurred on your integration
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
//	// paystackClient.transferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transferControl.BalanceLedger()
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
func (t *TransferControlClient) BalanceLedger() (*Response, error) {
	return t.APICall(http.MethodGet, "/balance/ledger", nil)
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
//	// paystackClient.transferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transferControl.ResendOTP("TRF_vsyqdmlzble3uii","resend_otp")
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
func (t *TransferControlClient) ResendOTP(transferCode string, reason string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["transfer_code"] = transferCode
	payload["reason"] = reason
	return t.APICall(http.MethodPost, "/transfer/resend_otp", payload)
}

// DisableOTP lets you complete transfers without use of OTPs.
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
//	// paystackClient.transferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transferControl.DisableOTP()
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
func (t *TransferControlClient) DisableOTP() (*Response, error) {
	return t.APICall(http.MethodPost, "/transfer/disable_otp", nil)
}

// FinalizeDisableOTP lets you finalize the request to disable OTP on your transfers.
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
//	// paystackClient.transferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transferControl.FinalizeDisableOTP("<otp>")
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
func (t *TransferControlClient) FinalizeDisableOTP(otp string) (*Response, error) {
	payload := map[string]interface{}{"otp": otp}
	return t.APICall(http.MethodPost, "/transfer/disable_otp", payload)
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
//	// paystackClient.transferControl field is a `TransferControlClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.transferControl.EnableOTP()
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
func (t *TransferControlClient) EnableOTP() (*Response, error) {
	return t.APICall(http.MethodPost, "/transfer/enable_otp", nil)
}
