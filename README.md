# go-paypal-classic

go-paypal-classic is library for [PayPal Classic API](https://developer.paypal.com/docs/classic/api/)

## Current 
Supported API list

- Express Checkout
    - [CreateRecurringPaymentsProfile](https://developer.paypal.com/docs/classic/api/merchant/CreateRecurringPaymentsProfile_API_Operation_NVP/)
    - [DoExpressCheckoutPayment](https://developer.paypal.com/docs/classic/api/merchant/DoExpressCheckoutPayment_API_Operation_NVP/)
    - [GetExpressCheckoutDetails](https://developer.paypal.com/docs/classic/api/merchant/GetExpressCheckoutDetails_API_Operation_NVP/)
    - [GetRecurringPaymentsProfileDetails](https://developer.paypal.com/docs/classic/api/merchant/GetRecurringPaymentsProfileDetails_API_Operation_NVP/)
    - [ManageRecurringPaymentsProfileStatus](https://developer.paypal.com/docs/classic/api/merchant/ManageRecurringPaymentsProfileStatus_API_Operation_NVP/)
    - [SetExpressCheckout](https://developer.paypal.com/docs/classic/api/merchant/SetExpressCheckout_API_Operation_NVP/)

## Quick Usage

### SetExpressCheckout

```go
import (
    "os"

    "github.com/evalphobia/go-paypal-classic/config"
    "github.com/evalphobia/go-paypal-classic/merchant"
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
    "github.com/evalphobia/go-paypal-classic/merchant"
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
    "github.com/evalphobia/go-paypal-classic/merchant"
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
    "github.com/evalphobia/go-paypal-classic/merchant"
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
    "github.com/evalphobia/go-paypal-classic/merchant"
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
    "github.com/evalphobia/go-paypal-classic/merchant"
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
