package client

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	p "github.com/Picovoice/picovoice/sdk/go/v2"

	pvrecorder "github.com/Picovoice/pvrecorder/sdk/go"
)

func PrintAudioDevices() {
	if devices, err := pvrecorder.GetAudioDevices(); err != nil {
		log.Fatalf("Error: %s.\n", err.Error())
	} else {
		for i, device := range devices {
			log.Printf("index: %d, device name: %s\n", i, device)
		}
	}
}

func ReadAudioData(audioDeviceIndex int, frameLength int, p *p.Picovoice) {
	// recorder, delete := setupDevice(audioDeviceIndex, frameLength)
	// defer delete()

	recorder := pvrecorder.PvRecorder{
		DeviceIndex:    audioDeviceIndex,
		FrameLength:    frameLength,
		BufferSizeMSec: 1000,
		LogOverflow:    0,
	}
	defer recorder.Delete()
	if err := recorder.Init(); err != nil {
		log.Fatalf("Error: %s.\n", err)
	}

	if err := recorder.Start(); err != nil {
		log.Fatalf("Error: %s.\n", err.Error())
	}

	log.Printf("Using device: %s", recorder.GetSelectedDevice())
	fmt.Println("Listening...")
	done := make(chan os.Signal, 1)
	//waitCh := make(chan struct{})
	signal.Notify(done, os.Interrupt)
	// go func() {
	// 	<-done
	// 	close(waitCh)
	// }()
	//outputCh := make(chan []int16)

waitLoop:
	for {
		select {
		case <-done:
			log.Println("Stopping...")
			break waitLoop
		default:
			pcm, err := recorder.Read()
			if err != nil {
				log.Fatalf("Error: %s.\n", err.Error())
			}

			err2 := p.Process(pcm)
			if err2 != nil {
				log.Fatal(err)
			}

		}
	}

}

func setupDevice(audioDeviceIndex int, frameLength int) (*pvrecorder.PvRecorder, func()) {
	recorder := pvrecorder.PvRecorder{
		DeviceIndex:    audioDeviceIndex,
		FrameLength:    frameLength,
		BufferSizeMSec: 1000,
		LogOverflow:    0,
	}
	delete := func() {
		recorder.Delete()
	}

	if err := recorder.Init(); err != nil {
		log.Fatalf("Error: %s.\n", err)
	}

	if err := recorder.Start(); err != nil {
		log.Fatalf("Error: %s.\n", err.Error())
	}

	log.Printf("Using device: %s", recorder.GetSelectedDevice())

	return &recorder, delete

}
