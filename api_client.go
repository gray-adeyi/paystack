package paystack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const Version = "0.1.0"
const BaseUrl = "https://api.paystack.co"

var ErrNoSecretKey = errors.New("Paystack secret key was not provided")

// Response is a struct containing the status code and data retrieved from paystack. Response.Data is a slice of
// byte that is JSON serializable.
type Response struct {
	StatusCode int
	Data       []byte
}

// ClientOptions is a type used to modify attributes of an APIClient. It can be passed into the NewAPIClient
// function while creating an APIClient
type ClientOptions = func(client *APIClient)

// WithSecretKey lets you set the secret key of an APIClient. It should be used when creating an APIClient
// with the NewAPIClient function.
//
// Example
//
//	import p "github.com/gray-adeyi/paystack"
//	client := p.NewAPIClient(p.WithSecretKey("<your-paystack-secret-key>"))
func WithSecretKey(secretKey string) ClientOptions {
	return func(client *APIClient) {
		client.secretKey = secretKey
	}
}

// WithBaseUrl lets you override paystack's base url for an APIClient. It should be used when creating an APIClient
// with the NewAPIClient function.
func WithBaseUrl(baseUrl string) ClientOptions {
	return func(client *APIClient) {
		client.baseUrl = baseUrl
	}
}

// OptionalPayloadParameter is a type for storing optional parameters used by some APIClient methods that needs
// to accept optional parameter.
type OptionalPayloadParameter = func(map[string]interface{}) map[string]interface{}

// WithOptionalParameter lets you add optional parameters when calling some client methods and you need to add
// optional parameters to your payload.
//
// Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	client := p.NewAPIClient(p.WithSecretKey("<your-paystack-secret-key>"))
//	resp, err := client.dedicatedVirtualAccounts.Create("481193", p.WithOptionalParameter("preferred_bank","wema-bank"))
//
// WithOptionalParameter is used to pass the `preferred_bank` optional parameter in the client method call
func WithOptionalParameter(key string, value interface{}) OptionalPayloadParameter {
	return func(m map[string]interface{}) map[string]interface{} {
		m[key] = value
		return m
	}
}

type baseAPIClient struct {
	secretKey  string
	baseUrl    string
	httpClient *http.Client
}

func (a *baseAPIClient) APICall(method string, endPointPath string, payload interface{}) (*Response, error) {
	var body *bytes.Buffer
	if payload != nil {
		payloadInBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(payloadInBytes)
	}
	apiRequest, err := http.NewRequest(method, a.baseUrl+endPointPath, body)
	if err != nil {
		return nil, err
	}
	err = a.setHeaders(apiRequest)
	if err != nil {
		return nil, err
	}
	r, err := a.httpClient.Do(apiRequest)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: r.StatusCode,
		Data:       data,
	}, nil
}

func (a *baseAPIClient) setHeaders(request *http.Request) error {
	if a.secretKey == "" {
		return ErrNoSecretKey
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.secretKey))
	request.Header.Set("User-Agent", fmt.Sprintf("github.com/gray-adeyi/paystack version %s", Version))
	request.Header.Add("Content-Type", "application/json")
	return nil
}

// APIClient is a struct that has other dedicated clients bound to it. This provides a convenience for interacting
// with all of paystack's endpoints in your Go project. It should not be instantiated directly but interacting but
// via the NewAPIClient function. As stated above, it has other dedicated clients bound to it as field, therefore,
// after creating an instance of the APIClient type. You can access the associated functions of each dedicated client
// via its field name.
//
//	Example
//	import p "github.com/gray-adeyi/paystack"
//
//	client := p.NewAPIClient(p.WithSecretKey("<your-paystack-secret-key>"))
//	resp, err := client.Transactions.Verify("<reference>")
type APIClient struct {
	baseAPIClient
	Transactions             *TransactionClient
	TransactionSplits        *TransactionSplitClient
	terminals                *TerminalClient
	customers                *CustomerClient
	dedicatedVirtualAccounts *DedicatedVirtualAccountClient
	applePay                 *ApplePayClient
	subAccounts              *SubAccountClient
	plans                    *PlanClient
	subscriptions            *SubscriptionClient
	products                 *ProductClient
	paymentPages             *PaymentPageClient
	paymentRequests          *PaymentRequestClient
	settlements              *SettlementClient
	transferRecipients       *TransferRecipientClient
	transfers                *TransferClient
	transferControl          *TransferControlClient
	bulkCharges              *BulkChargeClient
	integration              *IntegrationClient
	charges                  *ChargeClient
	disputes                 *DisputeClient
	refunds                  *RefundClient
	verification             *VerificationClient
	miscellaneous            *MiscellaneousClient
}

// NewAPIClient lets you create an APIClient. it can accept zero to many client options
//
//	Example
//	import p "github.com/gray-adeyi/paystack"
//
//	client := p.NewAPIClient(p.WithSecretKey("<your-paystack-secret-key>"))
func NewAPIClient(options ...ClientOptions) *APIClient {
	newClient := &APIClient{
		Transactions: &TransactionClient{
			&baseAPIClient{},
		},
	}
	newClient.baseUrl = BaseUrl

	httpClient := &http.Client{}
	newClient.httpClient = httpClient
	for _, opts := range options {
		opts(newClient)
	}

	newClient.Transactions.baseUrl = BaseUrl
	newClient.Transactions.secretKey = newClient.secretKey
	newClient.Transactions.httpClient = httpClient

	newClient.TransactionSplits.baseUrl = BaseUrl
	newClient.TransactionSplits.secretKey = newClient.secretKey
	newClient.TransactionSplits.httpClient = httpClient

	newClient.terminals.baseUrl = BaseUrl
	newClient.terminals.secretKey = newClient.secretKey
	newClient.terminals.httpClient = httpClient

	newClient.customers.baseUrl = BaseUrl
	newClient.customers.secretKey = newClient.secretKey
	newClient.customers.httpClient = httpClient

	newClient.dedicatedVirtualAccounts.baseUrl = BaseUrl
	newClient.dedicatedVirtualAccounts.secretKey = newClient.secretKey
	newClient.dedicatedVirtualAccounts.httpClient = httpClient

	newClient.applePay.baseUrl = BaseUrl
	newClient.applePay.secretKey = newClient.secretKey
	newClient.applePay.httpClient = httpClient

	newClient.subAccounts.baseUrl = BaseUrl
	newClient.subAccounts.secretKey = newClient.secretKey
	newClient.subAccounts.httpClient = httpClient

	newClient.plans.baseUrl = BaseUrl
	newClient.plans.secretKey = newClient.secretKey
	newClient.plans.httpClient = httpClient

	newClient.subscriptions.baseUrl = BaseUrl
	newClient.subscriptions.secretKey = newClient.secretKey
	newClient.subscriptions.httpClient = httpClient

	newClient.products.baseUrl = BaseUrl
	newClient.products.secretKey = newClient.secretKey
	newClient.products.httpClient = httpClient

	newClient.paymentPages.baseUrl = BaseUrl
	newClient.paymentPages.secretKey = newClient.secretKey
	newClient.paymentPages.httpClient = httpClient

	newClient.paymentRequests.baseUrl = BaseUrl
	newClient.paymentRequests.secretKey = newClient.secretKey
	newClient.paymentRequests.httpClient = httpClient

	newClient.settlements.baseUrl = BaseUrl
	newClient.settlements.secretKey = newClient.secretKey
	newClient.settlements.httpClient = httpClient

	newClient.transferRecipients.baseUrl = BaseUrl
	newClient.transferRecipients.secretKey = newClient.secretKey
	newClient.transferRecipients.httpClient = httpClient

	newClient.transfers.baseUrl = BaseUrl
	newClient.transfers.secretKey = newClient.secretKey
	newClient.transfers.httpClient = httpClient

	newClient.transferControl.baseUrl = BaseUrl
	newClient.transferControl.secretKey = newClient.secretKey
	newClient.transferControl.httpClient = httpClient

	newClient.bulkCharges.baseUrl = BaseUrl
	newClient.bulkCharges.secretKey = newClient.secretKey
	newClient.bulkCharges.httpClient = httpClient

	newClient.integration.baseUrl = BaseUrl
	newClient.integration.secretKey = newClient.secretKey
	newClient.integration.httpClient = httpClient

	newClient.charges.baseUrl = BaseUrl
	newClient.charges.secretKey = newClient.secretKey
	newClient.charges.httpClient = httpClient

	newClient.charges.baseUrl = BaseUrl
	newClient.charges.secretKey = newClient.secretKey
	newClient.charges.httpClient = httpClient

	newClient.disputes.baseUrl = BaseUrl
	newClient.disputes.secretKey = newClient.secretKey
	newClient.disputes.httpClient = httpClient

	newClient.refunds.baseUrl = BaseUrl
	newClient.refunds.secretKey = newClient.secretKey
	newClient.refunds.httpClient = httpClient

	newClient.verification.baseUrl = BaseUrl
	newClient.verification.secretKey = newClient.secretKey
	newClient.verification.httpClient = httpClient

	newClient.miscellaneous.baseUrl = BaseUrl
	newClient.miscellaneous.secretKey = newClient.secretKey
	newClient.miscellaneous.httpClient = httpClient

	return newClient
}

// Query helps represent key value pairs used in url query parametes
type Query struct {
	Key   string
	Value string
}

// WithQuery lets you create a Query from key value pairs
func WithQuery(key string, value string) Query {
	return Query{
		Key:   key,
		Value: value,
	}
}

// AddQueryParamsToUrl lets you add query parameters to a url
func AddQueryParamsToUrl(url string, queries ...Query) string {
	for _, query := range queries {
		if strings.Contains(url, "?") {
			url += fmt.Sprintf("&%s=%s", query.Key, query.Value)
		} else {
			url += fmt.Sprintf("?%s=%s", query.Key, query.Value)
		}
	}
	return url
}
