package transaction

import (
	"errors"
	"strconv"
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
