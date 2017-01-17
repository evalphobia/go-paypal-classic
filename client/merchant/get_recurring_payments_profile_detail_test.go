package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecurringPaymentsProfileDetails(t *testing.T) {
	assert := assert.New(t)

	cli := testNewDefault()

	// error
	svc := &GetRecurringPaymentsProfileDetails{
		ProfileID: "I-000000000000",
	}
	v, err := svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("11552", v.ErrorCode)
}
