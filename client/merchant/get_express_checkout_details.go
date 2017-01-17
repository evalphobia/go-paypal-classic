package merchant

import "github.com/evalphobia/go-paypal-classic/client"

// GetExpressCheckoutDetails is struct for GetExpressCheckoutDetails API
// see: https://developer.paypal.com/docs/classic/api/merchant/GetExpressCheckoutDetails_API_Operation_NVP/
type GetExpressCheckoutDetails struct {
	client.BaseRequest `url:",squash"`

	Token string `url:"TOKEN"`
}

// Do executes GetExpressCheckoutDetails operation
func (svc *GetExpressCheckoutDetails) Do(cli client.Client) (*GetExpressCheckoutDetailsResponse, error) {
	const method = "GetExpressCheckoutDetails"
	svc.BaseRequest.Method = method

	result := &GetExpressCheckoutDetailsResponse{}
	err := cli.Call(svc, result)
	return result, err
}

// GetExpressCheckoutDetailsResponse is struct for response of GetExpressCheckoutDetails API
type GetExpressCheckoutDetailsResponse struct {
	client.BaseResponse `url:",squash"`

	// success
	Token              string `url:"TOKEN"`
	RecurringAgreement string `url:"BILLINGAGREEMENTACCEPTEDSTATUS"`
	CheckoutStatus     string `url:"CHECKOUTSTATUS"`
	Email              string `url:"EMAIL"`
	PayerID            string `url:"PAYERID"`
	PayerStatus        string `url:"PAYERSTATUS"`
	FirstName          string `url:"FIRSTNAME"`
	LastName           string `url:"LASTNAME"`
	CountryCode        string `url:"COUNTRYCODE"`
	CurrencyCode       string `url:"CURRENCYCODE"`

	Amount                 string `url:"AMT"`
	ItemAmount             string `url:"ITEMAMT"`
	ShippingAmount         string `url:"SHIPPINGAMT"`
	HandlingAmount         string `url:"HANDLINGAMT"`
	TaxAmount              string `url:"TAXAMT"`
	InsuranceAmount        string `url:"INSURANCEAMT"`
	ShippingDiscountAmount string `url:"SHIPDISCAMT"`

	IsFinancing string `url:"PAYMENTINFO_0_ISFINANCING"`

	// failure
	PaymentErrorCode    string `url:"PAYMENTREQUESTINFO_0_ERRORCODE"`
	PaymentShortMessage string `url:"PAYMENTREQUESTINFO_0_SHORTMESSAGE"`
	PaymentLongMessage  string `url:"PAYMENTREQUESTINFO_0_LONGMESSAGE"`
	PaymentSeverityCode string `url:"PAYMENTREQUESTINFO_0_SEVERITYCODE"`
}

// IsPayerVerified checks the payer status is verified or not
func (r *GetExpressCheckoutDetailsResponse) IsPayerVerified() bool {
	return r.PayerStatus == payerStatusVerified
}

// IsSuccess checks the request is success or not
func (r *GetExpressCheckoutDetailsResponse) IsSuccess() bool {
	return r.IsRequestSuccess() && r.IsOperationSuccess()
}

// IsOperationSuccess checks the request is success or not
func (r *GetExpressCheckoutDetailsResponse) IsOperationSuccess() bool {
	return r.IsPayerVerified()
}

// Error returns error text
func (r *GetExpressCheckoutDetailsResponse) Error() string {
	var s []string
	if r.PaymentShortMessage != "" {
		s = append(s, r.PaymentShortMessage+" "+r.PaymentLongMessage)
	}
	return r.BaseResponse.Errors(s)
}
