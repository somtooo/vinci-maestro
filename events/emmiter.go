package events

import (
	"github.com/somtooo/vinci-maestro/iot"
	"log"
)

type Emitter struct {
	intentToFunc map[string]func(intent iot.IntentMessage)
}

func NewEmitter() *Emitter {
	e := &Emitter{make(map[string]func(intent iot.IntentMessage))}
	return e
}

func (e *Emitter) On(intent string, callback func(data iot.IntentMessage)) {
	e.intentToFunc[intent] = callback
}

func (e *Emitter) Emmit(intent string, data ...iot.IntentMessage) {
	callback, exists := e.intentToFunc[intent]
	if exists {
		if data != nil {
			callback(data[0])
		}
	} else {
		log.Println("No listner Present")
	}
}
