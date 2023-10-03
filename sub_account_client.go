package paystack

import (
	"fmt"
	"net/http"
)

type SubAccountClient struct {
	*baseAPIClient
}

func (s *SubAccountClient) Create(businessName string, settlementBank string, accountNumber string, percentageCharge float32, description string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["business_name"] = businessName
	payload["settlement_bank"] = settlementBank
	payload["account_number"] = accountNumber
	payload["percentage_charge"] = percentageCharge
	payload["description"] = description

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(http.MethodPost, "/subaccount", payload)
}

func (s *SubAccountClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/subaccount", queries...)
	return s.APICall(http.MethodGet, url, nil)
}

func (s *SubAccountClient) FetchOne(idOrCode string) (*Response, error) {
	return s.APICall(http.MethodGet, fmt.Sprintf("/subaccount/%s", idOrCode), nil)
}

func (s *SubAccountClient) Update(idOrCode string, businessName string, settlementBank string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["business_name"] = businessName
	payload["settlement_bank"] = settlementBank

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return s.APICall(http.MethodPut, fmt.Sprintf("/subaccount/%s", idOrCode), payload)
}
