package merchant

import (
	"time"

	"github.com/evalphobia/go-paypal-classic/client"
)

// CreateRecurringPaymentsProfile is struct for CreateRecurringPaymentsProfile API
// see: https://developer.paypal.com/docs/classic/api/merchant/CreateRecurringPaymentsProfile_API_Operation_NVP/
type CreateRecurringPaymentsProfile struct {
	client.BaseRequest `url:",squash"`

	BillingPeriod    string `url:"BILLINGPERIOD"`
	BillingFrequency int    `url:"BILLINGFREQUENCY"`
	BillingStartDate string `url:"PROFILESTARTDATE"`

	SubscriberName    string  `url:"SUBSCRIBERNAME"`
	Description       string  `url:"DESC"`
	Amount            float64 `url:"AMT"`
	Currency          string  `url:"CURRENCYCODE"`
	MaxFailedPayments int     `url:"MAXFAILEDPAYMENTS,omitempty"`

	Token   string `url:"TOKEN"`
	PayerID string `url:"PAYERID"`
}

// SetPeriodAsDay sets billing period as `day`
func (svc *CreateRecurringPaymentsProfile) SetPeriodAsDay(i int) {
	svc.BillingPeriod = "Day"
	svc.BillingFrequency = i
}

// SetPeriodAsWeek sets billing period as `week`
func (svc *CreateRecurringPaymentsProfile) SetPeriodAsWeek(i int) {
	svc.BillingPeriod = "Week"
	svc.BillingFrequency = i
}

// SetPeriodAsMonth sets billing period as `month`
func (svc *CreateRecurringPaymentsProfile) SetPeriodAsMonth(i int) {
	svc.BillingPeriod = "Month"
	svc.BillingFrequency = i
}

// SetPeriodAsYear sets billing period as `year`
func (svc *CreateRecurringPaymentsProfile) SetPeriodAsYear() {
	svc.BillingPeriod = "Year"
	svc.BillingFrequency = 1
}

// SetBillingStartDateFromNow sets billing start date respecting to billing period from today
func (svc *CreateRecurringPaymentsProfile) SetBillingStartDateFromNow() {
	dt := time.Now()
	switch svc.BillingPeriod {
	case "Day":
		svc.BillingStartDate = dt.AddDate(0, 0, svc.BillingFrequency).Format(time.RFC3339)
	case "Week":
		svc.BillingStartDate = dt.AddDate(0, 0, svc.BillingFrequency*7).Format(time.RFC3339)
	case "Month":
		// TODO: consider Jan 30th & Jan 31st for Feb 28th
		svc.BillingStartDate = dt.AddDate(0, svc.BillingFrequency, 0).Format(time.RFC3339)
	case "Year":
		svc.BillingStartDate = dt.AddDate(1, 0, 0).Format(time.RFC3339)
	}
}

// Do executes CreateRecurringPaymentsProfile operation
func (svc *CreateRecurringPaymentsProfile) Do(cli client.Client) (*CreateRecurringPaymentsProfileResponse, error) {
	const method = "CreateRecurringPaymentsProfile"
	svc.BaseRequest.Method = method
	svc.BaseRequest.Action = paymentActionSale

	result := &CreateRecurringPaymentsProfileResponse{}
	err := cli.Call(svc, result)
	return result, err
}

// CreateRecurringPaymentsProfileResponse is struct for response of CreateRecurringPaymentsProfile API
type CreateRecurringPaymentsProfileResponse struct {
	client.BaseResponse `url:",squash"`

	// success
	Token         string `url:"TOKEN"`
	ProfileID     string `url:"PROFILEID"`
	ProfileStatus string `url:"PROFILESTATUS"`
}

// IsSuccess checks the request is success or not
func (r *CreateRecurringPaymentsProfileResponse) IsSuccess() bool {
	return r.IsRequestSuccess() && r.IsOperationSuccess()
}

// IsOperationSuccess checks the request is success or not
func (r *CreateRecurringPaymentsProfileResponse) IsOperationSuccess() bool {
	return r.ProfileStatus == profileActive
}

// Error returns error text
func (r *CreateRecurringPaymentsProfileResponse) Error() string {
	return r.BaseResponse.Error()
}
