package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// DisputeClient interacts with endpoint related to paystack dispute resource that lets you
// manage transaction Disputes on your Integration.
type DisputeClient struct {
	*restClient
}

// NewDisputeClient creates a DisputeClient
func NewDisputeClient(options ...ClientOptions) *DisputeClient {
	client := NewClient(options...)
	return client.Disputes
}

// All lets you retrieve Disputes filed against you
//
// Default response: models.Response[[]models.Dispute]
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
//		var response models.Response[[]models.Dispute]
//		if err := client.Disputes.All(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Disputes.All(context.TODO(), &response,p.WithQuery("from","2023-01-01"), p.WithQuery("to","2023-12-31"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/dispute/
func (d *DisputeClient) All(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/dispute", queries...)
	return d.APICall(ctx, http.MethodGet, url, nil, response)
}

// FetchOne lets you retrieve more details about a dispute.
//
// Default response: models.Response[models.Dispute]
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
//		var response models.Response[models.Dispute]
//		if err := client.Disputes.FetchOne(context.TODO(),"<id>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (d *DisputeClient) FetchOne(ctx context.Context, id string, response any) error {
	return d.APICall(ctx, http.MethodGet, fmt.Sprintf("/dispute/%s", id), nil, response)
}

// AllTransactionDisputes lets you retrieve Disputes for a particular transaction
//
// Default response: models.Response[[]models.Dispute]
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
//		var response models.Response[[]models.Dispute]
//		if err := client.Disputes.AllTransactionDisputes(context.TODO(),<transactionId>, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (d *DisputeClient) AllTransactionDisputes(ctx context.Context, transactionId string, response any) error {
	return d.APICall(ctx, http.MethodGet, fmt.Sprintf("/dispute/transaction/%s", transactionId), nil, response)
}

// Update lets you update the details of a dispute on your Integration
//
// Default response: models.Response[models.Dispute]
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
//		var response models.Response[models.Dispute]
//		if err := client.Disputes.Update(context.TODO(),"<id>","<referenceAmount>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Disputes.Update(context.TODO(),"<id>","<referenceAmount>", &response,p.WithOptionalParameter("uploaded_filename","Disputes.pdf"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/dispute/
func (d *DisputeClient) Update(ctx context.Context, id string, referenceAmount int, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"reference_amount": referenceAmount,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(ctx, http.MethodPut, fmt.Sprintf("/dispute/%s", id), payload, response)
}

// AddEvidence lets you provide evidence for a dispute
//
// Default response: models.Response[models.DisputeEvidence]
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
//		var response models.Response[models.DisputeEvidence]
//		if err := client.Disputes.AddEvidence(context.TODO(),"<id>", "johndoe@example.com", "John Doe", "5085072209", "claim for buying product", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Disputes.AddEvidence(context.TODO(),"<id>", "johndoe@example.com", "John Doe", "5085072209", "claim for buying product", &response,p.WithOptionalParameter("delivery_address", "3a ladoke street ogbomoso"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/dispute/
func (d *DisputeClient) AddEvidence(ctx context.Context, id string, customerEmail string,
	customerName string, customerPhone string, serviceDetails string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"customer_email":  customerEmail,
		"customer_name":   customerName,
		"customer_phone":  customerPhone,
		"service_details": serviceDetails,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(ctx, http.MethodPost, fmt.Sprintf("/dispute/%s/evidence", id), payload, response)
}

// UploadUrl lets you retrieve Disputes for a particular transaction
//
// Default response: models.Response[models.DisputeUploadInfo]
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
//		var response models.Response[models.DisputeUploadInfo]
//		if err := client.Disputes.UploadUrl(context.TODO(),"<disputeId>", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Disputes.UploadUrl(context.TODO(),"<disputeId>", &response,p.WithQuery("upload_filename","filename.pdf"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/dispute/
func (d *DisputeClient) UploadUrl(ctx context.Context, id string, response any, queries ...Query) error {
	url := AddQueryParamsToUrl(fmt.Sprintf("/dispute/%s/upload_url", id), queries...)
	return d.APICall(ctx, http.MethodPost, url, nil, response)
}

// Resolve lets you resolve a dispute on your Integration
//
// Default response: models.Response[models.Dispute]
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
//		var response models.Response[models.Dispute]
//		if err := client.Disputes.Resolve(context.TODO(),enum.ResolutionMerchantAccepted,"Merchant accepted","qesp8a4df1xejihd9x5q",1002, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Disputes.Resolve(context.TODO(),enum.ResolutionMerchantAccepted,"Merchant accepted","qesp8a4df1xejihd9x5q",1002, &response,p.WithOptionalParameter("evidence", "<evidenceId>"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/dispute/
func (d *DisputeClient) Resolve(ctx context.Context, id string, resolution enum.Resolution, message string,
	refundAmount int, uploadedFilename string, response any,
	optionalPayloadParameters ...OptionalPayloadParameter) error {
	payload := map[string]any{
		"resolution":        resolution,
		"message":           message,
		"refund_amount":     refundAmount,
		"uploaded_filename": uploadedFilename,
	}
	for _, optionalPayloadParameter := range optionalPayloadParameters {
		payload = optionalPayloadParameter(payload)
	}
	return d.APICall(ctx, http.MethodPut, fmt.Sprintf("/dispute/%s/resolve", id), payload, response)
}

// Export lets you export Disputes available on your Integration
//
// Default response: models.Response[models.DisputeExportInfo]
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
//		var response models.Response[models.DisputeExportInfo]
//		if err := client.Disputes.Export(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Disputes.Export(context.TODO(), &response,p.WithQuery("from","2023-01-01"), p.WithQuery("to","2023-12-31"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/dispute/
func (d *DisputeClient) Export(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/dispute/export", queries...)
	return d.APICall(ctx, http.MethodGet, url, nil, response)
}
