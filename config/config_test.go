package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	conf := New("user", "password", "signature")
	assert.EqualValues("", conf.Mode)
	assert.Equal("", conf.AppID)
	assert.Equal("user", conf.User)
	assert.Equal("password", conf.Pass)
	assert.Equal("signature", conf.Signature)
	assert.Equal(defaultVersion, conf.Version)
}

func TestDefaultConfig(t *testing.T) {
	assert := assert.New(t)

	conf := DefaultConfig
	assert.EqualValues("sandbox", conf.Mode)
	assert.Equal("", conf.AppID)
	assert.Equal("sdk-three_api1.sdk.com", conf.User)
	assert.Equal("QFZCWN5HZM8VBG7Q", conf.Pass)
	assert.Equal("A-IzJhZZjhg29XQ2qnhapuwxIDzyAZQ92FRP5dqBzVesOkzbdUONzmOU", conf.Signature)
	assert.Equal(124, conf.Version)
}
