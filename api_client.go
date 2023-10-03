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

type Response struct {
	StatusCode int
	Data       []byte
}

type ClientOptions = func(client *APIClient)

func WithSecretKey(secretKey string) ClientOptions {
	return func(client *APIClient) {
		client.secretKey = secretKey
	}
}

func WithBaseUrl(baseUrl string) ClientOptions {
	return func(client *APIClient) {
		client.baseUrl = baseUrl
	}
}

type OptionalPayloadParameter = func(map[string]interface{}) map[string]interface{}

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

type APIClient struct {
	baseAPIClient
	transactions             *TransactionClient
	transactionSplits        *TransactionSplitClient
	terminals                *TerminalClient
	customers                *CustomerClient
	dedicatedVirtualAccounts *DedicatedVirtualAccountClient
	applePay                 *ApplePayClient
	subAccounts              *SubAccountClient
	plans                    *PlanClient
	subscriptions            *SubscriptionClient
	products                 *ProductClient
}

func NewAPIClient(options ...ClientOptions) *APIClient {
	newClient := &APIClient{
		transactions: &TransactionClient{
			&baseAPIClient{},
		},
	}
	newClient.baseUrl = BaseUrl

	httpClient := &http.Client{}
	newClient.httpClient = httpClient
	for _, opts := range options {
		opts(newClient)
	}

	newClient.transactions.baseUrl = BaseUrl
	newClient.transactions.secretKey = newClient.secretKey
	newClient.transactions.httpClient = httpClient

	newClient.transactionSplits.baseUrl = BaseUrl
	newClient.transactionSplits.secretKey = newClient.secretKey
	newClient.transactionSplits.httpClient = httpClient

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

	return newClient
}

type Query struct {
	Key   string
	Value string
}

func WithQuery(key string, value string) Query {
	return Query{
		Key:   key,
		Value: value,
	}
}

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
