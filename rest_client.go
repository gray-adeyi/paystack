package paystack

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
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

func (a *restClient) APICall(ctx context.Context, method string, endPointPath string, payload any, response any) error {
	var body *bytes.Buffer
	var apiRequest *http.Request
	var err error

	if payload != nil {
		payloadInBytes, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(payloadInBytes)
	}

	if payload != nil {
		apiRequest, err = http.NewRequestWithContext(ctx, method, a.baseUrl+endPointPath, body)
	} else {
		apiRequest, err = http.NewRequestWithContext(ctx, method, a.baseUrl+endPointPath, nil)
	}

	if err != nil {
		return err
	}
	err = a.setHeaders(apiRequest)
	if err != nil {
		return err
	}
	r, err := a.httpClient.Do(apiRequest)
	if err != nil {
		return err
	}
	return a.unMarshalResponse(r, response)
}

func (a *restClient) unMarshalResponse(httpResponse *http.Response, result any) error {
	raw, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(raw, result); err != nil {
		return err
	}

	// Get the value of the response interface
	value := reflect.ValueOf(result)

	// Check if it's a pointer and get the element it points to
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Check if it's a struct
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct of the generic model.Response struct, got %v", value.Kind())
	}

	// Check for StatusCode field
	statusCodeField := value.FieldByName("StatusCode")
	if !statusCodeField.IsValid() {
		return errors.New("response parameter has no StatusCode field, please ensure response type is of the generic model.Response struct")
	}

	// Check if StatusCode field is settable and of correct type
	if statusCodeField.CanSet() {
		switch statusCodeField.Kind() {
		case reflect.Int, reflect.Int32, reflect.Int64:
			statusCodeField.SetInt(int64(httpResponse.StatusCode))
		default:
			return errors.New("StatusCode field of the response parameter is not a valid integer")
		}
	} else {
		return errors.New("StatusCode field of the response parameter cannot be set")
	}

	// Check for Raw field
	rawField := value.FieldByName("Raw")
	if !rawField.IsValid() {
		return errors.New("response parameter has no Raw field, please ensure response type is of the generic model.Response struct")
	}

	// Check if Raw field is settable and of correct type
	if rawField.CanSet() {
		switch rawField.Kind() {
		case reflect.Slice:
			if rawField.Type().Elem().Kind() == reflect.Uint8 {
				rawField.SetBytes(raw)
			}
		default:
			return errors.New("StatusCode field of the response parameter is not a valid integer")
		}
	} else {
		return errors.New("StatusCode field of the response parameter cannot be set")
	}
	return nil
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
