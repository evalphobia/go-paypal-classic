package merchant

import (
	"strings"

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

	ackSuccess = "Success"
	ackFailure = "Failure"
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

// BaseRequest is base struct for api request
type BaseRequest struct {
	Method string `url:"METHOD"`
	Action string `url:"PAYMENTREQUEST_0_PAYMENTACTION,omitempty"`
}

type BaseResponse struct {
	Timestamp     string `url:"TIMESTAMP"`
	ACK           string `url:"ACK"`
	Version       string `url:"VERSION"`
	Build         string `url:"BUILD"`
	CorrelationID string `url:"CORRELATIONID"`

	// error0
	ErrorCode    string `url:"L_ERRORCODE0"`
	ShortMessage string `url:"L_SHORTMESSAGE0"`
	LongMessage  string `url:"L_LONGMESSAGE0"`
	SeverityCode string `url:"L_SEVERITYCODE0"`

	// error1
	ErrorCode1    string `url:"L_ERRORCODE1"`
	ShortMessage1 string `url:"L_SHORTMESSAGE1"`
	LongMessage1  string `url:"L_LONGMESSAGE1"`
	SeverityCode1 string `url:"L_SEVERITYCODE1"`
}

func (r BaseResponse) baseError() []string {
	var errs []string
	if r.ShortMessage != "" {
		errs = append(errs, r.ShortMessage+" "+r.LongMessage)
	}
	if r.ShortMessage1 != "" {
		errs = append(errs, r.ShortMessage1+" "+r.LongMessage1)
	}
	return errs
}

func (r BaseResponse) Error() string {
	return parseErrors(r.baseError())
}

func (r BaseResponse) Errors(s []string) string {
	errs := r.baseError()
	errs = append(errs, s...)
	return parseErrors(errs)
}

func parseErrors(errs []string) string {
	return strings.Join(errs, ",")
}
