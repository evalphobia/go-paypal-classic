package merchant

import (
	"github.com/evalphobia/go-paypal-classic/config"
	"github.com/evalphobia/go-paypal-classic/request"
)

const (
	endpointSandbox    = "https://api-3t.sandbox.paypal.com/nvp"
	endpointProduction = "https://api-3t.paypal.com/nvp"

	redirectSandbox    = "https://www.sandbox.paypal.com/webscr"
	redirectProduction = "https://www.paypal.com/webscr"

	apiVersion = 124.0

	paymentActionSale    = "sale"
	payerStatusVerified  = "verified"
	itemCategoryDigital  = "Digital"
	billingTypeRecurring = "RecurringPayments"

	ackSuccess            = "Success"
	ackFailure            = "Failure"
	statusActive          = "Active"
	profileActive         = "ActiveProfile"
	paymentStatusComleted = "Completed"
)

// Merchant is base struct for PayPal Classic API
type Merchant struct {
	Config *config.Config `url:",squash"`
}

// New creates Merchant with given config
func New(conf *config.Config) Merchant {
	return Merchant{
		Config: conf,
	}
}

// NewDefault creates Merchant with default config
func NewDefault() Merchant {
	return New(config.DefaultConfig)
}

// MerchantRequest is interface of each API
type MerchantRequest interface {
	SetMerchant(Merchant)
}

// call sends HTTP request to PayPal api
func (m Merchant) call(svc MerchantRequest, result interface{}) error {
	svc.SetMerchant(m)
	if m.Config.IsProduction() {
		return request.CallGET(endpointProduction, svc, result)
	}
	return request.CallGET(endpointSandbox, svc, result)
}

// redirectBase returns base url of PayPal ExpressCheckout
func (m Merchant) redirectBase() string {
	if m.Config.IsProduction() {
		return redirectProduction
	}
	return redirectSandbox
}
