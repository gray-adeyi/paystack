package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// TransferRecipientClient interacts with endpoints related to paystack transfer recipient resource
// that lets you create and manage beneficiaries that you send money to.
type TransferRecipientClient struct {
	*restClient
}

// NewTransferRecipientClient creates a TransferRecipientClient
func NewTransferRecipientClient(options ...ClientOptions) *TransferRecipientClient {
	client := NewClient(options...)
	return client.TransferRecipients
}

// Create lets you create a new recipient. A duplicate account number will lead to the retrieval of the existing record.
//
// Default response: models.Response[models.TransferRecipient]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//		"github.com/gray-adeyi/paystack/enum"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.TransferRecipient]
//		if err := client.TransferRecipients.Create(context.TODO(),enum.RecipientTypeNuban,"Tolu Robert","01000000010", "058", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.TransferRecipients.Create(context.TODO(),enum.RecipientTypeNuban,"Tolu Robert","01000000010", "058", &response, p.WithOptionalParameter("currency","NGN"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/transfer-recipient/
func (t *TransferRecipientClient) Create(ctx context.Context, recipientType enum.RecipientType, name string, accountNumber string,
	bankCode string, response any, optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"type":           recipientType,
		"name":           name,
		"account_number": accountNumber,
		"bank_code":      bankCode,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPost, "/transferrecipient", payload, response)
}

// BulkCreate lets you create multiple transfer recipients in batches. A duplicate account number will lead to the retrieval of the existing record.
//
// Default response: models.Response[models.TransferRecipientBulkCreateData]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//		
// 		batch := []map[string]any{
//		{"type":"nuban", "name" : "Habenero Mundane", "account_number": "0123456789","bank_code":"033","currency": "NGN"},
//		{"type":"nuban","name" : "Soft Merry","account_number": "98765432310","bank_code": "50211","currency": "NGN"},
//		}
// 
//		var response models.Response[models.TransferRecipientBulkCreateData]
//		if err := client.TransferRecipients.BulkCreate(context.TODO(),batch, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferRecipientClient) BulkCreate(ctx context.Context, batch, response any) error {
	payload := map[string]any{
		"batch": batch,
	}
	payload["batch"] = batch

	return t.APICall(ctx, http.MethodPost, "/transferrecipient/bulk", payload, response)
}

// All lets you retrieve transfer recipients available on your Integration
//
// Default response: models.Response[[]models.TransferRecipient]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[[]models.TransferRecipient]
//		if err := client.TranferRecipients.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.TransferRecipients.All(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/transfer-recipient/
func (t *TransferRecipientClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/transferrecipient", queries...)
	return t.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve the details of a transfer recipient
//
// Default response: models.Response[models.TranferRecipient]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[models.TransferRecipient]
//		if err := client.TransferRecipients.FetchOne(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferRecipientClient) FetchOne(ctx context.Context, idOrCode string, response any) error {
	return t.APICall(ctx, http.MethodGet, fmt.Sprintf("/transferrecipient/%s", idOrCode), nil, response)
}

// Update lets you update transfer recipients available on your Integration
//
// Default response: models.Response[struct{}]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//		"github.com/gray-adeyi/paystack/enum"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[struct{}]
//		if err := client.TransferRecipients.Update(context.TODO(),"<idOrCode>", "Rick Sanchez", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.TransferRecipients.Update(context.TODO(),"<idOrCode>", "Rick Sanchez", &response, p.WithOptionalParameter("email","johndoe@example.com"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/transfer-recipient/
func (t *TransferRecipientClient) Update(ctx context.Context, idOrCode string, name string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"name": name,
	}

	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return t.APICall(ctx, http.MethodPut, fmt.Sprintf("/transferrecipient/%s", idOrCode), payload, response)
}

// Delete lets you delete a transfer recipient (sets the transfer recipient to inactive)
//
// Default response: models.Response[struct{}]
//
// Example:
//
//	import (
//		"context"
//		"fmt"
//
//		p "github.com/gray-adeyi/paystack"
//		"github.com/gray-adeyi/paystack/models"
//	)
//
//	func main() {
//		client := p.NewClient(p.WithSecretKey("<paystack-secret-key>"))
//
//		var response models.Response[struct{}]
//		if err := client.TransferRecipients.Delete(context.TODO(),"<idOrCode>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (t *TransferRecipientClient) Delete(ctx context.Context, idOrCode string, response any) error {
	return t.APICall(ctx, http.MethodDelete, fmt.Sprintf("/transferrecipient/%s", idOrCode), nil, response)
}
