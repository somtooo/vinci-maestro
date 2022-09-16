package events_test

import (
	"fmt"
	"testing"

	"github.com/somtooo/vinci-maestro/events"
	"github.com/somtooo/vinci-maestro/iot"
)

func Test_NewEmitter(t *testing.T) {
	e := events.NewEmitter()
	c := func(data iot.IntentMessage) {
		fmt.Printf("This is the listner for the %s intent\n", data.Intent)
	}

	e.On("goat", c)
	e.Emmit("goat", iot.IntentMessage{Intent: "changeLightBulb"})
}
