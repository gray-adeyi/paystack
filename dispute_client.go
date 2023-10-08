package paystack

import (
	"fmt"
	"net/http"
)

type DisputeClient struct {
	*baseAPIClient
}

func (d *DisputeClient) All(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/dispute", queries...)
	return d.APICall(http.MethodGet, url, nil)
}

func (d *DisputeClient) FetchOne(id string) (*Response, error) {
	return d.APICall(http.MethodGet, fmt.Sprintf("/dispute/%s", id), nil)
}

func (d *DisputeClient) AllTransactionDisputes(transactionId string) (*Response, error) {
	return d.APICall(http.MethodGet, fmt.Sprintf("/dispute/transaction/%s", transactionId), nil)
}

func (d *DisputeClient) Update(id string, referenceAmount int, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"reference_amount": referenceAmount,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(http.MethodPut, fmt.Sprintf("/dispute/%s", id), payload)
}

func (d *DisputeClient) AddEvidence(id string, customerEmail string, customerName string, customerPhone string, serviceDetails string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"customer_email":  customerEmail,
		"customer_name":   customerName,
		"customer_phone":  customerPhone,
		"service_details": serviceDetails,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(http.MethodPost, fmt.Sprintf("/dispute/%s/evidence", id), payload)
}

func (d *DisputeClient) UploadURL(id string, queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl(fmt.Sprintf("/dispute/%s/upload_url", id), queries...)
	return d.APICall(http.MethodPost, url, nil)
}

func (d *DisputeClient) Resolve(id string, resolution string, message string, refundAmount int, uploadedFilename string, optionalPayloadParameters ...OptionalPayloadParameter) (*Response, error) {
	payload := map[string]interface{}{
		"resolution":        resolution,
		"message":           message,
		"refund_amount":     refundAmount,
		"uploaded_filename": uploadedFilename,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(http.MethodPut, fmt.Sprintf("/dispute/%s/resolve", id), payload)
}

func (d *DisputeClient) Export(queries ...Query) (*Response, error) {
	url := AddQueryParamsToUrl("/dispute/export", queries...)
	return d.APICall(http.MethodGet, url, nil)
}
