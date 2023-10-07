package paystack

import (
	"fmt"
	"net/http"
)

type PaymentRequestClient struct {
	*baseAPIClient
}

func (p *PaymentRequestClient) Create(customerIdOrCode string, amount int, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = customerIdOrCode
	payload["amount"] = amount

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPost, "/paymentrequest", payload)
}

func (p *PaymentRequestClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/paymentrequest", queries...)
	return p.APICall(http.MethodGet, url, nil)
}

func (p *PaymentRequestClient) FetchOne(idOrCode string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/paymentrequest/%s", idOrCode), nil)
}

func (p *PaymentRequestClient) Verify(code string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/paymentrequest/verify/%s", code), nil)
}

func (p *PaymentRequestClient) SendNotification(code string) (*Response, error) {
	return p.APICall(http.MethodPost, fmt.Sprintf("/paymentrequest/notify/%s", code), nil)
}

func (p *PaymentRequestClient) Total(code string) (*Response, error) {
	return p.APICall(http.MethodGet, "/paymentrequest/totals", nil)
}

func (p *PaymentRequestClient) Finalize(code string, sendNotification bool) (*Response, error) {
	payload := make(map[string]interface{})
	payload["send_notification"] = sendNotification
	return p.APICall(http.MethodPost, fmt.Sprintf("/paymentrequest/finalize/%s", code), nil)
}

func (p *PaymentRequestClient) Update(idOrCode string, customerIdOrCode string, amount int, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = customerIdOrCode
	payload["amount"] = amount

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPut, fmt.Sprintf("/paymentrequest/%s", idOrCode), nil)
}

func (p *PaymentRequestClient) Archive(idOrCode string) (*Response, error) {
	return p.APICall(http.MethodPut, fmt.Sprintf("/paymentrequest/archive/%s", idOrCode), nil)
}
