package config

import "github.com/jinzhu/configor"

type Config struct {
	FluentbitServer string
	Port            int
	MQTTUserName    string
	MQTTPassword    string
}

var _config *Config

func MustGetConfig() Config {
	if _config != nil {
		return *_config
	}

	_config = &Config{}
	err := configor.New(&configor.Config{ENVPrefix: "MFB"}).Load(_config)
	if err != nil {
		panic(err)
	}

	return *_config
}
