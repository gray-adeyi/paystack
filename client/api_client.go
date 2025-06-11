package client

import (
	"net/http"
)

// PaystackClient is a struct that has other dedicated clients bound to it. This provides a convenience for interacting
// with all of paystack's endpoints in your Go project. It should not be instantiated directly but interacting but
// via the NewAPIClient function. As stated above, it has other dedicated clients bound to it as field, therefore,
// after creating an instance of the PaystackClient type. You can access the associated functions of each dedicated client
// via its field name.
//
//	Example
//	import (
//		p "github.com/gray-adeyi/paystack"
//		"context"
//		)
//
//	client := p.NewAPIClient(p.WithSecretKey("<your-paystack-secret-key>"))
//	resp, err := client.Transactions.Verify(context.TODO(),"<reference>")
type PaystackClient struct {
	*restClient

	// Transactions let you interact with endpoints related to paystack Transaction resource
	// that allows you to create and manage payments on your Integration.
	Transactions *TransactionClient

	// TransactionSplits lets you interact with endpoints related to paystack Transaction Split resource
	// that allows you to split the settlement for a transaction across a payout account, and one or
	// more subaccounts.
	TransactionSplits *TransactionSplitClient

	// Terminals let you interact with endpoints related to paystack Terminal resource that allows you to
	// build delightful in-person payment experiences.
	Terminals *TerminalClient

	// Customers let you interact with endpoints related to paystack Customer resource
	// that allows you to create and manage Customers on your Integration.
	Customers *CustomerClient

	// DedicatedVirtualAccounts lets you interact with endpoints related to paystack dedicated virtual account
	// resource that enables Nigerian merchants to manage unique payment accounts of their Customers.
	DedicatedVirtualAccounts *DedicatedVirtualAccountClient

	// ApplePay lets you interact with endpoints related to paystack Apple Pay resource that
	// lets you register your application's top-level domain or subdomain.
	ApplePay *ApplePayClient

	// SubAccounts lets you interact with endpoints related to paystack subaccount resource that lets you
	// create and manage subaccounts on your Integration. Subaccounts can be used to split payment
	// between two accounts (your main account and a subaccount).
	SubAccounts *SubAccountClient

	// Plans lets you interact with endpoints related to paystack plan resource that lets you
	// create and manage installment payment options on your Integration.
	Plans *PlanClient

	// Subscriptions let you interact with endpoints related to paystack subscription resource that lets you
	// create and manage recurring payment on your Integration.
	Subscriptions *SubscriptionClient

	// Products let you interact with endpoints related to paystack product resource that allows you to create and
	// manage inventories on your Integration.
	Products *ProductClient

	// PaymentPages let you interact with endpoints related to paystack payment page resource
	// that lets you provide a quick and secure way to collect payment for Products.
	PaymentPages *PaymentPageClient

	// PaymentRequests let you interacts with endpoints related to paystack payment request resource that lets you manage requests
	// for payment of goods and services.
	PaymentRequests *PaymentRequestClient

	// Settlements let you interact with endpoints related to paystack settlement resource that lets you
	// gain insights into payouts made by Paystack to your bank account.
	Settlements *SettlementClient

	// TransferRecipients let you interact with endpoints related to paystack transfer recipient resource
	// that lets you create and manage beneficiaries that you send money to.
	TransferRecipients *TransferRecipientClient

	// Transfers let you interact with endpoints related to paystack transfer resource that lets you
	// automate sending money to your Customers.
	Transfers *TransferClient

	// TransferControl let you interact with endpoints related to paystack transfer control resource that lets
	// you manage settings of your Transfers.
	TransferControl *TransferControlClient

	// BulkCharges let you interact with endpoints related to paystack bulk Charges resource that lets
	// you create and manage multiple recurring payments from your Customers.
	BulkCharges *BulkChargeClient

	// Integration let you interact with endpoints related to paystack Integration resource
	// that lets you manage some settings on your Integration.
	Integration *IntegrationClient

	// Charge let you interact with endpoints related to paystack charge resource that
	// lets you configure a payment channel of your choice when initiating a payment.
	Charges *ChargeClient

	// Disputes let you interact with endpoint related to paystack dispute resource that lets you
	// manage transaction Disputes on your Integration.
	Disputes *DisputeClient

	// Refunds let you interact with endpoints related to paystack refund resource that lets you
	// create and manage transaction Refunds.
	Refunds *RefundClient

	// Verification let you interact with endpoints related to paystack Verification resource
	// that allows you to perform KYC processes.
	Verification *VerificationClient

	// Miscellaneous let you interact with endpoints related to paystack Miscellaneous resource that
	// provides information that is relevant to other client methods
	Miscellaneous *MiscellaneousClient
}

// NewPaystackClient lets you create an APIClient. it can accept zero to many client options
//
//	Example
//	import p "github.com/gray-adeyi/paystack"
//
//	client := p.NewPaystackClient(p.WithSecretKey("<your-paystack-secret-key>"))
func NewPaystackClient(options ...ClientOptions) *PaystackClient {
	restClient := &restClient{
		baseUrl:    BaseUrl,
		httpClient: &http.Client{},
	}

	for _, opts := range options {
		opts(restClient)
	}
	return &PaystackClient{
		restClient: restClient,
		Transactions: &TransactionClient{
			restClient,
		},
		TransactionSplits: &TransactionSplitClient{
			restClient,
		},
		Terminals: &TerminalClient{
			restClient,
		},
		Customers: &CustomerClient{
			restClient,
		},
		DedicatedVirtualAccounts: &DedicatedVirtualAccountClient{
			restClient,
		},
		ApplePay: &ApplePayClient{
			restClient,
		},
		SubAccounts: &SubAccountClient{
			restClient,
		},
		Plans: &PlanClient{
			restClient,
		},
		Subscriptions: &SubscriptionClient{
			restClient,
		},
		Products: &ProductClient{
			restClient,
		},
		PaymentPages: &PaymentPageClient{
			restClient,
		},
		PaymentRequests: &PaymentRequestClient{
			restClient,
		},
		Settlements: &SettlementClient{
			restClient,
		},
		TransferControl: &TransferControlClient{
			restClient,
		},
		TransferRecipients: &TransferRecipientClient{
			restClient,
		},
		Transfers: &TransferClient{
			restClient,
		},
		BulkCharges: &BulkChargeClient{
			restClient,
		},
		Integration: &IntegrationClient{
			restClient,
		},
		Charges: &ChargeClient{
			restClient,
		},
		Disputes: &DisputeClient{
			restClient,
		},
		Refunds: &RefundClient{
			restClient,
		},
		Verification: &VerificationClient{
			restClient,
		},
		Miscellaneous: &MiscellaneousClient{
			restClient,
		},
	}
}
