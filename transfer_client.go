package paystack

import (
	"fmt"
	"net/http"
)

type TransferClient struct {
	*baseAPIClient
}

func (t *TransferClient) Initiate(source string, amount int, recipient string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["source"] = source
	payload["amount"] = amount
	payload["recipient"] = recipient

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPost, "/transfer", payload)
}

func (t *TransferClient) Finalize(transferCode string, otp string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["transfer_code"] = transferCode
	payload["otp"] = otp
	return t.APICall(http.MethodPost, "/transfer/finalize_transfer", payload)
}

func (t *TransferClient) BulkInitiate(source string, transfers interface{}) (*Response, error) {
	payload := make(map[string]interface{})
	payload["source"] = source
	payload["transfers"] = transfers

	return t.APICall(http.MethodPost, "/transfer/bulk", payload)
}

func (t *TransferClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transfer", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

func (t *TransferClient) FetchOne(idOrCode string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transfer/%s", idOrCode), nil)
}

func (t *TransferClient) Verify(reference string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transfer/verify/%s", reference), nil)
}
