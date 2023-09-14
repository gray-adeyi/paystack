package paystack

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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
	transactions *TransactionClient
}

func NewAPIClient(options ...ClientOptions) *APIClient {
	newClient := &APIClient{
		transactions: &TransactionClient{
			&baseAPIClient{},
		},
	}
	newClient.baseUrl = BaseUrl
	newClient.httpClient = &http.Client{}
	for _, opts := range options {
		opts(newClient)
	}

	newClient.transactions.baseUrl = BaseUrl
	newClient.transactions.secretKey = newClient.secretKey
	newClient.transactions.httpClient = &http.Client{}
	return newClient
}
