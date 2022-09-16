package mqtt

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttSettings struct {
	Broker    string
	User      string
	Password  string
	Id        string
	Cleansess bool
	Store     string
}

// type mqttClient struct {
// 	MQTT.Client
// }

// type client MQTT.Client

func Connect(settings MqttSettings) MQTT.Client {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(settings.Broker)
	opts.SetClientID(settings.Id)
	opts.SetUsername(settings.User)
	opts.SetPassword(settings.Password)
	opts.SetCleanSession(settings.Cleansess)
	if settings.Store != "" {
		opts.SetStore(MQTT.NewFileStore(settings.Store))
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Error")
		panic(token.Error())
	}
	log.Printf("Started Mqtt Client %s", settings.Id)
	return client
}
