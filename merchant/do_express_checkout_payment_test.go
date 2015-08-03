package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoExpressCheckoutPayment(t *testing.T) {
	assert := assert.New(t)

	m := NewDefault()

	// error
	svc := &DoExpressCheckoutPayment{
		Token:       "EC-00000000000000000",
		PayerID:     "XXX",
		TotalAmount: 200.0,
		Currency:    CurrencyTWD,
	}
	v, err := svc.Do(m)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.Equal("000000", v.Build)
	assert.Equal("10410", v.ErrorCode)
}
