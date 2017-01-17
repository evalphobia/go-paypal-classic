package merchant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetExpressCheckout(t *testing.T) {
	assert := assert.New(t)
	redirectSuccess := "https://www.sandbox.paypal.com/webscr?cmd=_express-checkout&token=EC-"

	cli := testNewDefault()

	// error
	svc := &SetExpressCheckout{
		TotalAmount: 100.00,
	}
	v, err := svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("10404", v.ErrorCode)
	assert.Equal("10405", v.ErrorCode1)
	assert.Equal("Transaction refused because of an invalid argument. See additional error messages for details.", v.ShortMessage)
	assert.Equal("ReturnURL is missing.", v.LongMessage)
	assert.Equal("Error", v.SeverityCode)

	// success
	svc = &SetExpressCheckout{
		TotalAmount: 100.00,
		ReturnURL:   "http://localhost/",
		CancelURL:   "http://localhost/",
	}
	v, err = svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Success", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Empty(v.ErrorCode)
	assert.Empty(v.ErrorCode1)
	assert.Empty(v.ShortMessage)
	assert.Empty(v.LongMessage)
	assert.Empty(v.SeverityCode)
	assert.Contains(v.RedirectURL(), redirectSuccess)

	t.Logf("TestSetExpressCheckout: %s", v.RedirectURL())
}

func TestSetExpressCheckoutDigital(t *testing.T) {
	assert := assert.New(t)
	redirectSuccess := "https://www.sandbox.paypal.com/webscr?cmd=_express-checkout&token=EC-"

	cli := testNewDefault()

	// error
	svc := &SetExpressCheckout{
		TotalAmount: 100.00,
	}
	svc.SetAsDigitalCategory()
	v, err := svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("10404", v.ErrorCode)
	assert.Equal("10405", v.ErrorCode1)
	assert.Equal("Transaction refused because of an invalid argument. See additional error messages for details.", v.ShortMessage)
	assert.Equal("ReturnURL is missing.", v.LongMessage)
	assert.Equal("Error", v.SeverityCode)

	// success with digital item
	svc = &SetExpressCheckout{
		ItemAmount: 99,
		ItemName:   "hoge",
		ReturnURL:  "http://localhost/",
		CancelURL:  "http://localhost/",
		Currency:   CurrencyTWD,
	}
	svc.SetAsDigitalCategory()
	v, err = svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Success", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Empty(v.ErrorCode)
	assert.Empty(v.ErrorCode1)
	assert.Empty(v.ShortMessage)
	assert.Empty(v.LongMessage)
	assert.Empty(v.SeverityCode)
	assert.Contains(v.RedirectURL(), redirectSuccess)

	t.Logf("TestSetExpressCheckoutDigital: %s", v.RedirectURL())
}

func TestSetExpressCheckoutRecurring(t *testing.T) {
	assert := assert.New(t)
	redirectSuccess := "https://www.sandbox.paypal.com/webscr?cmd=_express-checkout&token=EC-"

	cli := testNewDefault()

	// error
	svc := &SetExpressCheckout{
		TotalAmount: 100.00,
	}
	svc.SetAsRecurringPayment("this is recurring")
	v, err := svc.Do(cli)

	assert.Nil(err)
	assert.Equal("Failure", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Equal("10404", v.ErrorCode)
	assert.Equal("10405", v.ErrorCode1)
	assert.Equal("Transaction refused because of an invalid argument. See additional error messages for details.", v.ShortMessage)
	assert.Equal("ReturnURL is missing.", v.LongMessage)
	assert.Equal("Error", v.SeverityCode)

	// success with digital item
	svc = &SetExpressCheckout{
		ItemAmount: 99,
		ItemName:   "hoge",
		ReturnURL:  "http://localhost/",
		CancelURL:  "http://localhost/",
		Currency:   CurrencyTWD,
	}
	svc.SetAsRecurringPayment("this is recurring")
	v, err = svc.Do(cli)
	assert.Nil(err)
	assert.Equal("Success", v.ACK)
	assert.Equal("124", v.Version)
	assert.NotEmpty(v.Build)
	assert.Empty(v.ErrorCode)
	assert.Empty(v.ErrorCode1)
	assert.Empty(v.ShortMessage)
	assert.Empty(v.LongMessage)
	assert.Empty(v.SeverityCode)
	assert.Contains(v.RedirectURL(), redirectSuccess)
	t.Logf("TestSetExpressCheckoutRecurring: %s", v.RedirectURL())
}
