package sti

import (
	"fmt"
	"log"
	"os"

	. "github.com/Picovoice/picovoice/sdk/go/v2"
	rhn "github.com/Picovoice/rhino/binding/go/v2"
	"github.com/somtooo/vinci-maestro/client"
	"github.com/somtooo/vinci-maestro/events"
)

var keywordPath string = os.Getenv("KEYWORD_PATH")
var contextPath string = os.Getenv("CONTEXT_PATH")
var accessKey string = os.Getenv("ACCESS_KEY")

func StartInference(event *events.Emitter) {
	pico := setup(event)
	audioChan := client.ReadAudioData(0, FrameLength)
	for v := range audioChan {
		err := pico.Process(v)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func setup(event *events.Emitter) Picovoice {
	p := Picovoice{
		RequireEndpoint: true,
	}
	p.AccessKey = accessKey
	p.KeywordPath = keywordPath
	p.ContextPath = contextPath
	p.PorcupineLibraryPath = ""
	p.RhinoLibraryPath = ""
	p.PorcupineModelPath = ""
	p.RhinoModelPath = ""
	ppnSensitivityFloat := float32(0.7)
	p.PorcupineSensitivity = ppnSensitivityFloat
	rhnSensitivityFloat := float32(0.7)
	p.RhinoSensitivity = rhnSensitivityFloat
	endpointDurationFloat := float32(1.0)
	p.EndpointDurationSec = endpointDurationFloat
	p.WakeWordCallback = wakeWordCallback

	f := func(inference rhn.RhinoInference) {

		if inference.IsUnderstood {
			d := events.IntentMessage{}
			d.Intent = inference.Intent
			d.Slots = inference.Slots
			event.Emmit(inference.Intent, d)
			fmt.Println("{")
			fmt.Printf("  intent : '%s'\n", inference.Intent)
			fmt.Println("  slots : {")
			for k, v := range inference.Slots {
				fmt.Printf("    %s : '%s'\n", k, v)
			}
			fmt.Println("  }")
			fmt.Println("}")
		} else {
			fmt.Println("Didn't understand the command")
		}
	}
	p.InferenceCallback = f
	err := p.Init()
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func wakeWordCallback() {
	fmt.Println("[Yes Sir!]")
}

func inferenceCallback(inference rhn.RhinoInference) {

	if inference.IsUnderstood {
		fmt.Println("{")
		fmt.Printf("  intent : '%s'\n", inference.Intent)
		fmt.Println("  slots : {")
		for k, v := range inference.Slots {
			fmt.Printf("    %s : '%s'\n", k, v)
		}
		fmt.Println("  }")
		fmt.Println("}")
	} else {
		fmt.Println("Didn't understand the command")
	}
}
