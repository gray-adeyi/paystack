package paystack

import (
	"fmt"
	"net/http"
)

type VerificationClient struct {
	*baseAPIClient
}

func (v *VerificationClient) ResolveAccount(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/bank/resolve", queries...)
	return v.APICall(http.MethodGet, url, nil)
}

func (v *VerificationClient) ValidateAccount(accountName string, accountNumber string, accountType string, bankCode string, countryCode string, documentType string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"account_name":   accountName,
		"account_number": accountNumber,
		"account_type":   accountType,
		"bank_code":      bankCode,
		"country_code":   countryCode,
		"document_type":  documentType,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return v.APICall(http.MethodGet, "/bank/validate", payload)
}

func (v *VerificationClient) ResolveBIN(bin string) (*Response, error) {
	return v.APICall(http.MethodGet, fmt.Sprintf("/decision/bin/%s", bin), nil)
}
