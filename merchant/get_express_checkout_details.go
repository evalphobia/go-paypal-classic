package merchant

// GetExpressCheckoutDetails is struct for GetExpressCheckoutDetails API
// see: https://developer.paypal.com/docs/classic/api/merchant/GetExpressCheckoutDetails_API_Operation_NVP/
type GetExpressCheckoutDetails struct {
	Merchant    `url:",squash"`
	BaseRequest `url:",squash"`

	Token string `url:"TOKEN"`
}

// SetMerchant sets Merchant
func (svc *GetExpressCheckoutDetails) SetMerchant(m Merchant) {
	svc.Merchant = m
}

// Do executes GetExpressCheckoutDetails operation
func (svc *GetExpressCheckoutDetails) Do(m Merchant) (*GetExpressCheckoutDetailsResponse, error) {
	const method = "GetExpressCheckoutDetails"
	svc.BaseRequest.Method = method

	result := &GetExpressCheckoutDetailsResponse{}
	err := m.call(svc, result)
	return result, err
}

// GetExpressCheckoutDetailsResponse is struct for response of GetExpressCheckoutDetails API
type GetExpressCheckoutDetailsResponse struct {
	BaseResponse `url:",squash"`

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

// IsSuccess checks the request is success or not
func (r *GetExpressCheckoutDetailsResponse) IsSuccess() bool {
	return r.ACK == ackSuccess
}

// IsPayerVerified checks the payer status is verified or not
func (r *GetExpressCheckoutDetailsResponse) IsPayerVerified() bool {
	return r.PayerStatus == payerStatusVerified
}

// Error returns error text
func (r *GetExpressCheckoutDetailsResponse) Error() string {
	var s []string
	if r.PaymentShortMessage != "" {
		s = append(s, r.PaymentShortMessage+" "+r.PaymentLongMessage)
	}
	return r.BaseResponse.Errors(s)
}
