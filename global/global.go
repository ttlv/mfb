package global

import (
	"fmt"
	"github.com/ttlv/mfb/config"
)

var (
	cfg      = config.MustGetConfig()
	Topic    = "$ke/events/device/+/data/update"
	Remote   = fmt.Sprintf("%v:%v", cfg.FluentbitServer, cfg.Port)
	UserName = cfg.MQTTUserName
	Password = cfg.MQTTPassword
)
