package client

import (
	"context"
	"fmt"
	"net/http"
)

// TerminalEvent specifies the supported terminal event by paystack
type TerminalEvent = string

const TerminalEventInvoice TerminalEvent = "invoice"
const TerminalEventTransaction TerminalEvent = "transaction"

// TerminalClient interacts with endpoints related to paystack Terminal resource that allows you to
// build delightful in-person payment experiences.
type TerminalClient struct {
	*restClient
}

// NewTerminalClient creates a TerminalClient
//
// Example:
//
//	import p "github.com/gray-adeyi/paystack"
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
func NewTerminalClient(options ...ClientOptions) *TerminalClient {
	client := NewPaystackClient(options...)
	return client.Terminals
}

// SendEvent lets you send an event from your application to the Paystack Terminal
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a terminal client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Terminals field is a `TerminalClient`
//	// Therefore, this is possible
//	// payload := map[string]interface{}{
//	//	"id": 7895939, "reference": 4634337895939 }
//	// resp, err := paystackClient.Terminals.SendEvent("30",p.TerminalEventInvoice,"process", payload)
//
//	payload := map[string]interface{}{
//	"id": 7895939, "reference": 4634337895939 }
//	resp, err := terminalClient.SendEvent("30",p.TerminalEventInvoice,"process", payload)
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TerminalClient) SendEvent(ctx context.Context, terminalId string, eventType TerminalEvent, action string, data, response any) error {
	payload := map[string]any{
		"type":   eventType,
		"action": action,
		"data":   data,
	}

	return t.APICall(ctx, http.MethodPost, fmt.Sprintf("/terminal/%s/event", terminalId), payload, response)
}

// EventStatus lets you check the status of an event sent to the Terminal
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a terminal client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Terminals field is a `TerminalClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Terminals.EventStatus("30","616d721e8c5cd40a0cdd54a6")
//
//	payload := map[string]interface{}{
//	"id": 7895939, "reference": 4634337895939 }
//	resp, err := terminalClient.EventStatus("30","616d721e8c5cd40a0cdd54a6")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TerminalClient) EventStatus(ctx context.Context, terminalId string, eventId string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/terminal/%s/event/%s", terminalId, eventId), nil, response)
}

// TerminalStatus lets you check the availability of a Terminal before sending an event to it
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a transaction client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Terminals field is a `TerminalClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Terminals.TerminalStatus("30")
//
//	resp, err := terminalClient.TerminalStatus("30")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TerminalClient) TerminalStatus(ctx context.Context, terminalId string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/terminal/%s/presence", terminalId), nil, response)
}

// All lets you retrieve the Terminals available on your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a terminal client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Terminals field is a `TerminalClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Transactions.All()
//
//	// All also accepts queries, so say you want to customize how many Terminals to retrieve,
//	// you can write it like so.
//	// resp, err := txnClient.All(p.WithQuery("perPage","50"))
//
//	// see https://paystack.com/docs/api/terminal/#list for supported query parameters
//
//	resp, err := terminalClient.All()
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TerminalClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/terminal", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you get the details of a Terminal
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a terminal client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Terminals field is a `TerminalClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Terminals.FetchOne("<terminalId>")
//
//	resp, err := terminalClient.FetchOne("<terminalId>")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TerminalClient) FetchOne(ctx context.Context, terminalId string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/terminal/%s", terminalId), nil, response)
}

// Update lets you update the details of a Terminal
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a terminal client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Terminals field is a `TerminalClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Terminals.Update("<terminalId>", "New Terminal","somewhere on earth")
//
//	resp, err := terminalClient.Update("<terminalId>", "New Terminal","somewhere on earth")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TerminalClient) Update(ctx context.Context, terminalId string, name string, address string, response any) error {
	payload := map[string]any{
		"name":    name,
		"address": address,
	}

	return t.APICall(ctx, http.MethodPut, fmt.Sprintf("/terminal/%s", terminalId), payload, response)
}

// Commission lets you activate your debug device by linking it to your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a terminal client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Terminals field is a `TerminalClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Terminals.Commission("<serialNumber>")
//
//	resp, err := terminalClient.Commission("<serialNumber>")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TerminalClient) Commission(ctx context.Context, serialNumber string, response any) error {
	payload := map[string]any{
		"serial_number": serialNumber,
	}

	return t.APICall(ctx, http.MethodPost, "/terminal/commission_device", payload, response)
}

// Decommission lets you unlink your debug device from your Integration
//
// Example:
//
//	import (
//		"fmt"
//		p "github.com/gray-adeyi/paystack"
//		"encoding/json"
//	)
//
//	terminalClient := p.NewTerminalClient(p.WithSecretKey("<paystack-secret-key>"))
//	// Alternatively, you can access a terminal client from an APIClient
//	// paystackClient := p.NewAPIClient(p.WithSecretKey("<paystack-secret-key>"))
//	// paystackClient.Terminals field is a `TerminalClient`
//	// Therefore, this is possible
//	// resp, err := paystackClient.Terminals.Commission("<serialNumber>")
//
//	resp, err := terminalClient.Commission("<serialNumber>")
//	if err != nil {
//		panic(err)
//	}
//	// you can have data be a custom structure based on the data your interested in retrieving from
//	// from paystack for simplicity, we're using `map[string]interface{}` which is sufficient to
//	// to serialize the json data returned by paystack
//	data := make(map[string]interface{})
//
//	err := json.Unmarshal(resp.Data, &data); if err != nil {
//		panic(err)
//	}
//	fmt.Println(data)
func (t *TerminalClient) Decommission(ctx context.Context, serialNumber string, response any) error {
	payload := map[string]string{
		"serial_number": serialNumber,
	}

	return t.APICall(ctx, http.MethodPost, "/terminal/decommission_device", payload, response)
}
