package paystack

import (
	"fmt"
	"net/http"
)

type PaymentPageClient struct {
	*baseAPIClient
}

func (p *PaymentPageClient) Create(name string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPost, "/page", payload)
}

func (p *PaymentPageClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/page", queries...)
	return p.APICall(http.MethodGet, url, nil)
}

func (p *PaymentPageClient) FetchOne(idOrSlug string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/page/%s", idOrSlug), nil)
}

func (p *PaymentPageClient) Update(idOrSlug string, name string, description string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["description"] = description

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return p.APICall(http.MethodPut, fmt.Sprintf("/page/%s", idOrSlug), payload)
}

func (p *PaymentPageClient) CheckSlug(slug string) (*Response, error) {
	return p.APICall(http.MethodGet, fmt.Sprintf("/page/check_slug_availability/%s", slug), nil)
}

func (p *PaymentPageClient) AddProducts(id string, products []string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["product"] = products
	return p.APICall(http.MethodGet, fmt.Sprintf("/page/%s/product", id), nil)
}
