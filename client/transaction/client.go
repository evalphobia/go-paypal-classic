package transaction

import (
	"time"

	"github.com/evalphobia/go-paypal-classic/client"
	"github.com/evalphobia/go-paypal-classic/config"
)

// New creates client.Client with given config
func New(conf *config.Config) client.Client {
	return client.Client{
		Config: conf,
	}
}

// stringToTime converts string("YYYY-MM-DD HH:mm:ssZ") to time.Time.
func stringToTime(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}

// timeToString converts time.Time to string("YYYY-MM-DD HH:mm:ssZ").
func timeToString(dt time.Time) string {
	return dt.Format(time.RFC3339)
}
