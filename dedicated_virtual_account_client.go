package paystack

import "net/http"

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
