package paystack

import "net/http"

type MiscellaneousClient struct {
	*baseAPIClient
}

func (p *MiscellaneousClient) Banks(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/bank", queries...)
	return p.APICall(http.MethodGet, url, nil)
}

func (p *MiscellaneousClient) Countries(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/country", queries...)
	return p.APICall(http.MethodGet, url, nil)
}

func (p *MiscellaneousClient) States(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/address_verification/states", queries...)
	return p.APICall(http.MethodGet, url, nil)
}
