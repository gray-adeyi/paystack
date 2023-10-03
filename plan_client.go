package paystack

import (
	"fmt"
	"net/http"
)

type PlanClient struct {
	*baseAPIClient
}

func (p *PlanClient) Create(name string, amount int, interval string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["amount"] = amount
	payload["interval"] = interval

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPost, "/plan", payload)
}

func (p *PlanClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/plan", queries...)
	return p.APICall(http.MethodGet, url, nil)
}

func (p *PlanClient) FetchOne(idOrCode string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/plan/%s", idOrCode), nil)
}

func (p *PlanClient) Update(idOrCode string, name string, amount int, interval string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["amount"] = amount
	payload["interval"] = interval

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPut, fmt.Sprintf("/plan/%s", idOrCode), payload)
}
