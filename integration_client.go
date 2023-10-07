package paystack

import "net/http"

type IntegrationClient struct {
	*baseAPIClient
}

func (i *IntegrationClient) Timeout() (*Response, error) {
	return i.APICall(http.MethodGet, "/integration/payment_session_timeout", nil)
}

func (i *IntegrationClient) UpdateTimeout(timeout int) (*Response, error) {
	payload := map[string]interface{}{
		"timeout": timeout,
	}
	return i.APICall(http.MethodPut, "/integration/payment_session_timeout", payload)
}
