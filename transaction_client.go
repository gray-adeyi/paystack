package paystack

import (
	"fmt"
	"net/http"
)

// TransactionClient interacts with endpoints related to paystack's Transaction resource
// that allows you to create and manage payments on your integration.
type TransactionClient struct {
	*baseAPIClient
}

// NewTransactionClient creates a TransactionClient
func NewTransactionClient(options ...ClientOptions) *TransactionClient {
	client := NewAPIClient(options...)
	return client.transactions
}

// Initialize lets you initialize a transaction from your backend
func (t *TransactionClient) Initialize(amount int, email string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := make(map[string]interface{})
	payload["amount"] = amount
	payload["email"] = email

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(http.MethodPost, "/transaction/initialize", payload)
}

// Verify lets you confirm the status of a transaction
func (t *TransactionClient) Verify(reference string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transaction/verify/%s", reference), nil)
}

// All lets you list transactions carried out on your integration
func (t *TransactionClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

// FetchOne lets you get the details of a transaction carried out on your integration
func (t *TransactionClient) FetchOne(id string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transaction/%s", id), nil)
}

// ChargeAuthorization lets you charge authorizations that are marked as reusable
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

// Timeline lets you view the timeline of a transaction
func (t *TransactionClient) Timeline(idOrReference string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/transaction/timeline/%s", idOrReference), nil)
}

// Total lets you retrieve the total amount received on your account
func (t *TransactionClient) Total(idOrReference string, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction/totals", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

// Export lets you export a list of transactions carried out on your integration
func (t *TransactionClient) Export(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/transaction/export", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

// PartialDebit lets you retrieve part of a payment from a customer
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
