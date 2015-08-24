package merchant

// GetRecurringPaymentsProfileDetails is struct for GetRecurringPaymentsProfileDetails API
// see: https://developer.paypal.com/docs/classic/api/merchant/GetRecurringPaymentsProfileDetails_API_Operation_NVP/
type GetRecurringPaymentsProfileDetails struct {
	Merchant    `url:",squash"`
	BaseRequest `url:",squash"`

	ProfileID string `url:"PROFILEID"`
}

// SetMerchant sets Merchant
func (svc *GetRecurringPaymentsProfileDetails) SetMerchant(m Merchant) {
	svc.Merchant = m
}

// Do executes GetRecurringPaymentsProfileDetails operation
func (svc *GetRecurringPaymentsProfileDetails) Do(m Merchant) (*GetRecurringPaymentsProfileDetailsResponse, error) {
	const method = "GetRecurringPaymentsProfileDetails"
	svc.BaseRequest.Method = method

	result := &GetRecurringPaymentsProfileDetailsResponse{}
	err := m.call(svc, result)
	return result, err
}

// GetRecurringPaymentsProfileDetailsResponse is struct for response of GetRecurringPaymentsProfileDetails API
type GetRecurringPaymentsProfileDetailsResponse struct {
	BaseResponse `url:",squash"`

	// success
	ProfileID             string `url:"PROFILEID"`
	Status                string `url:"STATUS"`
	AutoBilloutAmount     string `url:"AUTOBILLOUTAMT"`
	Description           string `url:"DESC"`
	MaxFailedPayments     string `url:"MAXFAILEDPAYMENTS"`
	SubscriberName        string `url:"SUBSCRIBERNAME"`
	ProfileStartDate      string `url:"PROFILESTARTDATE"`
	NextBillingDate       string `url:"NEXTBILLINGDATE"`
	CompletedCyclesNumber string `url:"NUMCYCLESCOMPLETED"`
	RemainingCyclesNumber string `url:"NUMCYCLESREMAINING"`
	OutstandingBalance    string `url:"OUTSTANDINGBALANCE"`
	FailedPaymentCount    string `url:"FAILEDPAYMENTCOUNT"`

	TrialAmountPaid         string `url:"TRIALAMTPAID"`
	RegularAmountPaid       string `url:"REGULARAMTPAID"`
	AggregateAmount         string `url:"AGGREGATEAMT"`
	AggregateOptionalAmount string `url:"AGGREGATEOPTIONALAMT"`
	FinalPaymentDueDate     string `url:"FINALPAYMENTDUEDATE"`

	BillingPeriod      string `url:"BILLINGPERIOD"`
	BilligFrequency    string `url:"BILLINGFREQUENCY"`
	TotalBillingCycles string `url:"TOTALBILLINGCYCLES"`
	CurrencyCode       string `url:"CURRENCYCODE"`
	Amount             string `url:"AMT"`
	ShippingAmount     string `url:"SHIPPINGAMT"`
	TaxAmount          string `url:"TAXAMT"`

	RegularBillingPeriod    string `url:"REGULARBILLINGPERIOD"`
	RegularBillingFrequency string `url:"REGULARBILLINGFREQUENCY"`
	RegularBillingCycles    string `url:"REGULARTOTALBILLINGCYCLES"`
	RegularCurrencyCode     string `url:"REGULARCURRENCYCODE"`
	RegularAmount           string `url:"REGULARAMT"`
	RegularShippingAmount   string `url:"REGULARSHIPPINGAMT"`
	RegularTaxAmount        string `url:"REGULARTAXAMT"`

	ShippingResponse `url:",squash"`
}

// IsActive checks the recurring payment is still active
func (r *GetRecurringPaymentsProfileDetailsResponse) IsActive() bool {
	return r.Status == statusActive
}

// IsSuccess checks the request is success or not
func (r *GetRecurringPaymentsProfileDetailsResponse) IsSuccess() bool {
	return r.IsRequestSuccess() && r.IsOperationSuccess()
}

// IsRequestSuccess checks the request is success or not
func (r *GetRecurringPaymentsProfileDetailsResponse) IsRequestSuccess() bool {
	return r.ACK == ackSuccess
}

// IsOperationSuccess checks the request is success or not
func (r *GetRecurringPaymentsProfileDetailsResponse) IsOperationSuccess() bool {
	return r.IsActive()
}

// Error returns error text
func (r *GetRecurringPaymentsProfileDetailsResponse) Error() string {
	return r.BaseResponse.Error()
}
