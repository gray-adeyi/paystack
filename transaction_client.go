package paystack

import (
	"fmt"
	"net/http"
)

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

func (t *TransactionClient) Verify(reference string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transaction/verify/%s", reference), nil)
}

func (t *TransactionClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

func (t *TransactionClient) FetchOne(id string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transaction/%s", id), nil)
}

func (t *TransactionClient) ChargeAuthorization(amount int, email string, authorizationCode string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["amount"] = amount
	payload["email"] = email
	payload["authorization_code"] = authorizationCode

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPost, "/transaction/charge_authorization", payload)
}

func (t *TransactionClient) Timeline(idOrReference string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transaction/timeline/%s", idOrReference), nil)
}

func (t *TransactionClient) Total(idOrReference string, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction/totals", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

func (t *TransactionClient) Export(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction/export", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

func (t *TransactionClient) PartialDebit(authorizationCode string, currency string, amount string, email string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["authorization_code"] = authorizationCode
	payload["currency"] = currency
	payload["amount"] = amount
	payload["email"] = email

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPost, "/transaction/partial_debit", payload)
}
