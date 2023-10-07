package paystack

import (
	"fmt"
	"net/http"
)

type ChargeClient struct {
	*baseAPIClient
}

func (c *ChargeClient) Create(email string, amount string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["email"] = email
	payload["amount"] = amount

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return c.APICall(http.MethodPost, "/charge", payload)
}

func (c *ChargeClient) SubmitPin(pin string, reference string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["pin"] = pin
	payload["reference"] = reference

	return c.APICall(http.MethodPost, "/charge/submit_pin", payload)
}

func (c *ChargeClient) SubmitPhone(phone string, reference string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["phone"] = phone
	payload["reference"] = reference

	return c.APICall(http.MethodPost, "/charge/submit_phone", payload)
}

func (c *ChargeClient) SubmitBirthday(birthday string, reference string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["birthday"] = birthday
	payload["reference"] = reference

	return c.APICall(http.MethodPost, "/charge/submit_birthday", payload)
}

func (c *ChargeClient) SubmitAddress(address string, reference string, city string, state string, zipCode string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["address"] = address
	payload["reference"] = reference
	payload["city"] = city
	payload["state"] = state
	payload["zipcode"] = zipCode

	return c.APICall(http.MethodPost, "/charge/submit_address", payload)
}

func (c *ChargeClient) PendingCharge(reference string) (*Response, error) {
	return c.APICall(http.MethodGet, fmt.Sprintf("/charge/%s", reference), nil)
}
