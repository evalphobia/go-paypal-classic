package transaction

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/evalphobia/go-paypal-classic/client"
)

// TransactionSearch is struct for TransactionSearch API
// see: https://developer.paypal.com/docs/classic/api/merchant/TransactionSearch_API_Operation_NVP/
type TransactionSearch struct {
	client.BaseRequest `url:",squash"`
	ParamStartDate     string `url:"STARTDATE"`
	ParamEndDate       string `url:"ENDDATE,omitempty"`

	// Required
	StartDate time.Time `url:"-"`

	// Optional
	EndDate time.Time `url:"-"`

	ProfileID         string `url:"PROFILEID,omitempty"`
	TransactionID     string `url:"TRANSACTIONID,omitempty"`
	RceiptID          string `url:"RECEIPTID,omitempty"`
	InvNum            string `url:"INVNUM,omitempty"`
	AuctionItemNumber string `url:"AUCTIONITEMNUMBER,omitempty"`

	Email    string `url:"EMAIL,omitempty"`
	Receiver string `url:"RECEIVER,omitempty"`

	Status           string `url:"STATUS,omitempty"`
	TransactionClass string `url:"TRANSACTIONCLASS,omitempty"`
}

// Do executes TransactionSearch operation
func (svc *TransactionSearch) Do(cli client.Client) (*TransactionSearchResponse, error) {
	const method = "TransactionSearch"
	svc.BaseRequest.Method = method

	if !svc.StartDate.IsZero() {
		svc.ParamStartDate = timeToString(svc.StartDate)
	}
	if !svc.EndDate.IsZero() {
		svc.ParamEndDate = timeToString(svc.EndDate)
	}

	result := &TransactionSearchResponse{}
	err := cli.Call(svc, result)
	return result, err
}

// TransactionSearchResponse is struct for response of GetRecurringPaymentsProfileDetails API
type TransactionSearchResponse struct {
	client.BaseResponse `url:",squash"`

	Items []TransactionSearchItem
}

// IsSuccess checks the request is success or not
func (r *TransactionSearchResponse) IsSuccess() bool {
	return r.IsRequestSuccess()
}

// Error returns error text
func (r *TransactionSearchResponse) Error() string {
	return r.BaseResponse.Error()
}

// Unmarshal assigns results from map.
func (r *TransactionSearchResponse) Unmarshal(mapData map[string]interface{}) error {
	if _, ok := mapData["L_TRANSACTIONID0"]; !ok {
		return errors.New("Cannot unmarshal result")
	}

	max := 1
	for {
		_, ok := mapData["L_TRANSACTIONID"+strconv.Itoa(max)]
		if !ok {
			break
		}
		max++
	}

	r.Items = make([]TransactionSearchItem, max)
	for i := 0; i < max; i++ {
		r.Items[i] = createTransactionSearchItem(mapData, i)
	}
	return nil
}

// Transactions returns list of transactions.
func (r *TransactionSearchResponse) Transactions() (list []TransactionSearchItem) {
	for _, item := range r.Items {
		if item.IsTransaction() {
			list = append(list, item)
		}
	}
	return
}

// TransactionIDList returns list of transaction id for payment.
func (r *TransactionSearchResponse) TransactionIDList() (list []string) {
	for _, item := range r.Items {
		if item.IsTransaction() {
			list = append(list, item.TransactionID)
		}
	}
	return
}

// GetSubscribeiStartedDate returns time of subscription started.
func (r *TransactionSearchResponse) GetSubscribeStartedDate() (time.Time, bool) {
        for i := len(r.Items) - 1; i >= 0; i-- {
                item := r.Items[i]
                if item.IsProfile() && item.IsCreated() {
                        return item.Timestamp, true
                }
        }
        return time.Time{}, false
}

// GetCancelDate returns time of canceled.
func (r *TransactionSearchResponse) GetCancelDate() (time.Time, bool) {
	for i := len(r.Items) - 1; i >= 0; i-- {
		item := r.Items[i]
		if item.IsProfile() && item.IsCanceled() {
			return item.Timestamp, true
		}
	}
	return time.Time{}, false
}

// TransactionSearchItem is single transaction.
type TransactionSearchItem struct {
	Timestamp     time.Time
	Timezone      string
	Type          string
	Email         string
	Name          string
	TransactionID string
	Status        string
	Amount        float64
	Currency      string
	FeeAmount     float64
	NetAmount     float64
}

func createTransactionSearchItem(mapData map[string]interface{}, i int) TransactionSearchItem {
	result := TransactionSearchItem{}
	num := strconv.Itoa(i)

	if v, ok := mapData["L_TIMESTAMP"+num].(string); ok {
		result.Timestamp, _ = stringToTime(v)
	}
	if v, ok := mapData["L_TIMEZONE"+num].(string); ok {
		result.Timezone = v
	}
	if v, ok := mapData["L_TYPE"+num].(string); ok {
		result.Type = v
	}
	if v, ok := mapData["L_EMAIL"+num].(string); ok {
		result.Email = v
	}
	if v, ok := mapData["L_NAME"+num].(string); ok {
		result.Name = v
	}
	if v, ok := mapData["L_TRANSACTIONID"+num].(string); ok {
		result.TransactionID = v
	}
	if v, ok := mapData["L_STATUS"+num].(string); ok {
		result.Status = v
	}
	if v, ok := mapData["L_AMT"+num].(float64); ok {
		result.Amount = v
	}
	if v, ok := mapData["L_CURRENCYCODE"+num].(string); ok {
		result.Currency = v
	}
	if v, ok := mapData["L_FEEAMT"+num].(float64); ok {
		result.FeeAmount = v
	}
	if v, ok := mapData["L_NETAMT"+num].(float64); ok {
		result.NetAmount = v
	}
	return result
}

// IsProfile checks the TransactionID is Profile or not.
func (item TransactionSearchItem) IsProfile() bool {
	return strings.HasPrefix(item.TransactionID, "I-")
}

// IsTransaction checks the TransactionID is Transaction or not.
func (item TransactionSearchItem) IsTransaction() bool {
	return !item.IsProfile()
}

// IsRecurringPayment checks the transaction is Recurring Payment or not.
func (item TransactionSearchItem) IsRecurringPayment() bool {
	return item.Type == "Recurring Payment"
}

// IsPayment checks the transaction is Payment or not.
func (item TransactionSearchItem) IsPayment() bool {
	return item.Type == "Payment"
}

// IsCreated checks the transaction is created or not.
func (item TransactionSearchItem) IsCreated() bool {
	return item.Status == "Created"
}

// IsCanceled checks the transaction is canceled or not.
func (item TransactionSearchItem) IsCanceled() bool {
	return item.Status == "Canceled"
}

// IsCompleted checks the transaction is completed or not.
func (item TransactionSearchItem) IsCompleted() bool {
	return item.Status == "Completed"
}
