package merchant

// GetTransactionDetails is struct for GetTransactionDetails API
// see: https://developer.paypal.com/docs/classic/api/merchant/GetTransactionDetails_API_Operation_NVP/
type GetTransactionDetails struct {
	Merchant    `url:",squash"`
	BaseRequest `url:",squash"`

	TransactionID string `url:"TRANSACTIONID"`
}

// SetMerchant sets Merchant
func (svc *GetTransactionDetails) SetMerchant(m Merchant) {
	svc.Merchant = m
}

// Do executes GetTransactionDetails operation
func (svc *GetTransactionDetails) Do(m Merchant) (*GetTransactionDetailsResponse, error) {
	const method = "GetTransactionDetails"
	svc.BaseRequest.Method = method

	result := &GetTransactionDetailsResponse{}
	err := m.call(svc, result)
	return result, err
}

// GetTransactionDetailsResponse is struct for response of GetTransactionDetails API
type GetTransactionDetailsResponse struct {
	BaseResponse `url:",squash"`

	// success
	TransactionID   string `url:"TRANSACTIONID"`
	TransactionType string `url:"TRANSACTIONTYPE"`
	PaymentType     string `url:"PAYMENTTYPE"`
	Subject         string `url:"SUBJECT"`

	CurrencyCode   string `url:"CURRENCYCODE"`
	Amount         string `url:"AMT"`
	ShippingAmount string `url:"SHIPPINGAMT"`
	TaxAmount      string `url:"TAXAMT"`
	FeeAmount      string `url:"FEEAMT"`
	OrderTime      string `url:"ORDERTIME"`

	PaymentStatus             string `url:"PAYMENTSTATUS"`
	PendingReason             string `url:"PENDINGREASON"`
	ReasonCode                string `url:"REASONCODE"`
	ProtectionEligibility     string `url:"PROTECTIONELIGIBILITY"`
	ProtectionEligibilityType string `url:"PROTECTIONELIGIBILITYTYPE"`

	PayerID       string `url:"PAYERID"`
	PayerStatus   string `url:"PAYERSTATUS"`
	FirstName     string `url:"FIRSTNAME"`
	LastName      string `url:"LASTNAME"`
	Emain         string `url:"EMAIL"`
	ReceiverID    string `url:"RECEIVERID"`
	ReceiverEmail string `url:"RECEIVEREMAIL"`
	CountryCode   string `url:"COUNTRYCODE"`
	AddressOwner  string `url:"ADDRESSOWNER"`
	AddressStatus string `url:"ADDRESSSTATUS"`

	SalesTax         string `url:"SALESTAX"`
	ShipAmount       string `url:"SHIPAMOUNT"`
	ShipHandleAmount string `url:"SHIPHANDLEAMOUNT"`

	ShippingResponse `url:",squash"`
}

// IsSuccess checks the request is success or not
func (r *GetTransactionDetailsResponse) IsSuccess() bool {
	return r.IsRequestSuccess() && r.IsOperationSuccess()
}

// IsRequestSuccess checks the request is success or not
func (r *GetTransactionDetailsResponse) IsRequestSuccess() bool {
	return r.ACK == ackSuccess
}

// IsOperationSuccess checks the request is success or not
func (r *GetTransactionDetailsResponse) IsOperationSuccess() bool {
	return r.PaymentStatus == paymentStatusComleted
}

// Error returns error text
func (r *GetTransactionDetailsResponse) Error() string {
	return r.BaseResponse.Error()
}
