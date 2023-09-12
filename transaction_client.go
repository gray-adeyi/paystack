package paystack

import "net/http"

type TransactionClient struct {
	*baseAPIClient
}

func (t *TransactionClient) Initialize(amount int, email string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["amount"] = amount
	payload["email"] = email

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPost, "/transaction/initialize", payload)
}
