package paystack

import (
	"fmt"
	"net/http"
)

type DedicatedVirtualAccountClient struct {
	*baseAPIClient
}

func (d *DedicatedVirtualAccountClient) Create(customerIdOrCode string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = customerIdOrCode

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(http.MethodPost, "/dedicated_account", payload)
}

func (d *DedicatedVirtualAccountClient) Assign(email string, firstName string, lastName string, phone string, preferredBank string, country string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["email"] = email
	payload["first_name"] = firstName
	payload["last_name"] = lastName
	payload["phone"] = phone
	payload["preferred_bank"] = preferredBank
	payload["country"] = country

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(http.MethodPost, "/dedicated_account/assign", payload)
}

func (d *DedicatedVirtualAccountClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/dedicated_account", queries...)
	return d.APICall(http.MethodGet, url, nil)
}

func (d *DedicatedVirtualAccountClient) FetchOne(emailOrCode string) (*Response, error) {
	return d.APICall(http.MethodGet, fmt.Sprintf("/dedicated_account/%s", emailOrCode), nil)
}

func (d *DedicatedVirtualAccountClient) Requery(queries ...Query) (*Response, error) {
	return d.All(queries...)
}

func (d *DedicatedVirtualAccountClient) Deactivate(id string) (*Response, error) {
	return d.APICall(http.MethodDelete, fmt.Sprintf("/dedicated_account/%s", id), nil)
}

func (d *DedicatedVirtualAccountClient) Split(customerIdOrCode string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["customer"] = customerIdOrCode

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return d.APICall(http.MethodPost, "/dedicated_account/split", payload)
}

func (d *DedicatedVirtualAccountClient) RemoveSplit(accountNumber string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["account_number"] = accountNumber
	return d.APICall(http.MethodDelete, "/dedicated_account/split", payload)
}

func (d *DedicatedVirtualAccountClient) BankProviders() (*Response, error) {
	return d.APICall(http.MethodPost, "/dedicated_account/available_providers", nil)
}
