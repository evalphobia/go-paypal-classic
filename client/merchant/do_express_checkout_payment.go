package merchant

import "github.com/evalphobia/go-paypal-classic/client"

// DoExpressCheckoutPayment is struct for DoExpressCheckoutPayment API
// see: https://developer.paypal.com/docs/classic/api/merchant/DoExpressCheckoutPayment_API_Operation_NVP/
type DoExpressCheckoutPayment struct {
	client.BaseRequest `url:",squash"`

	Token   string `url:"TOKEN"`
	PayerID string `url:"PAYERID"`

	TotalAmount float64 `url:"PAYMENTREQUEST_0_AMT"`
	ItemAmount  float64 `url:"PAYMENTREQUEST_0_ITEMAMT"`
	TaxAmount   float64 `url:"PAYMENTREQUEST_0_TAXAMT"`
	Currency    string  `url:"PAYMENTREQUEST_0_CURRENCYCODE"`
}

// Do executes DoExpressCheckoutPayment operation
func (svc *DoExpressCheckoutPayment) Do(cli client.Client) (*DoExpressCheckoutPaymentResponse, error) {
	const method = "DoExpressCheckoutPayment"
	svc.BaseRequest.Method = method
	svc.BaseRequest.Action = paymentActionSale

	if svc.TotalAmount == 0 {
		svc.TotalAmount = svc.ItemAmount + svc.TaxAmount
	}

	result := &DoExpressCheckoutPaymentResponse{}
	err := cli.Call(svc, result)
	return result, err
}

// DoExpressCheckoutPaymentResponse is struct for response of DoExpressCheckoutPayment API
type DoExpressCheckoutPaymentResponse struct {
	client.BaseResponse `url:",squash"`

	// success
	Token string `url:"TOKEN"`

	PaymentACK      string `url:"PAYMENTINFO_0_ACK"`
	TransactionID   string `url:"PAYMENTINFO_0_TRANSACTIONID"`
	TransactionType string `url:"PAYMENTINFO_0_TRANSACTIONTYPE"`
	PaymentType     string `url:"PAYMENTINFO_0_PAYMENTTYPE"`
	OrderTime       string `url:"PAYMENTINFO_0_ORDERTIME"`
	Amount          string `url:"PAYMENTINFO_0_AMT"`
	FeeAmount       string `url:"PAYMENTINFO_0_FEEAMT"`
	TaxAmount       string `url:"PAYMENTINFO_0_TAXAMT"`
	CurrecyCode     string `url:"PAYMENTINFO_0_CURRENCYCODE"`
	PaymentStatus   string `url:"PAYMENTINFO_0_PAYMENTSTATUS"`

	PendingReason string `url:"PAYMENTINFO_0_PENDINGREASON"`
	ReasonCode    string `url:"PAYMENTINFO_0_REASONCODE"`

	ProtectionEligibility     string `url:"PAYMENTINFO_0_PROTECTIONELIGIBILITY"`
	ProtectionEligibilityType string `url:"PAYMENTINFO_0_PROTECTIONELIGIBILITYTYPE"`
	SecureMerchantAccountID   string `url:"PAYMENTINFO_0_SECUREMERCHANTACCOUNTID"`

	InsuranceOptionSelected string `url:"INSURANCEOPTIONSELECTED"`
	ShippingOptionIsDefault string `url:"SHIPPINGOPTIONISDEFAULT"`

	// failure
	PaymentErrorCode    string `url:"PAYMENTINFO_0_ERRORCODE"`
	PaymentShortMessage string `url:"PAYMENTINFO_0_SHORTMESSAGE"`
	PaymentLongMessage  string `url:"PAYMENTINFO_0_LONGMESSAGE"`
	PaymentSeverityCode string `url:"PAYMENTINFO_0_SEVERITYCODE"`
}

// IsSuccess checks the request is success or not
func (r *DoExpressCheckoutPaymentResponse) IsSuccess() bool {
	return r.IsRequestSuccess() && r.IsOperationSuccess()
}

// IsOperationSuccess checks the request is success or not
func (r *DoExpressCheckoutPaymentResponse) IsOperationSuccess() bool {
	return r.PaymentACK == ackSuccess
}

// Error returns error text
func (r *DoExpressCheckoutPaymentResponse) Error() string {
	var s []string
	if r.PaymentShortMessage != "" {
		s = append(s, r.PaymentShortMessage+" "+r.PaymentLongMessage)
	}
	return r.BaseResponse.Errors(s)
}
