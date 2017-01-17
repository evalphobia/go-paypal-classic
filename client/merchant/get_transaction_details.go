package merchant

import "github.com/evalphobia/go-paypal-classic/client"

// GetTransactionDetails is struct for GetTransactionDetails API
// see: https://developer.paypal.com/docs/classic/api/merchant/GetTransactionDetails_API_Operation_NVP/
type GetTransactionDetails struct {
	client.BaseRequest `url:",squash"`

	TransactionID string `url:"TRANSACTIONID"`
}

// Do executes GetTransactionDetails operation
func (svc *GetTransactionDetails) Do(cli client.Client) (*GetTransactionDetailsResponse, error) {
	const method = "GetTransactionDetails"
	svc.BaseRequest.Method = method

	result := &GetTransactionDetailsResponse{}
	err := cli.Call(svc, result)
	return result, err
}

// GetTransactionDetailsResponse is struct for response of GetTransactionDetails API
type GetTransactionDetailsResponse struct {
	client.BaseResponse `url:",squash"`

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

	client.ShippingResponse `url:",squash"`
}

// IsSuccess checks the request is success or not
func (r *GetTransactionDetailsResponse) IsSuccess() bool {
	return r.IsRequestSuccess() && r.IsOperationSuccess()
}

// IsOperationSuccess checks the request is success or not
func (r *GetTransactionDetailsResponse) IsOperationSuccess() bool {
	return r.PaymentStatus == paymentStatusComleted
}

// Error returns error text
func (r *GetTransactionDetailsResponse) Error() string {
	return r.BaseResponse.Error()
}
