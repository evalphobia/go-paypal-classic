# go-paypal-classic

[![GoDoc][1]][2] [![Apache 2.0 License][3]][4] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Downloads][15]][16]

[1]: https://godoc.org/github.com/evalphobia/go-paypal-classic?status.svg
[2]: https://godoc.org/github.com/evalphobia/go-paypal-classic
[3]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/go-paypal-classic.svg
[6]: https://github.com/evalphobia/go-paypal-classic/releases/latest
[7]: https://travis-ci.org/evalphobia/go-paypal-classic.svg?branch=master
[8]: https://travis-ci.org/evalphobia/go-paypal-classic
[9]: https://coveralls.io/repos/evalphobia/go-paypal-classic/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/go-paypal-classic?branch=master
[11]: https://codecov.io/github/evalphobia/go-paypal-classic/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/go-paypal-classic?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/go-paypal-classic
[14]: https://goreportcard.com/report/github.com/evalphobia/go-paypal-classic
[15]: https://img.shields.io/github/downloads/evalphobia/go-paypal-classic/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/go-paypal-classic/releases
[17]: https://img.shields.io/github/stars/evalphobia/go-paypal-classic.svg
[18]: https://github.com/evalphobia/go-paypal-classic/stargazers


go-paypal-classic is library for [PayPal Classic API](https://developer.paypal.com/docs/classic/api/)

## Current Supported API list

- Express Checkout
    - [CreateRecurringPaymentsProfile](https://developer.paypal.com/docs/classic/api/merchant/CreateRecurringPaymentsProfile_API_Operation_NVP/)
    - [DoExpressCheckoutPayment](https://developer.paypal.com/docs/classic/api/merchant/DoExpressCheckoutPayment_API_Operation_NVP/)
    - [GetExpressCheckoutDetails](https://developer.paypal.com/docs/classic/api/merchant/GetExpressCheckoutDetails_API_Operation_NVP/)
    - [GetRecurringPaymentsProfileDetails](https://developer.paypal.com/docs/classic/api/merchant/GetRecurringPaymentsProfileDetails_API_Operation_NVP/)
    - [ManageRecurringPaymentsProfileStatus](https://developer.paypal.com/docs/classic/api/merchant/ManageRecurringPaymentsProfileStatus_API_Operation_NVP/)
    - [SetExpressCheckout](https://developer.paypal.com/docs/classic/api/merchant/SetExpressCheckout_API_Operation_NVP/)
- Transaction
    - [TransactionSearch](https://developer.paypal.com/docs/classic/api/merchant/TransactionSearch_API_Operation_NVP/)

## Quick Usage

### SetExpressCheckout

```go
import (
    "os"

    "github.com/evalphobia/go-paypal-classic/config"
    "github.com/evalphobia/go-paypal-classic/client/merchant"
)

func main() {
    user := os.Getenv("PAYPAL_USER")
    pwd := os.Getenv("PAYPAL_PWD")
    signature := os.Getenv("PAYPAL_SIGNATURE")
    conf := config.New(user, pwd, signature)

    mer := merchant.New(conf)
    sec := &merchant.SetExpressCheckout{
        TotalAmount: 100.00,
        ReturnURL:   "http://localhost/confirm",
        CancelURL:   "http://localhost/cancel",
        Currency:    merchant.CurrencyUSD,
    }

    resp, err := sec.Do(mer)
    if err != nil {
        panic("error occured on SetExpressCheckout api request")
    }

    if resp.IsSuccess() {
        resp.RedirectURL()
        // => https://www.paypal.com/webscr?cmd=_express-checkout&token=<TOKEN>
    }
}
```

### GetExpressCheckoutDetails

```go
import (
    "github.com/evalphobia/go-paypal-classic/client/merchant"
)

func main() {
    gecd = &merchant.GetExpressCheckoutDetails{
        Token: "EC-XXXXXXXXXXXX",
    }

    resp, err := gecd.Do(mer)
    if err != nil {
        panic("error occured on GetExpressCheckoutDetails api request")
    }

    if resp.IsSuccess() {
        // verified or unverified
        resp.IsPayerVerified()
    }
}
```

### DoExpressCheckoutPayment

```go
import (
    "github.com/evalphobia/go-paypal-classic/client/merchant"
)

func main() {
    decp = &merchant.DoExpressCheckoutPayment{
        Token:       "EC-XXXXXXXXXXXX",
        PayerID:     "XXX",
        TotalAmount: 100.0,
        Currency:    merchant.CurrencyUSD,
    }

    resp, err := decp.Do(mer)
    if err != nil {
        panic("error occured on DoExpressCheckoutPayment api request")
    }

    if resp.IsSuccess() {
        // payment request is success or not
        resp.IsPaymentSuccess()
    }
}
```

### DoExpressCheckoutPayment

```go
import (
    "github.com/evalphobia/go-paypal-classic/client/merchant"
)

func main() {
    decp = &merchant.CreateRecurringPaymentsProfile{
        Token:       "EC-XXXXXXXXXXXX",
        PayerID:     "XXX",
        TotalAmount: 100.0,
        Currency:    merchant.CurrencyUSD,
        Description: "this is recurring payment",
    }
    decp.SetPeriodAsMonth(3) // once every three months
    decp.SetBillingStartDateFromNow() // the 1st billing starts three month later

    resp, err := decp.Do(mer)
    if err != nil {
        panic("error occured on CreateRecurringPaymentsProfile api request")
    }

    if resp.IsSuccess() {
        // created recurring profile id
        // resp.ProfileID
    }
}
```

### GetRecurringPaymentsProfileDetails

```go
import (
    "github.com/evalphobia/go-paypal-classic/client/merchant"
)

func main() {
    grpp = &merchant.GetRecurringPaymentsProfileDetails{
        ProfileID: "I-000000000000",
    }
    resp, err := grpp.Do(mer)
    if err != nil {
        panic("error occured on GetRecurringPaymentsProfileDetails api request")
    }

    if resp.IsSuccess() {
        // recurring payment status
        // resp.Status
    }
}
```

### ManageRecurringPaymentsProfileStatus

```go
import (
    "github.com/evalphobia/go-paypal-classic/client/merchant"
)

func main() {
    mrpp := &merchant.ManageRecurringPaymentsProfileStatus{
        ProfileID: "I-000000000000",
    }
    svc.SetAsCancel("You must pay my bill!")

    resp, err := mrpp.Do(mer)
    if err != nil {
        panic("error occured on ManageRecurringPaymentsProfileStatus api request")
    }

    if resp.IsSuccess() {
        // profile id will be present when success
        // resp.ProfileID
    }
}
```

### TransactionSearch

```go
import (
    "time"

    "github.com/evalphobia/go-paypal-classic/client/transaction"
)

func main() {
    ts = &transaction.TransactionSearch{
        StartDate: time.Now(),
        ProfileID: "I-000000000000",
    }
    resp, err := ts.Do(cli)
    if err != nil {
        panic("error occured on TransactionSearch api request")
    }

    if resp.IsSuccess() {
        // transaction list
        // resp.Items
    }
}
```
