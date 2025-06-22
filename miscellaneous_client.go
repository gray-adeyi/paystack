package paystack

import (
	"context"
	"net/http"

	"github.com/gray-adeyi/paystack/enum"
)

// MiscellaneousClient interacts with endpoints related to paystack Miscellaneous resource that
// provides information that is relevant to other client methods
type MiscellaneousClient struct {
	*restClient
}

// NewMiscellaneousClient creates a MiscellaneousClient
func NewMiscellaneousClient(options ...ClientOptions) *MiscellaneousClient {
	client := NewClient(options...)
	return client.Miscellaneous
}

// Banks lets you retrieve a list of all supported banks and their properties
//
// Default response: models.Response[[]models.Bank]
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
//		var response models.Response[[]models.Bank]
//		if err := client.Miscellaneous.Banks(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Miscellaneous.Banks(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/miscellaneous/
func (p *MiscellaneousClient) Banks(ctx context.Context, country enum.Country, response any, queries ...Query) error {
	countryFullNames := map[enum.Country]string{
		enum.CountryNigeria:     "nigeria",
		enum.CountryGhana:       "ghana",
		enum.CountrySouthAfrica: "south africa",
		enum.CountryKenya:       "kenya",
		enum.CountryCoteDIvoire: "côte d'ivoire",
		enum.CountryEgypt:       "egypt",
		enum.CountryRwanda:      "rwanda",
	}
	fullName := countryFullNames[country]
	queries = append(queries, WithQuery("country", fullName))
	url := AddQueryParamsToUrl("/bank", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil, response)
}

// Countries let you retrieve a list of countries that Paystack currently supports
//
// Default response: models.Response[[]models.PaystackSupportedCountry]
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
//		var response models.Response[[]models.PastackSupportedCountry]
//		if err := client.Miscellaneous.Countries(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//	}
func (p *MiscellaneousClient) Countries(ctx context.Context, response any) error {
	return p.APICall(ctx, http.MethodGet, "/country", nil, response)
}

// States lets you retrieve a list of states for a country for address Verification
//
// Default response: models.Response[[]models.State]
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
//		var response models.Response[[]models.State]
//		if err := client.Miscellaneous.States(context.TODO(), &response); err != nil {
//			panic(err)
//		}
//
//		fmt.Println(response)
//
//		// With query parameters
//		// err := client.Miscellaneous.States(context.TODO(), &response,p.WithQuery("perPage","50"), p.WithQuery("page","2"))
//	}
//
// For supported query parameters, see:
// https://paystack.com/docs/api/miscellaneous/
func (p *MiscellaneousClient) States(ctx context.Context, response any, queries ...Query) error {
	url := AddQueryParamsToUrl("/address_verification/states", queries...)
	return p.APICall(ctx, http.MethodGet, url, nil, response)
}
