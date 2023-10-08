package paystack

import (
	"fmt"
	"net/http"
)

type RefundClient struct {
	*baseAPIClient
}

func (r *RefundClient) Create(transaction string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["transaction"] = transaction

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return r.APICall(http.MethodPost, "/refund", payload)
}

func (r *RefundClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/refund", queries...)
	return r.APICall(http.MethodGet, url, nil)
}

func (r *RefundClient) FetchOne(reference string) (*Response, error) {
	return r.APICall(http.MethodGet, fmt.Sprintf("/refund/%s", reference), nil)
}
