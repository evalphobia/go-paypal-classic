package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManageRecurringPaymentsProfileStatus(t *testing.T) {
	assert := assert.New(t)

	cli := testNewDefault()

	// error
	svc := &ManageRecurringPaymentsProfileStatus{
		ProfileID: "I-000000000000",
	}
	svc.SetAsSuspend("foo")
	v, err := svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("11552", v.ErrorCode)
}
