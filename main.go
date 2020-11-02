package main

import (
	"github.com/ttlv/mfb/global"
	"github.com/ttlv/mfb/mqtt"
	"sync"
)

func main() {
	var (
		mc  mqtt.MQTTClient
		err error
		wg  sync.WaitGroup
	)
	if err, mc = mqtt.NewMQClient(global.Remote, global.UserName, global.Password); err != nil {
		panic(err)
	}
	for {
		wg.Add(1)
		defer wg.Done()
		mc.Subscribe(global.Topic)
	}
	wg.Wait()
}
