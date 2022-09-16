package sti

import (
	"fmt"
	"log"
	"os"

	. "github.com/Picovoice/picovoice/sdk/go/v2"
	rhn "github.com/Picovoice/rhino/binding/go/v2"
	"github.com/somtooo/vinci-maestro/client"
	"github.com/somtooo/vinci-maestro/iot"
)

var keywordPath string = os.Getenv("KEYWORD_PATH")
var contextPath string = os.Getenv("CONTEXT_PATH")
var accessKey string = os.Getenv("ACCESS_KEY")

func StartInference() {
	pico := setup()
	audioChan := client.ReadAudioData(0, FrameLength)
	for v := range audioChan {
		err := pico.Process(v)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func setup() Picovoice {
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
	ppnSensitivityFloat := float32(0.5)
	p.PorcupineSensitivity = ppnSensitivityFloat
	rhnSensitivityFloat := float32(0.5)
	p.RhinoSensitivity = rhnSensitivityFloat
	endpointDurationFloat := float32(1.0)
	p.EndpointDurationSec = endpointDurationFloat
	p.WakeWordCallback = wakeWordCallback
	p.InferenceCallback = inferenceCallback
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
		r := iot.IntentMessage{Intent: inference.Intent,
			Slots: inference.Slots}
		iot.EmitIntentMessage(r)
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
