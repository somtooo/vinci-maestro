package main

import (
	"fmt"

	"github.com/somtooo/vinci-maestro/events"
	"github.com/somtooo/vinci-maestro/iot"
	"github.com/somtooo/vinci-maestro/mqtt"
	"github.com/somtooo/vinci-maestro/sti"
)

func main() {
	e := events.NewEmitter()
	settings := mqtt.MqttSettings{
		Broker: "tcp://localhost:1883",
		Id:     "testgoid",
	}

	c := mqtt.Connect(settings)
	fmt.Println(c.IsConnected())

	iot.Start(e, c)

	sti.StartInference(e)
}
