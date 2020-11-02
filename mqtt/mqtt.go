package mqtt

import (
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"github.com/ttlv/common_utils/utils"
	"github.com/ttlv/mfb/global"
)

type MQTTClient struct {
	MC mqtt.Client
}

func subCallBackFunc(client mqtt.Client, msg mqtt.Message) {
	logrus.Infof("Subscribe: Topic is [%s]; msg is [%s]\n", msg.Topic(), string(msg.Payload()))
	if err := utils.PushToFluentBit(global.Remote, global.Topic, global.UserName, global.Password, string(msg.Payload())); err != nil {
		logrus.Errorf("can't push meaasge to fluent-bit,err is :%v", err)
	}
}

func NewMQClient(broker, user, passwd string) (err error, client MQTTClient) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetUsername(user)
	opts.SetPassword(passwd)

	client.MC = mqtt.NewClient(opts)
	if token := client.MC.Connect(); token.Wait() && token.Error() != nil {
		return token.Error(), client
	}
	return
}

func (mc *MQTTClient) Subscribe(topic string) {
	mc.MC.Subscribe(topic, 0x00, subCallBackFunc)
}

func (mc *MQTTClient) Publish(topic string, content string) {
	mc.MC.Publish(topic, 0x00, true, content)
}
