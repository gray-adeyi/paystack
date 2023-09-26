package paystack

import (
	"fmt"
	"net/http"
)

type TerminalEvent = string

const TerminalEventInvoice TerminalEvent = "invoice"
const TerminalEventTransaction TerminalEvent = "transaction"

type TerminalClient struct {
	*baseAPIClient
}

func (t *TerminalClient) SendEvent(terminalId string, eventType TerminalEvent, action string, data interface{}) (*Response, error) {
	payload := make(map[string]interface{})
	payload["type"] = eventType
	payload["action"] = action
	payload["data"] = data

	return t.APICall(http.MethodPost, fmt.Sprintf("/terminal/%s/event", terminalId), payload)
}

func (t *TerminalClient) EventStatus(terminalId string, eventId string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/terminal/%s/event/%s", terminalId, eventId), nil)
}

func (t *TerminalClient) TerminalStatus(terminalId string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/terminal/%s/presence", terminalId), nil)
}

func (t *TerminalClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/terminal", queries...)
	return t.APICall(http.MethodGet, url, nil)
}

func (t *TerminalClient) FetchOne(terminalId string) (*Response, error) {
	return t.APICall(http.MethodGet, fmt.Sprintf("/terminal/%s", terminalId), nil)
}

func (t *TerminalClient) Update(terminalId string, name string, address string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["name"] = name
	payload["address"] = address

	return t.APICall(http.MethodPut, fmt.Sprintf("/terminal/%s", terminalId), payload)
}

func (t *TerminalClient) Commission(serialNumber string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["serial_number"] = serialNumber

	return t.APICall(http.MethodPost, "/terminal/commission_device", payload)
}

func (t *TerminalClient) Decommission(serialNumber string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["serial_number"] = serialNumber

	return t.APICall(http.MethodPost, "/terminal/decommission_device", payload)
}
