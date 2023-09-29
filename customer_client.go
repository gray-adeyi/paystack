package paystack

import (
	"fmt"
	"net/http"
)

type CustomerClient struct {
	*baseAPIClient
}

func (c *CustomerClient) Create(email string, firstName string, lastName string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["email"] = email
	payload["first_name"] = firstName
	payload["last_name"] = lastName

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(http.MethodPost, "/customer", payload)
}

func (c *CustomerClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/terminal", queries...)
	return c.APICall(http.MethodGet, url, nil)
}

func (c *CustomerClient) FetchOne(emailOrCode string) (*Response, error) {
	return c.APICall(http.MethodGet, fmt.Sprintf("/customer/%s", emailOrCode), nil)
}

func (c *CustomerClient) Update(code string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(http.MethodPut, fmt.Sprintf("/customer/%s", code), payload)
}

func (c *CustomerClient) Validate(code string, firstName string, lastName string, identificationType string, value string, country string, bvn string, bankCode string, accountNumber string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["first_name"] = firstName
	payload["last_name"] = lastName
	payload["type"] = identificationType
	payload["value"] = value
	payload["country"] = country
	payload["bvn"] = bvn
	payload["bank_code"] = bankCode
	payload["account_number"] = accountNumber

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(http.MethodPost, fmt.Sprintf("/customer/%s/identification", code), payload)
}

func (c *CustomerClient) Flag(emailOrCode string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = emailOrCode

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(http.MethodPost, "/customer/set_risk_action", payload)
}

func (c *CustomerClient) Deactivate(authorizationCode string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["authorization_code"] = authorizationCode

	return c.APICall(http.MethodPost, "/customer/deactivate_authorization", payload)
}
