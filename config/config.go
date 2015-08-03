package config

// mode
var (
	modeSanbox     = Mode("sandbox")
	modeProduction = Mode("live")
)

// default sandbox parameters
const (
	defaultUser      = "sdk-three_api1.sdk.com"
	defaultPass      = "QFZCWN5HZM8VBG7Q"
	defaultSignature = "A-IzJhZZjhg29XQ2qnhapuwxIDzyAZQ92FRP5dqBzVesOkzbdUONzmOU"
	defaultVersion   = 124
)

// DefaultConfig is used for development
var DefaultConfig = &Config{
	Mode:      modeSanbox,
	User:      defaultUser,
	Pass:      defaultPass,
	Signature: defaultSignature,
	Version:   defaultVersion,
}

// Mode stands for development mode or live mode
type Mode string

// IsProduction checks the mode is production(=live) or not
func (m Mode) IsProduction() bool {
	return m == modeProduction
}

// Config is struct for base setting for PayPal Classic API
type Config struct {
	Mode      `url:"-"`
	AppID     string `url:"app_id,omitempty"`
	User      string `url:"USER"`
	Pass      string `url:"PWD"`
	Signature string `url:"SIGNATURE"`
	Version   int    `url:"VERSION"`
}

// New returns initialized *Config
func New(user, pass, signature string) *Config {
	return &Config{
		User:      user,
		Pass:      pass,
		Signature: signature,
		Version:   defaultVersion,
	}
}

// SetAsProduction set to live mode
func (c *Config) SetAsProduction() {
	c.Mode = modeProduction
}
