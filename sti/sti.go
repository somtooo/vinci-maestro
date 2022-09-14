package sti

import (
	"fmt"
	"log"
	"os"

	//"os/signal"

	. "github.com/Picovoice/picovoice/sdk/go/v2"
	//pvrecorder "github.com/Picovoice/pvrecorder/sdk/go"
	rhn "github.com/Picovoice/rhino/binding/go/v2"
	"github.com/somtooo/vinci-maestro/client"
)

var keywordPath string = os.Getenv("KEYWORD_PATH")
var contextPath string = os.Getenv("CONTEXT_PATH")
var accessKey string = os.Getenv("ACCESS_KEY")

func StartInference() {
	// pico := setup()

	// go func() {
	// 	for v := range c {
	// 		err := pico.Process(v)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// }()
	p := setup()
	client.ReadAudioData(0, FrameLength, &p)
	// 	recorder := pvrecorder.PvRecorder{
	// 		DeviceIndex:    0,
	// 		FrameLength:    FrameLength,
	// 		BufferSizeMSec: 1000,
	// 		LogOverflow:    0,
	// 	}
	// 	defer recorder.Delete()
	// 	if err := recorder.Init(); err != nil {
	// 		log.Fatalf("Error: %s.\n", err)
	// 	}

	// 	if err := recorder.Start(); err != nil {
	// 		log.Fatalf("Error: %s.\n", err.Error())
	// 	}

	// 	log.Printf("Using device: %s", recorder.GetSelectedDevice())
	// 	fmt.Println("Listening...")
	// 	done := make(chan os.Signal, 1)
	// 	//waitCh := make(chan struct{})
	// 	signal.Notify(done, os.Interrupt)
	// 	// go func() {
	// 	// 	<-done
	// 	// 	close(waitCh)
	// 	// }()
	// 	//outputCh := make(chan []int16)

	// waitLoop:
	// 	for {
	// 		select {
	// 		case <-done:
	// 			log.Println("Stopping...")
	// 			break waitLoop
	// 		default:
	// 			pcm, err := recorder.Read()
	// 			if err != nil {
	// 				log.Fatalf("Error: %s.\n", err.Error())
	// 			}

	// 			err2 := p.Process(pcm)
	// 			if err2 != nil {
	// 				log.Fatal(err)
	// 			}

	// 		}
	// 	}
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

	// validate Porcupine sensitivity
	ppnSensitivityFloat := float32(0.5)
	p.PorcupineSensitivity = ppnSensitivityFloat

	// validate Rhino sensitivity
	rhnSensitivityFloat := float32(0.5)
	p.RhinoSensitivity = rhnSensitivityFloat

	// validate endpoint duration
	endpointDurationFloat := float32(1.0)
	p.EndpointDurationSec = endpointDurationFloat

	p.WakeWordCallback = func() { fmt.Println("[Yes Sir!]") }
	p.InferenceCallback = func(inference rhn.RhinoInference) {
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
