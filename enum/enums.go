package enum

type (
	TerminalEvent                    string
	TerminalEventAction              string
	Currency                         string
	Interval                         string
	Channel                          string
	Bearer                           string
	TransactionStatus                string
	Split                            string
	Country                          string
	RiskAction                       string
	Identification                   string
	RecipientType                    string
	Document                         string
	GenericStatus                    string
	Schedule                         string
	Reason                           string
	Gateway                          string
	AccountType                      string
	Resolution                       string
	BankType                         string
	DisputeStatus                    string
	Domain                           string
	BulkChargeStatus                 string
	SupportedCountryRelationshipType string
)

const (
	// TerminalEvent enum variants
	TerminalEventTransaction TerminalEvent = "transaction"
	TerminalEventInvoice     TerminalEvent = "invoice"

	// TerminalEventAction enum variants
	TerminalEventActionProcess TerminalEventAction = "process"
	TerminalEventActionView    TerminalEventAction = "view"
	TerminalEventActionPrint   TerminalEventAction = "print"

	// Currency enum variants
	CurrencyNgn Currency = "NGN"
	CurrencyGhs Currency = "GHS"
	CurrencyZar Currency = "ZAR"
	CurrencyUsd Currency = "USD"
	CurrencyKes Currency = "KES"
	CurrencyXof Currency = "XOF"
	CurrencyEgp Currency = "EGP"
	CurrencyRwf Currency = "RWF"

	// Interval enum variants
	IntervalHourly     Interval = "hourly"
	IntervalDaily      Interval = "daily"
	IntervalWeekly     Interval = "weekly"
	IntervalMonthly    Interval = "monthly"
	IntervalQuarterly  Interval = "quarterly"
	IntervalBiannually Interval = "biannually"
	IntervalAnnually   Interval = "annually"

	// Channel enum variants
	ChannelCard         Channel = "card"
	ChannelBank         Channel = "bank"
	ChannelUssd         Channel = "ussd"
	ChannelQr           Channel = "qr"
	ChannelMobileMoney  Channel = "mobile_money"
	ChannelBankTransfer Channel = "bank_transfer"

	// Bearer enum variants
	BearerAccount         Bearer = "account"
	BearerSubAccount      Bearer = "subaccount"
	BearerAllProportional Bearer = "all-proportional"
	BearerAll             Bearer = "all"

	// TransactionStatus enum variants
	TransactionStatusFailed    TransactionStatus = "failed"
	TransactionStatusSuccess   TransactionStatus = "success"
	TransactionStatusAbandoned TransactionStatus = "abandoned"

	// Split enum variants
	SplitPercentage Split = "percentage"
	SplitFlat       Split = "flat"

	// Country enum variants
	CountryNigeria     Country = "NG"
	CountryGhana       Country = "GH"
	CountrySouthAfrica Country = "ZA"
	CountryKenya       Country = "KE"
	CountryCoteDIvoire Country = "CI"
	CountryEgypt       Country = "EG"
	CountryRwanda      Country = "RW"

	// RiskAction enum variants
	RiskActionDefault   RiskAction = "default"
	RiskActionWhitelist RiskAction = "allow"
	RiskActionBlacklist RiskAction = "deny"

	// Identification enum variants
	IdentificationBvn         Identification = "bvn"
	IdentificationBankAccount Identification = "bank_account"

	// RecipientType enum variants
	RecipientTypeNuban       RecipientType = "nuban"
	RecipientTypeMobileMoney RecipientType = "mobile_money"
	RecipientTypeBasa        RecipientType = "basa"

	// Document enum variants
	DocumentIdentityNumber             Document = "identityNumber"
	DocumentPassportNumber             Document = "passportNumber"
	DocumentBusinessRegistrationNumber Document = "businessRegistrationNumber"

	// GenericStatus enum variants
	GenericStatusPending GenericStatus = "pending"
	GenericStatusSuccess GenericStatus = "success"
	GenericStatusFailed  GenericStatus = "failed"

	// Schedule enum variants
	ScheduleAuto    Schedule = "auto"
	ScheduleWeekly  Schedule = "weekly"
	ScheduleMonthly Schedule = "monthly"
	ScheduleManual  Schedule = "manual"

	// Reason enum variants
	ReasonResendOtp  Reason = "resend_otp"
	ReasonTransfer   Reason = "transfer"
	ReasonDisableOtp Reason = "disable_otp"

	// Gateway enum variants
	GatewayEmandate           Gateway = "emandate"
	GatewayDigitalBankMandate Gateway = "digitalbankmandate"

	// AccountType enum variants
	AccountTypePersonal AccountType = "personal"
	AccountTypeBusiness AccountType = "business"

	// Resolution enum variants
	ResolutionMerchantAccepted Resolution = "merchant-accepted"
	ResolutionDeclined         Resolution = "declined"

	// BankType enum variants
	BankTypeGhipps      BankType = "ghipps"
	BankTypeMobileMoney BankType = "mobile_money"

	// DisputeStatus enum variants
	DisputeStatusPending                  DisputeStatus = "pending"
	DisputeStatusResolved                 DisputeStatus = "resolved"
	DisputeStatusAwaitingBankFeedBack     DisputeStatus = "awaiting-bank-feedback"
	DisputeStatusAwaitingMerchantFeedback DisputeStatus = "awaiting-merchant-feedback"
	DisputeStatusArchived                 DisputeStatus = "archived"

	// Domain enum variants
	DomainLive Domain = "live"
	DomainTest Domain = "test"

	// BulkChargeStatus enum variants
	BulkChargeStatusActive   BulkChargeStatus = "active"
	BulkChargeStatusPaused   BulkChargeStatus = "paused"
	BulkChargeStatusComplete BulkChargeStatus = "complete"

	// SupportedCountryRelationshipType enum variants
	SupportedCountryRelationshipTypeCurrency           SupportedCountryRelationshipType = "currency"
	SupportedCountryRelationshipTypeIntegrationFeature SupportedCountryRelationshipType = "integration_feature"
	SupportedCountryRelationshipTypeIntegrationType    SupportedCountryRelationshipType = "integration_type"
	SupportedCountryRelationshipTypePaymentMethod      SupportedCountryRelationshipType = "payment_method"
)
