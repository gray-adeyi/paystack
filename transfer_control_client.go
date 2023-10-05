package paystack

import "net/http"

type TransferControlClient struct {
	*baseAPIClient
}

func (t *TransferControlClient) Balance() (*Response, error) {
	return t.APICall(http.MethodGet, "/balance", nil)
}

func (t *TransferControlClient) BalanceLedger() (*Response, error) {
	return t.APICall(http.MethodGet, "/balance/ledger", nil)
}

func (t *TransferControlClient) ResendOTP(transferCode string, reason string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["transfer_code"] = transferCode
	payload["reason"] = reason
	return t.APICall(http.MethodPost, "/transfer/resend_otp", payload)
}

func (t *TransferControlClient) DisableOTP() (*Response, error) {
	return t.APICall(http.MethodPost, "/transfer/disable_otp", nil)
}

func (t *TransferControlClient) FinalizeDisableOTP(otp string) (*Response, error) {
	payload := map[string]interface{}{"otp": otp}
	return t.APICall(http.MethodPost, "/transfer/disable_otp", payload)
}

func (t *TransferControlClient) EnableOTP() (*Response, error) {
	return t.APICall(http.MethodPost, "/transfer/enable_otp", nil)
}
