package paystack

import (
	"fmt"
	"net/http"
)

type SubscriptionClient struct {
	*baseAPIClient
}

func (s *SubscriptionClient) Create(customer string, plan string, authorization string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = customer
	payload["plan"] = plan
	payload["authorization"] = authorization

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(http.MethodPost, "/subscription", payload)
}

func (s *SubscriptionClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/subscription", queries...)
	return s.APICall(http.MethodGet, url, nil)
}

func (s *SubscriptionClient) FetchOne(idOrCode string) (*Response, error) {
	return s.APICall(http.MethodGet, fmt.Sprintf("/subscription/%s", idOrCode), nil)
}

func (s *SubscriptionClient) Enable(code string, token string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["code"] = code
	payload["token"] = token
	return s.APICall(http.MethodPost, "/subscription/enable", payload)
}

func (s *SubscriptionClient) Disable(code string, token string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["code"] = code
	payload["token"] = token
	return s.APICall(http.MethodPost, "/subscription/disable", payload)
}

func (s *SubscriptionClient) GenerateLink(code string) (*Response, error) {
	return s.APICall(http.MethodGet, fmt.Sprintf("/subscription/%s/manage/link/", code), nil)
}

func (s *SubscriptionClient) SendLink(code string) (*Response, error) {
	return s.APICall(http.MethodGet, fmt.Sprintf("/subscription/%s/manage/email/", code), nil)
}
