package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecurringPaymentsProfileDetails(t *testing.T) {
	assert := assert.New(t)

	m := NewDefault()

	// error
	svc := &GetRecurringPaymentsProfileDetails{
		ProfileID: "I-000000000000",
	}
	v, err := svc.Do(m)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.Equal("000000", v.Build)
	assert.Equal("10001", v.ErrorCode)
}
