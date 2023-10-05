package paystack

import (
	"fmt"
	"net/http"
)

type ProductClient struct {
	*baseAPIClient
}

func (p *ProductClient) Create(name string, description string, price int, currency string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["description"] = description
	payload["price"] = price
	payload["currency"] = currency

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPost, "/product", payload)
}

func (p *ProductClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/product", queries...)
	return p.APICall(http.MethodGet, url, nil)
}

func (p *ProductClient) FetchOne(id string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/product/%s", id), nil)
}

func (p *ProductClient) Update(id string, name string, description string, price int, currency string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["description"] = description
	payload["price"] = price
	payload["currency"] = currency

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPut, fmt.Sprintf("/product/%s", id), nil)
}
