package iot_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/somtooo/vinci-maestro/events"
	"github.com/somtooo/vinci-maestro/iot"
	"github.com/somtooo/vinci-maestro/mqtt"
)

func Test_Send(t *testing.T) {
	e := events.NewEmitter()
	settings := mqtt.MqttSettings{
		Broker: "tcp://localhost:1883",
		Id:     "testgoid",
	}

	c := mqtt.Connect(settings)
	fmt.Println(c.IsConnected())

	iot.Start(e, c)
	e.Emmit("smarthome.lights.checkState", events.IntentMessage{})
	time.Sleep(5000000)

}
