package merchant

// SetExpressCheckout is struct for SetExpressCheckout API
// see: https://developer.paypal.com/docs/classic/api/merchant/SetExpressCheckout_API_Operation_NVP/
type SetExpressCheckout struct {
	Merchant    `url:",squash"`
	BaseRequest `url:",squash"`

	Action string `url:"PAYMENTREQUEST_0_PAYMENTACTION"`

	ReturnURL   string  `url:"RETURNURL"`
	CancelURL   string  `url:"CANCELURL"`
	TotalAmount float64 `url:"PAYMENTREQUEST_0_AMT"`
	Currency    string  `url:"PAYMENTREQUEST_0_CURRENCYCODE,omitempty"`

	ItemName   string  `url:"L_PAYMENTREQUEST_0_NAME0,omitempty"`
	ItemQty    int     `url:"L_PAYMENTREQUEST_0_QTY0,omitempty"`
	ItemAmount float64 `url:"L_PAYMENTREQUEST_0_AMT0,omitempty"`

	Category        string `url:"L_PAYMENTREQUEST_0_ITEMCATEGORY0,omitempty"`
	ConfirmShipping string `url:"REQCONFIRMSHIPPING,omitempty"`
	NoShipping      string `url:"NOSHIPPING,omitempty"`

	BillingType string `url:"L_BILLINGTYPE0"`
	Description string `url:"L_BILLINGAGREEMENTDESCRIPTION0"`

	Locale string `url:"LOCALECODE"`
}

// SetMerchant sets Merchant
func (svc *SetExpressCheckout) SetMerchant(m Merchant) {
	svc.Merchant = m
}

// SetAsRecurringPayment sets billing type as recurring payment
func (svc *SetExpressCheckout) SetAsRecurringPayment(desc string) *SetExpressCheckout {
	svc.BillingType = billingTypeRecurring
	svc.Description = desc
	return svc
}

// SetAsDigitalCategory sets item category as digital item
func (svc *SetExpressCheckout) SetAsDigitalCategory() *SetExpressCheckout {
	svc.Category = itemCategoryDigital
	svc.ConfirmShipping = "0"
	svc.NoShipping = "1"
	if svc.ItemQty == 0 {
		svc.ItemQty = 1
	}
	return svc
}

// Do executes SetExpressCheckout operation
func (svc *SetExpressCheckout) Do(m Merchant) (*SetExpressCheckoutResponse, error) {
	const method = "SetExpressCheckout"
	svc.BaseRequest.Method = method
	svc.BaseRequest.Action = paymentActionSale

	if svc.TotalAmount == 0 {
		svc.TotalAmount = svc.ItemAmount
	}

	result := &SetExpressCheckoutResponse{
		redirectURL: m.redirectBase(),
	}
	err := m.call(svc, result)
	return result, err
}

// SetExpressCheckoutResponse is struct for response of SetExpressCheckout API
type SetExpressCheckoutResponse struct {
	BaseResponse `url:",squash"`
	redirectURL  string `url:"-"`

	// success
	Token string `url:"TOKEN"`
}

// RedirectURL returns full URL of PayPal express checkout
func (r *SetExpressCheckoutResponse) RedirectURL() string {
	if r.Token == "" {
		return ""
	}
	return r.redirectURL + "?cmd=_express-checkout&token=" + r.Token
}

// IsSuccess checks the request is success or not
func (r *SetExpressCheckoutResponse) IsSuccess() bool {
	return r.IsRequestSuccess()
}

// IsRequestSuccess checks the request is success or not
func (r *SetExpressCheckoutResponse) IsRequestSuccess() bool {
	return r.ACK == ackSuccess
}

// IsOperationSuccess checks the request is success or not
func (r *SetExpressCheckoutResponse) IsOperationSuccess() bool {
	return r.IsRequestSuccess()
}

// Error returns error text
func (r *SetExpressCheckoutResponse) Error() string {
	return r.BaseResponse.Error()
}
