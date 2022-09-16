package iot

import "fmt"

type IntentMessage struct {
	Intent string
	Slots  map[string]string
}

func EmitIntentMessage(data IntentMessage) {
	fmt.Println("emit 'intentMessage' data")
	fmt.Println(data)
}

func OnIntentMessage(callback func(data IntentMessage)) {
	fmt.Println(". on intentMessage callback")
}
