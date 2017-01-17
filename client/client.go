package client

import (
	"github.com/evalphobia/go-paypal-classic/config"
	"github.com/evalphobia/go-paypal-classic/request"
)

const (
	endpointSandbox    = "https://api-3t.sandbox.paypal.com/nvp"
	endpointProduction = "https://api-3t.paypal.com/nvp"

	redirectSandbox    = "https://www.sandbox.paypal.com/webscr"
	redirectProduction = "https://www.paypal.com/webscr"
)

// Client is base struct for PayPal Classic API
type Client struct {
	Config *config.Config `url:",squash"`
}

// New creates Client with given config
func New(conf *config.Config) Client {
	return Client{
		Config: conf,
	}
}

// NewDefault creates Client with default config
func NewDefault() Client {
	return New(config.DefaultConfig)
}

// Call sends HTTP request to PayPal api.
func (c Client) Call(param interface{}, result interface{}) error {
	p := parameter{
		Common: c,
		Extra:  param,
	}
	if c.Config.IsProduction() {
		return request.CallPOST(endpointProduction, p, result)
	}
	return request.CallPOST(endpointSandbox, p, result)
}

type parameter struct {
	Common Client      `url:",squash"`
	Extra  interface{} `url:",squash"`
}

// RedirectBase returns base url of PayPal ExpressCheckout
func (c Client) RedirectBase() string {
	if c.Config.IsProduction() {
		return redirectProduction
	}
	return redirectSandbox
}
