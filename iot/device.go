package iot

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/somtooo/vinci-maestro/events"
)

type exposes struct {
	State      string `json:"state"`
	Brightness int    `json:"brightness"`
	Color      color  `json:"color"`
}

type brightness struct {
	Brightness int `json:"brightness"`
}

type brightness_move struct {
	Brightness_move int `json:"brightness_move"`
}

type state struct {
	State string `json:"state"`
}

type color struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

const baseTopic = "zigbee2mqtt/room/set"

func Start(event *events.Emitter, client MQTT.Client) {

	event.On("smarthome.lights.changeColor", func(data events.IntentMessage) {
		c := exposes{}
		if data.Slots["color"] == "red" {
			c = exposes{
				Color: color{
					R: 255,
					G: 0,
					B: 0,
				},
			}
		} else if data.Slots["color"] == "blue" {
			c = exposes{
				Color: color{
					R: 0,
					G: 0,
					B: 255,
				},
			}
		} else if data.Slots["color"] == "green" {
			c = exposes{
				Color: color{
					R: 0,
					G: 255,
					B: 0,
				},
			}
		} else if data.Slots["color"] == "yellow" {
			c = exposes{
				Color: color{
					R: 255,
					G: 255,
					B: 0,
				},
			}
		} else if data.Slots["color"] == "white" {
			c = exposes{
				Color: color{
					R: 255,
					G: 255,
					B: 255,
				},
			}
		}
		topic := "zigbee2mqtt/" + data.Slots["friendlyName"] + "/set"
		b, err := json.Marshal(c)
		if err != nil {
			panic(err)
		}
		publish(topic, 0, false, b, client)
	})

	event.On("smarthome.lights.changeState", func(data events.IntentMessage) {
		//topic := "zigbee2mqtt/" + data.Slots["friendlyName"] + "/set"
		c := state{State: data.Slots["state"]}
		b, err := json.Marshal(c)
		fmt.Printf("This is the Json %s", b)
		if err != nil {
			panic(err)
		}
		publish(baseTopic, 0, false, b, client)

	})

	event.On("smarthome.lights.checkState", func(data events.IntentMessage) {
		//topic := "zigbee2mqtt/" + data.Slots["friendlyName"] + "/get"
		c := state{}
		b, err := json.Marshal(c)
		fmt.Printf("This is the Json %s", b)
		if err != nil {
			panic(err)
		}

		f := func(c MQTT.Client, m MQTT.Message) {
			fmt.Println(m.Topic())
			fmt.Println(string(m.Payload()[:]))
		}
		if token := client.Subscribe("zigbee2mqtt/#", byte(0), f); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

		publish("zigbee2mqtt/room/get", 0, false, b, client)

	})

	event.On("smarthome.lights.setBrightness", func(data events.IntentMessage) {
		intVar, err1 := strconv.Atoi(data.Slots["brightness"])
		if err1 != nil {
			log.Fatal("Conversion failed")

		}
		c := brightness{Brightness: intVar}
		b, err := json.Marshal(c)
		fmt.Printf("This is the Json %s", b)
		if err != nil {
			panic(err)
		}
		publish(baseTopic, 0, false, b, client)
	})

	event.On("smarthome.lights.checkBrightness", func(data events.IntentMessage) {
		c := brightness{}
		b, err := json.Marshal(c)
		fmt.Printf("This is the Json %s", b)
		if err != nil {
			panic(err)
		}

		f := func(c MQTT.Client, m MQTT.Message) {
			fmt.Println(m.Topic())
			fmt.Println(string(m.Payload()[:]))
		}
		if token := client.Subscribe("zigbee2mqtt/room/get", byte(0), f); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}

		publish("zigbee2mqtt/room/set", 0, false, b, client)

	})

	event.On("smarthome.lights.brightnessInc", func(data events.IntentMessage) {
		c := brightness_move{Brightness_move: 30}
		b, err := json.Marshal(c)
		fmt.Printf("This is the Json %s", b)
		if err != nil {
			panic(err)
		}
		publish("zigbee2mqtt/room/set", 0, false, b, client)

	})
	event.On("smarthome.lights.brightnessDec", func(data events.IntentMessage) {
		c := brightness_move{Brightness_move: -30}
		b, err := json.Marshal(c)
		fmt.Printf("This is the Json %s", b)
		if err != nil {
			panic(err)
		}
		publish("zigbee2mqtt/room/set", 0, false, b, client)

	})
	event.On("smarthome.switch.changeState", func(data events.IntentMessage) {})
	event.On("smarthome.switch.checkState", func(data events.IntentMessage) {})
}

func publish(topic string, qos byte, retained bool, payload interface{}, client MQTT.Client) {
	token := client.Publish(topic, byte(qos), false, payload)
	fmt.Println("published " + topic)
	token.Wait()
}
