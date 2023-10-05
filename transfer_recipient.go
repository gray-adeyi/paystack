package paystack

import (
	"fmt"
	"net/http"
)

type TransferRecipientClient struct {
	*baseAPIClient
}

func (t *TransferRecipientClient) Create(recipientType string, name string, accountNumber string, bankCode string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["type"] = recipientType
	payload["name"] = name
	payload["account_number"] = accountNumber
	payload["bank_code"] = bankCode

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPost, "/transferrecipient", payload)
}

func (t *TransferRecipientClient) BulkCreate(batch interface{}) (*Response, error) {
	payload := make(map[string]interface{})
	payload["batch"] = batch

	return t.APICall(http.MethodPost, "/transferrecipient/bulk", payload)
}

func (t *TransferRecipientClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transferrecipient", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

func (t *TransferRecipientClient) FetchOne(idOrCode string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transferrecipient/%s", idOrCode), nil)
}

func (t *TransferRecipientClient) Update(idOrCode string, name string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPut, fmt.Sprintf("/transferrecipient/%s", idOrCode), nil)
}

func (t *TransferRecipientClient) Delete(idOrCode string) (*Response, error) {
	return t.APICall(http.MethodDelete, fmt.Sprintf("/transferrecipient/%s", idOrCode), nil)
}
