package events

import (
	"log"
)

type IntentMessage struct {
	Intent string
	Slots  map[string]string
}
type Emitter struct {
	intentToFunc map[string]func(intent IntentMessage)
}

func NewEmitter() *Emitter {
	e := &Emitter{make(map[string]func(intent IntentMessage))}
	return e
}

func (e *Emitter) On(intent string, callback func(data IntentMessage)) {
	e.intentToFunc[intent] = callback
}

func (e *Emitter) Emmit(intent string, data IntentMessage) {
	callback, exists := e.intentToFunc[intent]
	if exists {
		callback(data)
	} else {
		log.Println("No listner Present")
	}
}
