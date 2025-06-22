package paystack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// VerificationClient interacts with endpoints related to paystack Verification resource
// that allows you to perform KYC processes.
type VerificationClient struct {
	*restClient
}

// NewVerificationClient creates a VerificationClient
//
//	Example
//
//	import p "github.com/gray-adeyi/paystack"
//
//	vClient := p.NewVerificationClient(p.WithSecretKey("<paystack-secret-key>"))
func NewVerificationClient(options ...ClientOptions) *VerificationClient {
	client := NewClient(options...)
	return client.Verification
}

// ResolveAccount lets you confirm an account belongs to the right customer
//
// Default response: models.Response[models.BankAccountInfo]
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
//		var response models.Response[models.BankAccountInfo]
//		if err := client.Verification.ResolveAccount(context.TODO(), &response,p.WithQuery("account_number","0022728151"),p.WithQuery("bank_code","063")); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (v *VerificationClient) ResolveAccount(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/bank/resolve", queries...)
	return v.APICall(ctx, http.MethodGet, url, nil, response)
}

// ValidateAccount lets you confirm the authenticity of a customer's account number before sending money
//
// Default response: models.Response[models.AccountVerificationInfo]
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
//		var response models.Response[models.AccountVerificationInfo]
//		if err := client.Verification.ValidateAccount(context.TODO(),"Ann Bron","0123456789",enum.AccountTypePersonal,"632005",enum.CountrySouthAfrica,enum.DocumentIdentityNumber, &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With optional parameters
//		// err := client.Verification.ValidateAccount(context.TODO(),"Ann Bron","0123456789",enum.AccountTypePersonal,"632005",enum.CountrySouthAfrica,enum.DocumentIdentityNumber, &response, p.WithOptionalPayload("document_number","1234567890123"))
//	}
//
// For supported optional parameters, see:
// https://paystack.com/docs/api/verification/
func (v *VerificationClient) ValidateAccount(ctx context.Context, accountName string, accountNumber string, accountType enum.AccountType, bankCode string, countryCode enum.Country, documentType enum.Document, response any, optionalPayloads ...OptionalPayload) error {
	payload := map[string]any{
		"account_name":   accountName,
		"account_number": accountNumber,
		"account_type":   accountType,
		"bank_code":      bankCode,
		"country_code":   countryCode,
		"document_type":  documentType,
	}
	for _, optionalPayloadParameter := range optionalPayloads {
		payload = optionalPayloadParameter(payload)
	}
	return v.APICall(ctx, http.MethodGet, "/bank/validate", payload, response)
}

// ResolveBin lets you retrieve more information about a customer's card
//
// Default response: models.Response[models.CardBin]
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
//		var response models.Response[models.CardBin]
//		if err := client.Verification.ResolveBin(context.TODO(),"539983", &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (v *VerificationClient) ResolveBin(ctx context.Context, bin string, response any) error {
	return v.APICall(ctx, http.MethodGet, fmt.Sprintf("/decision/bin/%s", bin), nil, response)
}
