package paystack

import (
	"fmt"
	"net/http"
)

type BulkChargeClient struct {
	*baseAPIClient
}

func (b *BulkChargeClient) Initiate(charges interface{}) (*Response, error) {
	return b.APICall(http.MethodPost, "/bulkcharge", charges)
}

func (b *BulkChargeClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/bulkcharge", queries...)
	return b.APICall(http.MethodGet, url, nil)
}

func (b *BulkChargeClient) FetchOne(idOrCode string) (*Response, error) {
	return b.APICall(http.MethodGet, fmt.Sprintf("/bulkcharge/%s", idOrCode), nil)
}

func (b *BulkChargeClient) Charges(idOrCode string, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl(fmt.Sprintf("/bulkcharge/%s/charges", idOrCode), queries...)
	return b.APICall(http.MethodGet, url, nil)
}

func (b *BulkChargeClient) Pause(idOrCode string) (*Response, error) {
	return b.APICall(http.MethodGet, fmt.Sprintf("/bulkcharge/pause/%s", idOrCode), nil)
}

func (b *BulkChargeClient) Resume(idOrCode string) (*Response, error) {
	return b.APICall(http.MethodGet, fmt.Sprintf("/bulkcharge/resume/%s", idOrCode), nil)
}
