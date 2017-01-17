package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/evalphobia/go-paypal-classic/config"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	conf := config.New("user", "password", "signature")
	m := New(conf)
	assert.Equal(conf, m.Config)
}

func TestNewDefault(t *testing.T) {
	assert := assert.New(t)

	m := NewDefault()
	assert.Equal(config.DefaultConfig, m.Config)
}
