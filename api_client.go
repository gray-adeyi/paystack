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
//	resp, err := client.DedicatedVirtualAccounts.Create("481193", p.WithOptionalParameter("preferred_bank","wema-bank"))
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
	Terminals                *TerminalClient
	Customers                *CustomerClient
	DedicatedVirtualAccounts *DedicatedVirtualAccountClient
	ApplePay                 *ApplePayClient
	SubAccounts              *SubAccountClient
	Plans                    *PlanClient
	Subscriptions            *SubscriptionClient
	Products                 *ProductClient
	PaymentPages             *PaymentPageClient
	PaymentRequests          *PaymentRequestClient
	Settlements              *SettlementClient
	TransferRecipients       *TransferRecipientClient
	Transfers                *TransferClient
	TransferControl          *TransferControlClient
	BulkCharges              *BulkChargeClient
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

	newClient.Terminals.baseUrl = BaseUrl
	newClient.Terminals.secretKey = newClient.secretKey
	newClient.Terminals.httpClient = httpClient

	newClient.Customers.baseUrl = BaseUrl
	newClient.Customers.secretKey = newClient.secretKey
	newClient.Customers.httpClient = httpClient

	newClient.DedicatedVirtualAccounts.baseUrl = BaseUrl
	newClient.DedicatedVirtualAccounts.secretKey = newClient.secretKey
	newClient.DedicatedVirtualAccounts.httpClient = httpClient

	newClient.ApplePay.baseUrl = BaseUrl
	newClient.ApplePay.secretKey = newClient.secretKey
	newClient.ApplePay.httpClient = httpClient

	newClient.SubAccounts.baseUrl = BaseUrl
	newClient.SubAccounts.secretKey = newClient.secretKey
	newClient.SubAccounts.httpClient = httpClient

	newClient.Plans.baseUrl = BaseUrl
	newClient.Plans.secretKey = newClient.secretKey
	newClient.Plans.httpClient = httpClient

	newClient.Subscriptions.baseUrl = BaseUrl
	newClient.Subscriptions.secretKey = newClient.secretKey
	newClient.Subscriptions.httpClient = httpClient

	newClient.Products.baseUrl = BaseUrl
	newClient.Products.secretKey = newClient.secretKey
	newClient.Products.httpClient = httpClient

	newClient.PaymentPages.baseUrl = BaseUrl
	newClient.PaymentPages.secretKey = newClient.secretKey
	newClient.PaymentPages.httpClient = httpClient

	newClient.PaymentRequests.baseUrl = BaseUrl
	newClient.PaymentRequests.secretKey = newClient.secretKey
	newClient.PaymentRequests.httpClient = httpClient

	newClient.Settlements.baseUrl = BaseUrl
	newClient.Settlements.secretKey = newClient.secretKey
	newClient.Settlements.httpClient = httpClient

	newClient.TransferRecipients.baseUrl = BaseUrl
	newClient.TransferRecipients.secretKey = newClient.secretKey
	newClient.TransferRecipients.httpClient = httpClient

	newClient.Transfers.baseUrl = BaseUrl
	newClient.Transfers.secretKey = newClient.secretKey
	newClient.Transfers.httpClient = httpClient

	newClient.TransferControl.baseUrl = BaseUrl
	newClient.TransferControl.secretKey = newClient.secretKey
	newClient.TransferControl.httpClient = httpClient

	newClient.BulkCharges.baseUrl = BaseUrl
	newClient.BulkCharges.secretKey = newClient.secretKey
	newClient.BulkCharges.httpClient = httpClient

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
