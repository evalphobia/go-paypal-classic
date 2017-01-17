package merchant

import "github.com/evalphobia/go-paypal-classic/client"

const (
	paymentActionSale    = "sale"
	payerStatusVerified  = "verified"
	itemCategoryDigital  = "Digital"
	billingTypeRecurring = "RecurringPayments"

	statusActive          = "Active"
	profileActive         = "ActiveProfile"
	paymentStatusComleted = "Completed"

	ackSuccess = client.ACKSuccess
)

// PayPal supported currencies
const (
	CurrencyAUD = "AUD"
	CurrencyBRL = "BRL"
	CurrencyCAD = "CAD"
	CurrencyCZK = "CZK"
	CurrencyDKK = "DKK"
	CurrencyEUR = "EUR"
	CurrencyHKD = "HKD"
	CurrencyHUF = "HUF"
	CurrencyILS = "ILS"
	CurrencyJPY = "JPY"
	CurrencyMYR = "MYR"
	CurrencyMXN = "MXN"
	CurrencyNOK = "NOK"
	CurrencyNZD = "NZD"
	CurrencyPHP = "PHP"
	CurrencyPLN = "PLN"
	CurrencyGBP = "GBP"
	CurrencyRUB = "RUB"
	CurrencySGD = "SGD"
	CurrencySEK = "SEK"
	CurrencyCHF = "CHF"
	CurrencyTWD = "TWD"
	CurrencyTHB = "THB"
	CurrencyTRY = "TRY"
	CurrencyUSD = "USD"
)
