package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManageRecurringPaymentsProfileStatus(t *testing.T) {
	assert := assert.New(t)

	m := NewDefault()

	// error
	svc := &ManageRecurringPaymentsProfileStatus{
		ProfileID: "I-000000000000",
	}
	svc.SetAsSuspend("foo")
	v, err := svc.Do(m)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.Equal("000000", v.Build)
	assert.Equal("11552", v.ErrorCode)
}