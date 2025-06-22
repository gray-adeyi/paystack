package models

import (
	"time"

	"github.com/gray-adeyi/paystack/enum"
)

// Response is a struct containing the status code and data retrieved from paystack.
type Response[T any] struct {
	// StatusCode is the http status code returned from making an http request to Paystack
	StatusCode int 
	Status     bool   `json:"status"`
	Message    string `json:"message"`
	Data T       `json:"data"`
	Meta *Meta   `json:"meta"`
	Type *string `json:"type"`
	Code *string `json:"code"`
	Raw  []byte
}

type Meta struct {
	Total     int `json:"total"`
	Skipped   int `json:"skipped"`
	PerPage  string `json:"perPage"`
	Page      int `json:"page"`
	PageCount int `json:"pageCount"`
}

type State struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Abbreviation string `json:"abbreviation"`
}

type IntegrationTimeout struct {
	PaymentSessionTimeout time.Duration `json:"payment_session_timeout"`
}

type IntegrationBalance struct {
	Currency enum.Currency `json:"currency"`
	Balance  int           `json:"balance"`
}

type Integration struct {
	Key               string          `json:"key"`
	Name              string          `json:"name"`
	Logo              string          `json:"logo"`
	AllowedCurrencies []enum.Currency `json:"allowed_currencies"`
}

type ApplePayDomains struct {
	DomainNames []string `json:"domain_names"`
}

type BulkCharge struct {
	BatchCode      string             `json:"batch_code"`
	Reference      *string            `json:"reference"`
	Id             int                `json:"id"`
	Integration    *int               `json:"integration"`
	Domain         enum.Domain        `json:"domain"`
	Status         enum.DisputeStatus `json:"status"`
	TotalCharges   *int               `json:"total_charges"`
	PendingCharges *int               `json:"pending_charges"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}

type BulkChargeUnitCharge struct {
	Integration   int           `json:"integration"`
	Bulkcharge    int           `json:"bulkcharge"`
	Customer      Customer      `json:"customer"`
	Authorization Authorization `json:"authorization"`
	Transaction   Transaction   `json:"transaction"`
	Domain        enum.Domain   `json:"domain"`
	Amount        int           `json:"amount"`
	Currency      enum.Currency `json:"currency"`
	Status        string        `json:"status"`
	Id            string        `json:"id"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type Customer struct {
	Integration              *int            `json:"integration"`
	Id                       int             `json:"id"`
	FirstName                *string         `json:"first_name"`
	LastName                 *string         `json:"last_name"`
	Email                    string          `json:"email"`
	CustomerCode             string          `json:"customer_code"`
	Phone                    *string         `json:"phone"`
	Metadata                 map[string]any  `json:"metadata"`
	RiskAction               enum.RiskAction `json:"risk_action"`
	InternationalPhoneFormat *string         `json:"international_phone_format"`
	Identified               *bool           `json:"identified"`
	Identifications          *any            `json:"identifications"`
	Transactions             []Transaction   `json:"transactions"`
	Subscriptions            []Subscription  `json:"subscriptions"`
	Authorizations           []Authorization `json:"authorizations"`
	CreatedAt                *time.Time      `json:"created_at"`
	UpdatedAt                *time.Time      `json:"updated_at"`
	TotalTransactions        *int            `json:"total_transactions"`
	TotalTransactionValue    []any           `json:"total_transaction_value"`
	DedicatedAccount         *string         `json:"dedicated_account"`
	DedicatedAccounts        []any           `json:"dedicated_accounts"`
}

type Authorization struct {
	AuthorizationCode *string       `json:"authorization_code"`
	Bin               *string       `json:"bin"`
	Last4             *string       `json:"last4"`
	ExpMonth          *string       `json:"exp_month"`
	ExpYear           *string       `json:"exp_year"`
	Channel           *string       `json:"channel"`
	CardType          *string       `json:"card_type"`
	Bank              *string       `json:"bank"`
	CountryCode       *enum.Country `json:"country_code"`
	Brand             *string       `json:"brand"`
	Reusable          *bool         `json:"reusable"`
	Signature         *string       `json:"signature"`
	AccountName       *string       `json:"account_name"`
}

type InitTransaction struct {
	AuthorizationUrl string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type TransactionHistory struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

type TransactionLog struct {
	StartTime int                  `json:"start_time"`
	TimeSpent int                  `json:"time_spent"`
	Attempts  int                  `json:"attempts"`
	Errors    int                  `json:"errors"`
	Success   bool                 `json:"success"`
	Mobile    bool                 `json:"mobile"`
	Input     []any                `json:"input"`
	History   []TransactionHistory `json:"history"`
}

type TransactionTotal struct {
	TotalTransactions          int     `json:"total_transactions"`
	TotalVolume                int     `json:"total_volume"`
	TotalVolumeByCurrency      []Money `json:"total_volume_by_currency"`
	PendingTransfers           int     `json:"pending_tranfers"`
	PendingTransfersByCurrency []Money `json:"pending_transfers_by_currency"`
}

type TransactionExport struct {
	Path      string    `json:"path"`
	ExpiresAt time.Time `json:"expires_at"`
}

type TransactionSource struct {
	Source     string `json:"source"`
	Type       string `json:"type"`
	Identifier any    `json:"identifier"`
	EntryPoint string `json:"entry_point"`
}

type Transaction struct {
	Id                  int                    `json:"id"`
	Domain              enum.Domain            `json:"domain"`
	Status              enum.TransactionStatus `json:"status"`
	Reference           string                 `json:"reference"`
	Amount              int                    `json:"amount"`
	Message             *string                `json:"message"`
	GatewayResponse     *string                `json:"gateway_response"`
	PaidAt              time.Time              `json:"paid_at"`
	Channel             string                 `json:"channel"`
	Currency            enum.Currency          `json:"currency"`
	IpAddress           *string                `json:"ip_address"`
	Metadata            map[string]any         `json:"metadata"`
	Log                 *TransactionLog        `json:"log"`
	Fees                *int                   `json:"fees"`
	FeesSplit           any                    `json:"fees_split"`
	Customer            map[string]any         `json:"customer"`
	Authorization       map[string]any         `json:"authorization"`
	Plan                any                    `json:"plan"`
	Split               any                    `json:"split"`
	Subaccount          any                    `json:"subaccount"`
	OrderId             *string                `json:"order_id"`
	CreatedAt           time.Time              `json:"created_at"`
	RequestedAmount     *int                   `json:"requested_amount"`
	Source              *TransactionSource     `json:"source"`
	Connect             any                    `json:"connect"`
	PostTransactionData any                    `json:"post_transaction_data"`
}

type TransactionSplitSubAccount struct {
	Subaccount SubAccount `json:"subaccount"`
	Share      int        `json:"share"`
}

type TransactionSplit struct {
	Id               int                          `json:"id"`
	Name             string                       `json:"name"`
	Type             string                       `json:"type"`
	Currency         enum.Currency                `json:"currency"`
	Integration      int                          `json:"integration"`
	Domain           enum.Domain                  `json:"domain"`
	SplitCode        string                       `json:"split_code"`
	Active           bool                         `json:"active"`
	BearerType       string                       `json:"bearer_type"`
	BearerSubaccount *string                      `json:"bearer_subaccount"`
	CreatedAt        time.Time                    `json:"created_at"`
	UpdatedAt        time.Time                    `json:"updated_at"`
	IsDynamic        bool                         `json:"is_dynamic"`
	Subaccount       []TransactionSplitSubAccount `json:"subaccount"`
	TotalSubaccounts int                          `json:"total_subaccounts"`
}

type Subscription struct {
	Customer          any         `json:"customer"`
	Plan              any         `json:"plan"`
	Integration       int         `json:"integration"`
	Domain            enum.Domain `json:"domain"`
	Start             *int        `json:"start"`
	Status            string      `json:"status"`
	Quantity          *int        `json:"quantity"`
	Amount            int         `json:"amount"`
	SubscriptionCode  string      `json:"subscription_code"`
	EmailToken        string      `json:"email_token"`
	Authorization     any         `json:"authorization"`
	EasyCronId        *string     `json:"easy_cron_id"`
	CronExpression    string      `json:"cron_expression"`
	NextPaymentDate   *time.Time  `json:"next_payment_date"`
	OpenInvoice       any         `json:"open_invoice"`
	InvoiceLimit      int         `json:"invoice_limit"`
	Id                int         `json:"id"`
	SplitCode         *string     `json:"split_code"`
	CancelledAt       *time.Time  `json:"cancelled_at"`
	UpdatedAt         *time.Time  `json:"updated_at"`
	PaymentsCount     *int        `json:"payments_count"`
	MostRecentInvoice *Invoice    `json:"most_recent_invoice"`
	InvoiceHistory    []any       `json:"invoice_history"`
}

type SubscriptionLink struct {
	Link string `json:"link"`
}

type Invoice struct {
	Subscription     int         `json:"subscription"`
	Integration      int         `json:"integration"`
	Domain           enum.Domain `json:"domain"`
	InvoiceCode      string      `json:"invoice_code"`
	Customer         string      `json:"customer"`
	Transaction      int         `json:"transaction"`
	Amount           int         `json:"amount"`
	PeriodStart      string      `json:"period_start"`
	PeriodEnd        string      `json:"period_end"`
	Status           string      `json:"status"`
	Paid             any         `json:"paid"`
	Retries          int         `json:"retries"`
	Authorization    int         `json:"authorization"`
	PaidAt           time.Time   `json:"paid_at"`
	NextNotification string      `json:"next_notification"`
	NotificationFlag any         `json:"notification_flag"`
	Description      *string     `json:"description"`
	Id               int         `json:"id"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
}

type PaymentPage struct {
	Integration       int            `json:"integration"`
	Plan              *int           `json:"plan"`
	Domain            enum.Domain    `json:"domain"`
	Name              string         `json:"name"`
	Description       *string        `json:"description"`
	Amount            *int           `json:"amount"`
	Currency          enum.Currency  `json:"currency"`
	Slug              string         `json:"slug"`
	CustomFields      map[string]any `json:"custom_fields"`
	Type              string         `json:"type"`
	RedirectUrl       *string        `json:"redirect_url"`
	SuccessMessage    *string        `json:"success_message"`
	CollectPhone      bool           `json:"collect_phone"`
	Active            bool           `json:"active"`
	Published         bool           `json:"published"`
	Migrate           bool           `json:"migrate"`
	NotificationEmail *string        `json:"notification_email"`
	Metadata          map[string]any `json:"metadata"`
	SplitCode         *string        `json:"split_code"`
	Id                int            `json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	Products          *[]Product     `json:"products"`
}

type PaymentRequestNotification struct {
	SentAt  time.Time `json:"sent_at"`
	Channel string    `json:"channel"`
}

type LineItem struct {
	Name     string `json:"name"`
	Amount   int    `json:"amount"`
	Quantity int    `json:"quantity"`
}

type Tax struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type PaymentRequest struct {
	Id               int                          `json:"id"`
	Integration      any                          `json:"integration"`
	Domain           enum.Domain                  `json:"domain"`
	Amount           int                          `json:"amount"`
	Currency         enum.Currency                `json:"currency"`
	DueDate          time.Time                    `json:"due_date"`
	HasInvoice       *bool                        `json:"has_invoice"`
	InvoiceNumber    *int                         `json:"invoice_number"`
	Description      *string                      `json:"description"`
	PdfUrl           *string                      `json:"pdf_url"`
	LineItems        []LineItem                   `json:"line_items"`
	Tax              []Tax                        `json:"tax"`
	RequestCode      string                       `json:"request_code"`
	Status           string                       `json:"status"`
	Paid             bool                         `json:"paid"`
	PaidAt           time.Time                    `json:"paid_at"`
	Metadata         map[string]any               `json:"metadata"`
	Notifications    []PaymentRequestNotification `json:"notifications"`
	OfflineReference string                       `json:"offline_reference"`
	Customer         any                          `json:"customer"`
	CreatedAt        time.Time                    `json:"created_at"`
	Discount         *string                      `json:"discount"`
	SplitCode        *string                      `json:"split_code"`
	Transactions     *[]Transaction               `json:"transactions"`
	Archived         *bool                        `json:"archived"`
	Source           *string                      `json:"source"`
	PaymentMethod    any                          `json:"payment_method"`
	Note             any                          `json:"note"`
	AmountPaid       *int                         `json:"amount_paid"`
	UpdatedAt        time.Time                    `json:"updated_at"`
	PendingAmount    *int                         `json:"pending_amount"`
}

type Money struct {
	Currency enum.Currency `json:"currency"`
	Amount   int           `json:"amount"`
}

type PaymentRequestStat struct {
	Pending    []Money `json:"pending"`
	Successful []Money `json:"successful"`
	Total      []Money `json:"total"`
}

type PlanSubscriber struct {
	CustomerCode            string        `json:"customer_code"`
	CustomerFirstName       string        `json:"customer_first_name"`
	CustomerLastName        string        `json:"customer_last_name"`
	CustomerEmail           string        `json:"customer_email"`
	SubscriptionStatus      string        `json:"subscription_status"`
	Currency                enum.Currency `json:"currency"`
	CustomerTotalAmountPaid int           `json:"customer_total_amount_paid"`
}

type Plan struct {
	Subscriptions             *[]Subscription   `json:"subscriptions"`
	Pages                     *[]PaymentPage    `json:"pages"`
	Domain                    *enum.Domain      `json:"domain"`
	Name                      string            `json:"name"`
	PlanCode                  string            `json:"plan_code"`
	Description               *string           `json:"description"`
	Amount                    int               `json:"amount"`
	Interval                  enum.Interval     `json:"interval"`
	InvoiceLimit              *int              `json:"invoice_limit"`
	SendInvoices              bool              `json:"send_invoices"`
	SendSms                   bool              `json:"send_sms"`
	HostedPage                *bool             `json:"hosted_page"`
	HostedPageUrl             *bool             `json:"hosted_page_url"`
	HostedPageSummary         *string           `json:"hosted_page_summary"`
	Currency                  enum.Currency     `json:"currency"`
	Migrate                   *bool             `json:"migrate"`
	IsDeleted                 *bool             `json:"is_deleted"`
	IsArchived                *bool             `json:"is_archived"`
	Id                        int               `json:"id"`
	Integration               *int              `json:"integration"`
	CreatedAt                 *time.Time        `json:"created_at"`
	UpdatedAt                 *time.Time        `json:"updated_at"`
	TotalSubscriptions        *int              `json:"total_subscriptions"`
	ActiveSubscriptions       *int              `json:"active_subscriptions"`
	TotalSubscriptionsRevenue *int              `json:"total_subscriptions_revenue"`
	PagesCount                *int              `json:"pages_count"`
	SubscribersCount          *int              `json:"subscribers_count"`
	ActiveSubscriptionsCount  *int              `json:"active_subscriptions_count"`
	TotalRevenue              *int              `json:"total_revenue"`
	Subscribers               *[]PlanSubscriber `json:"subscribers"`
}

type SubAccount struct {
	Id                   int            `json:"id"`
	SubaccountCode       string         `json:"subaccount_code"`
	BusinessName         string         `json:"business_name"`
	Description          *string        `json:"description"`
	PrimaryContactName   *string        `json:"primary_contact_name"`
	PrimaryContactEmail  *string        `json:"primary_contact_email"`
	PrimaryContactPhone  *string        `json:"primary_contact_phone"`
	Metadata             map[string]any `json:"metadata"`
	PercentageCharge     *float64       `json:"percentage_charge"`
	SettlementBank       string         `json:"settlement_bank"`
	BankId               *int           `json:"bank_id"`
	AccountNumber        string         `json:"account_number"`
	Currency             enum.Currency  `json:"currency"`
	Active               any            `json:"active"`
	IsVerified           *bool          `json:"is_verified"`
	Integration          *int           `json:"integration"`
	Bank                 *int           `json:"bank"`
	ManagedByIntegration *int           `json:"managed_by_integration"`
	Domain               *enum.Domain   `json:"domain"`
	Migrate              *bool          `json:"migrate"`
	AccountName          *string        `json:"account_name"`
	Product              *string        `json:"product"`
}

type Product struct {
	Id                 int            `json:"id"`
	Name               string         `json:"name"`
	Description        *string        `json:"description"`
	ProductCode        string         `json:"product_code"`
	Slug               string         `json:"slug"`
	Currency           enum.Currency  `json:"currency"`
	Price              int            `json:"price"`
	Quantity           int            `json:"quantity"`
	QuantitySold       *int           `json:"quantity_sold"`
	Active             bool           `json:"active"`
	Domain             enum.Domain    `json:"domain"`
	Type               string         `json:"type"`
	InStock            bool           `json:"in_stock"`
	Unlimited          bool           `json:"unlimited"`
	Metadata           map[string]any `json:"metadata"`
	Files              *[]any         `json:"files"`
	FilePath           *string        `json:"file_path"`
	SuccessMessage     *string        `json:"success_message"`
	RedirectUrl        *string        `json:"redirect_url"`
	SplitCode          *string        `json:"split_code"`
	NotificationEmails *[]any         `json:"notification_emails"`
	MinimumOrderable   int            `json:"minimum_orderable"`
	MaximumOrderable   *int           `json:"maximum_orderable"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	Features           any            `json:"features"`
	DigitalAssets      *[]any         `json:"digital_assets"`
	VariantOptions     *[]any         `json:"variant_options"`
	IsShippable        bool           `json:"is_shippable"`
	ShippingFields     map[string]any `json:"shipping_fields"`
	Integration        int            `json:"integration"`
	LowStockAlert      int            `json:"low_stock_alert"`
	StockThreshold     any            `json:"stock_threshold"`
	ExpiresIn          any            `json:"expires_in"`
}

type Terminal struct {
	Id           int         `json:"id"`
	SerialNumber *string     `json:"serial_number"`
	DeviceMake   *string     `json:"device_make"`
	TerminalId   string      `json:"terminal_id"`
	Integration  int         `json:"integration"`
	Domain       enum.Domain `json:"domain"`
	Name         string      `json:"name"`
	Address      *string     `json:"address"`
	Status       *string     `json:"status"`
}

type TerminalEventData struct {
	Id string `json:"id"`
}

type TerminalEvenStatusData struct {
	Delivered bool `json:"delivered"`
}

type TerminalStatusData struct {
	Online    bool `json:"online"`
	Available bool `json:"available"`
}

type DedicatedAccountBank struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Slug string `json:"slug"`
}

type DedicatedAccountAssignment struct {
	Integration  int       `json:"integration"`
	AssigneeId   int       `json:"assignee_id"`
	AssigneeType string    `json:"assignee_type"`
	Expired      bool      `json:"expired"`
	AccountType  string    `json:"account_type"`
	AssignedAt   time.Time `json:"assigned_at"`
}

type DedicatedAccount struct {
	Bank          DedicatedAccountBank       `json:"bank"`
	AccountName   string                     `json:"account_name"`
	AccountNumber string                     `json:"account_number"`
	Assigned      bool                       `json:"assigned"`
	Currency      enum.Currency              `json:"currency"`
	Metadata      map[string]any             `json:"metadata"`
	Active        bool                       `json:"active"`
	Id            int                        `json:"id"`
	CreatedAt     time.Time                  `json:"created_at"`
	UpdatedAt     time.Time                  `json:"updated_at"`
	Assignment    DedicatedAccountAssignment `json:"assignment"`
	SplitConfig   map[string]any             `json:"split_config"`
}

type DedicatedAccountProvider struct {
	ProviderSlug string `json:"provider_slug"`
	BankId       int    `json:"bank_id"`
	BankName     string `json:"bank_name"`
	Id           int    `json:"id"`
}

type Settlement struct {
	Id              int           `json:"id"`
	Domain          enum.Domain   `json:"domain"`
	Status          string        `json:"status"`
	Currency        enum.Currency `json:"currency"`
	Integration     int           `json:"integration"`
	TotalAmount     int           `json:"total_amount"`
	EffectiveAmount int           `json:"effective_amount"`
	TotalFees       int           `json:"total_fees"`
	TotalProcessed  int           `json:"total_processed"`
	Deductions      any           `json:"deductions"`
	SettlmentDate   time.Time     `json:"settlement_date"`
	SettledBy       any           `json:"settled_by"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

type TransferRecipientDetail struct {
	AuthorizationCode *string `json:"authorization_code"`
	AccountName       *string `json:"account_name"`
	BankCode          string  `json:"bank_code"`
	BankName          string  `json:"bank_name"`
}

type TransferRecipient struct {
	Active           bool                    `json:"active"`
	CreatedAt        time.Time               `json:"created_at"`
	Currency         enum.Currency           `json:"currency"`
	Description      *string                 `json:"description"`
	Domain           enum.Domain             `json:"domain"`
	Email            *string                 `json:"email"`
	Id               int                     `json:"id"`
	Integration      int                     `json:"integration"`
	Metadata         map[string]any          `json:"metadata"`
	Name             string                  `json:"name"`
	RecipientCode    string                  `json:"recipient_code"`
	Type             string                  `json:"type"`
	UpdatedAt        time.Time               `json:"updated_at"`
	IsDeleted        bool                    `json:"is_deleted"`
	RecipientAccount *string                 `json:"recipient_account"`
	InstitutionCode  *string                 `json:"institution_code"`
	Details          TransferRecipientDetail `json:"details"`
}

type TransferRecipientBulkCreateData struct {
	Success []TransferRecipient `json:"success"`
	Errors  []any               `json:"errors"`
}

type TransferSession struct {
	Provider any `json:"provider"`
	Id       any `json:"id"`
}

type Tranfer struct {
	Integration     int             `json:"integration"`
	Domain          enum.Domain     `json:"domain"`
	Amount          int             `json:"amount"`
	Currency        enum.Currency   `json:"currency"`
	Source          string          `json:"source"`
	SourceDetails   any             `json:"source_details"`
	Failures        any             `json:"failures"`
	TitanCode       any             `json:"titan_code"`
	TransferredAt   *time.Time      `json:"transferred_at"`
	Reference       *string         `json:"reference"`
	Request         *int            `json:"request"`
	Reason          string          `json:"reason"`
	Recipient       any             `json:"recipient"`
	Status          string          `json:"status"`
	TransferCode    string          `json:"transfer_code"`
	Id              int             `json:"id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"update_at"`
	Session         TransferSession `json:"session"`
	FeeCharged      *int            `json:"fee_charged"`
	FeesBreakdown   int             `json:"fees_breakdown"`
	GatewayResponse any             `json:"gateway_response"`
}

type BulkTransferItem struct {
	Reference    string        `json:"reference"`
	Recipient    string        `json:"recipient"`
	Amount       int           `json:"amount"`
	TransferCode string        `json:"transfer_code"`
	Currency     enum.Currency `json:"currency"`
	Status       string        `json:"status"`
}

type BalanceLedgerItem struct {
	Integration      int           `json:"integration"`
	Domain           enum.Domain   `json:"domain"`
	Balance          int           `json:"balance"`
	Currency         enum.Currency `json:"currency"`
	Difference       int           `json:"difference"`
	Reason           string        `json:"reason"`
	ModelResponsible string        `json:"model_responsible"`
	ModelRow         int           `json:"model_row"`
	Id               int           `json:"id"`
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}

type DisputeHistory struct {
	Status    enum.DisputeStatus `json:"status"`
	By        string             `json:"by"`
	CreatedAt time.Time          `json:"created_at"`
}

type DisputeMessage struct {
	Sender    string    `json:"sender"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type Dispute struct {
	Id                   int                `json:"id"`
	RefundAmount         *int               `json:"refund_amount"`
	Currency             *enum.Currency     `json:"currency"`
	Status               enum.DisputeStatus `json:"status"`
	Resolution           any                `json:"resolution"`
	Domain               enum.Domain        `json:"domain"`
	Transaction          Transaction        `json:"transaction"`
	TransactionReference *string            `json:"transaction_reference"`
	Category             any                `json:"category"`
	Customer             Customer           `json:"customer"`
	Bin                  *string            `json:"bin"`
	Last4                *string            `json:"last4"`
	DueAt                *time.Time         `json:"due_at"`
	ResolvedAt           *time.Time         `json:"resolved_at"`
	Evidence             any                `json:"evidence"`
	Attachments          any                `json:"attachments"`
	Note                 any                `json:"note"`
	History              []DisputeHistory   `json:"history"`
	Messages             []DisputeMessage   `json:"messages"`
	CreatedAt            time.Time          `json:"create_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
}

type DisputeEvidence struct {
	CustomerEmail   string    `json:"customer_email"`
	CustomerName    string    `json:"customer_name"`
	CustomerPhone   string    `json:"customer_phone"`
	ServiceDetails  string    `json:"service_details"`
	DeliveryAddress string    `json:"delivery_address"`
	Dispute         int       `json:"dispute"`
	Id              int       `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type DisputeUploadInfo struct {
	SignedUrl string `json:"signed_url"`
	FileName  string `json:"file_name"`
}

type DisputeExportInfo struct {
	Path      string    `json:"path"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Refund struct {
	Integration          int           `json:""`
	Transaction          any           `json:""`
	Dispute              any           `json:""`
	Settlement           any           `json:""`
	Id                   int           `json:""`
	Domain               enum.Domain   `json:""`
	Currency             enum.Currency `json:""`
	Amount               int           `json:""`
	Status               string        `json:""`
	RefundedAt           *time.Time    `json:""`
	RefundedBy           string        `json:""`
	CustomerNote         string        `json:""`
	MerchantNote         string        `json:""`
	DeductedAmount       int           `json:""`
	FullyDeducted        any           `json:""`
	CreatedAt            any           `json:""`
	BankReference        any           `json:""`
	TransactionReference *string       `json:""`
	Reason               *string       `json:""`
	Customer             *Customer     `json:""`
	RefundType           *string       `json:""`
	TransactionAmount    *int          `json:""`
	InitiatedBy          *string       `json:""`
	RefundChannel        *string       `json:""`
	SessionId            any           `json:""`
	CollectAccountNumber *bool         `json:""`
}

type CardBin struct {
	Bin          string       `json:"bin"`
	Brand        string       `json:"brand"`
	SubBrand     string       `json:"sub_brand"`
	CountryCode  enum.Country `json:"country_code"`
	CountryName  string       `json:"country_name"`
	CardType     string       `json:"card_type"`
	Bank         string       `json:"bank"`
	LinkedBankId int          `json:"linked_bank_id"`
}

type Bank struct {
	Name             string        `json:"name"`
	Slug             string        `json:"slug"`
	Code             string        `json:"code"`
	Longcode         string        `json:"longcode"`
	Gateway          *string       `json:"gateway"`
	PayWithBank      bool          `json:"pay_with_bank"`
	SupportsTransfer bool          `json:"supports_transfer"`
	Active           bool          `json:"active"`
	IsDeleted        bool          `json:"is_deleted"`
	Country          string        `json:"country"`
	Currency         enum.Currency `json:"currency"`
	Type             string        `json:"type"`
	Id               int           `json:"id"`
	CreatedAt        *time.Time    `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}

type BankAccountInfo struct {
	AccountNumber string `json:"account_number"`
	AcountName    string `json:"account_name"`
}

type AccountVerificationInfo struct {
	Verified            bool   `json:"verified"`
	VerificationMessage string `json:"verification_message"`
}

type PaystackSupportedCountry struct {
	Id                           int                                                                            `json:"id"`
	ActiveForDashboardOnboarding bool                                                                           `json:"active_for_dashboard_onboarding"`
	Name                         string                                                                         `json:"name"`
	IsoCode                      string                                                                         `json:"iso_code"`
	DefaultCurrencyCode          enum.Currency                                                                  `json:"default_currency_code"`
	IntegrationDefaults          map[string]any                                                                 `json:"integration_defaults"`
	CallingCode                  string                                                                         `json:"calling_code"`
	PilotMode                    bool                                                                           `json:"pilot_mode"`
	Relationships                map[enum.SupportedCountryRelationshipType]SupportedCountryCurrencyRelationship `json:"relationships"`
	CanGoLiveAutomatically       *bool                                                                          `json:"can_go_live_automatically"`
}

type SupportedCountryRelationship[T, D any] struct {
	Type T   `json:"type"`
	Data []D `json:"data"`
}

type SupportedCountryCurrencyRelationship struct {
	Type                enum.SupportedCountryRelationshipType      `json:"type"`
	Data                []string                                   `json:"data"`
	SupportedCurrencies map[enum.Currency]SupportedCountryCurrency `json:"supported_currencies"`
	IntegrationType     map[string]any                             `json:"integration_type"`
	PaymentMethod       map[string]any                             `json:"payment_method"`
}

type SupportedCountryCurrencyMobileMoney struct {
	BankType                    string               `json:"bank_type"`
	PhoneNumberLabel            string               `json:"phone_number_label"`
	AccountNumberPattern        AccountNumberPattern `json:"account_number_pattern"`
	Placeholder                 *string              `json:"placeholder"`
	AccountVerificationRequired *bool                `json:"account_verification_requried"`
}

type SupportedCountryCurrencyEft struct {
	AccountNumberPattern AccountNumberPattern `json:"account_number_pattern"`
	Placeholder          string               `json:"placeholder"`
}

type SupportedCountryCurrency struct {
	Bank                SupportedCountryBank                 `json:"bank"`
	MobileMoney         *SupportedCountryCurrencyMobileMoney `json:"mobile_money"`
	MobileMoneyBusiness *SupportedCountryCurrencyMobileMoney `json:"mobile_money_business"`
	Eft                 *SupportedCountryCurrencyEft         `json:"eft"`
}

type SupportedCountryBank struct {
	BankType                    string               `json:"bank_type"`
	RequiredFields              *[]string            `json:"required_fields"`
	BranchCode                  any                  `json:"branch_code"`
	BranchCodeType              string               `json:"branch_code_type"`
	AccountName                 any                  `json:"account_name"`
	AccountVerificationRequired bool                 `json:"account_verification_requried"`
	AccountNumberLabel          string               `json:"account_number_label"`
	AccountNumberPattern        AccountNumberPattern `json:"account_number_pattern"`
	Documents                   *[]string            `json:"documents"`
	Notices                     *[]string            `json:"notices"`
	ShowAccountNumberTooltip    *bool                `json:"show_account_number_tooltip"`
}

type AccountNumberPattern struct {
	ExactMatch bool   `json:"exact_match"`
	Pattern    string `json:"pattern"`
}
