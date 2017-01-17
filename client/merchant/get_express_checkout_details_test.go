package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExpressCheckoutDetails(t *testing.T) {
	assert := assert.New(t)

	cli := testNewDefault()

	// error
	svc := &GetExpressCheckoutDetails{
		Token: "EC-00000000000000000",
	}
	v, err := svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("10410", v.ErrorCode)
	assert.Empty(v.Token)
	assert.Empty(v.PayerID)
	assert.Empty(v.PayerStatus)
	assert.Empty(v.PaymentErrorCode)

	// success
	setSVC := &SetExpressCheckout{
		TotalAmount: 99,
		ReturnURL:   "http://localhost/",
		CancelURL:   "http://localhost/",
		Currency:    CurrencyTWD,
	}
	setResp, err := setSVC.Do(cli)
	assert.Nil(err)

	svc = &GetExpressCheckoutDetails{
		Token: setResp.Token,
	}
	v, err = svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Success", v.ACK)
	assert.Equal("124", v.Version)
	assert.Equal(setResp.Token, v.Token)
	assert.Equal("0", v.PaymentErrorCode)
	assert.Empty(v.PayerID)
	assert.Empty(v.PayerStatus)
	assert.Empty(v.ErrorCode)
}
