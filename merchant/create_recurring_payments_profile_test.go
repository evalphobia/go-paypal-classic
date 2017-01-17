package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRecurringPaymentsProfile(t *testing.T) {
	assert := assert.New(t)

	m := NewDefault()

	// error
	svc := &CreateRecurringPaymentsProfile{
		Token:       "EC-00000000000000000",
		PayerID:     "XXX",
		Amount:      200.0,
		Currency:    CurrencyTWD,
		Description: "this is recurring",
	}
	svc.SetPeriodAsMonth(13)
	svc.SetBillingStartDateFromNow()
	v, err := svc.Do(m)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("11502", v.ErrorCode)
}
