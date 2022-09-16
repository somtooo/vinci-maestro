package mqtt_test

import (
	"fmt"
	"testing"

	"github.com/somtooo/vinci-maestro/mqtt"
)

func Test_connect(t *testing.T) {
	settings := mqtt.MqttSettings{
		Broker: "tcp://localhost:1883",
		Id:     "testgoid",
	}

	c := mqtt.Connect(settings)
	fmt.Println(c.IsConnected())
	c.Disconnect(200)

}
