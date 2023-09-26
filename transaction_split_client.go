package paystack

import (
	"fmt"
	"net/http"
)

type TransactionSplitClient struct {
	*baseAPIClient
}

func (t *TransactionSplitClient) Create(name string, transactionSplitType string, currency string, subaccounts interface{}, bearerType string, bearerSubaccount string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["type"] = transactionSplitType
	payload["currency"] = currency
	payload["subaccounts"] = subaccounts
	payload["bearer_type"] = bearerType
	payload["bearer_subaccount"] = bearerSubaccount

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPost, "/split", payload)
}

func (t *TransactionSplitClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/split", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

func (t *TransactionSplitClient) FetchOne(id string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/split/%s", id), nil)
}

func (t *TransactionSplitClient) Update(id string, name string, active bool, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["active"] = active

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}

	return t.APICall(http.MethodPut, fmt.Sprintf("/split/%s", id), payload)
}

func (t *TransactionSplitClient) Add(id string, subAccount string, share int) (*Response, error) {
	payload := make(map[string]interface{})
	payload["subaccount"] = subAccount
	payload["share"] = share

	return t.APICall(http.MethodPost, fmt.Sprintf("/split/%s/add", id), payload)
}

func (t *TransactionSplitClient) Remove(id string, subAccount string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["subaccount"] = subAccount

	return t.APICall(http.MethodPost, fmt.Sprintf("/split/%s/remove", id), payload)
}
