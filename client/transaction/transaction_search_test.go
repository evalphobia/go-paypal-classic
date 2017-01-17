package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/evalphobia/go-paypal-classic/client"
)

func testNewDefault() client.Client {
	return client.NewDefault()
}

func TestTransactionSearch(t *testing.T) {
	assert := assert.New(t)

	cli := testNewDefault()

	// error
	svc := &TransactionSearch{
		ParamStartDate: "2011-01-01T00:00:00Z",
		ProfileID:      "I-XXXXXXXXXXXX",
	}
	v, err := svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("10001", v.ErrorCode)
}
