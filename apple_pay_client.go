package paystack

import "net/http"

type ApplePayClient struct {
	*baseAPIClient
}

func (a *ApplePayClient) Register(domainName string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["domainName"] = domainName
	return a.APICall(http.MethodPost, "/apple-pay/domain", payload)
}

func (a *ApplePayClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/apple-pay/domain", queries...)
	return a.APICall(http.MethodGet, url, nil)
}

func (a *ApplePayClient) Unregister(domainName string) (*Response, error) {
	payload := make(map[string]interface{})
	payload["domainName"] = domainName
	return a.APICall(http.MethodDelete, "/apple-pay/domain", payload)
}
