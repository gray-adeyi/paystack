package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gray-adeyi/paystack/models"
)

const Version = "0.1.0"
const BaseUrl = "https://api.paystack.co"

var ErrNoSecretKey = errors.New("Paystack secret key was not provided")



// ClientOptions is a type used to modify attributes of an APIClient. It can be passed into the NewAPIClient
// function while creating an APIClient
type ClientOptions = func(client *restClient)

type restClient struct {
	secretKey  string
	baseUrl    string
	httpClient *http.Client
}

// WithSecretKey lets you set the secret key of an APIClient. It should be used when creating an APIClient
// with the NewAPIClient function.
//
// Example
//
//	import p "github.com/gray-adeyi/paystack"
//	client := p.NewAPIClient(p.WithSecretKey("<your-paystack-secret-key>"))
func WithSecretKey(secretKey string) ClientOptions {
	return func(client *restClient) {
		client.secretKey = secretKey
	}
}

// WithBaseUrl lets you override paystack's base url for an APIClient. It should be used when creating an APIClient
// with the NewAPIClient function.
func WithBaseUrl(baseUrl string) ClientOptions {
	return func(client *restClient) {
		client.baseUrl = baseUrl
	}
}

func (a *restClient) APICall(ctx context.Context, method string, endPointPath string, payload interface{}) (*models.Response, error) {
	var body *bytes.Buffer
	var apiRequest *http.Request
	var err error

	if payload != nil {
		payloadInBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(payloadInBytes)
	}

	if payload != nil {
		apiRequest, err = http.NewRequestWithContext(ctx, method, a.baseUrl+endPointPath, body)
	} else {
		apiRequest, err = http.NewRequestWithContext(ctx, method, a.baseUrl+endPointPath, nil)
	}

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
	return &models.Response{
		StatusCode: r.StatusCode,
		Data:       data,
	}, nil
}

func (a *restClient) setHeaders(request *http.Request) error {
	if a.secretKey == "" {
		return ErrNoSecretKey
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.secretKey))
	request.Header.Set("User-Agent", fmt.Sprintf("github.com/gray-adeyi/paystack version %s", Version))
	request.Header.Add("Content-Type", "application/json")
	return nil
}