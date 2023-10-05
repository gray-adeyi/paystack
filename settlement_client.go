package paystack

import (
	"fmt"
	"net/http"
)

type SettlementClient struct {
	*baseAPIClient
}

func (s *SettlementClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/settlement", queries...)
	return s.APICall(http.MethodGet, url, nil)
}

func (s *SettlementClient) AllTransactions(settlementId string, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl(fmt.Sprintf("/settlement/%s", settlementId), queries...)
	return s.APICall(http.MethodGet, url, nil)
}
