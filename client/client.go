package client

import (
	"fmt"
	"log"
	"os"
	"os/signal"

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

func ReadAudioData(audioDeviceIndex int, frameLength int) chan []int16 {
	recorder, delete := setupDevice(audioDeviceIndex, frameLength)
	log.Printf("Using device: %s", recorder.GetSelectedDevice())
	fmt.Println("Listening...")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	outputCh := make(chan []int16)
	go func() {
		defer delete()
	waitLoop:
		for {
			select {
			case <-done:
				log.Println("Stopping...")
				close(outputCh)
				break waitLoop
			default:
				pcm, err := recorder.Read()
				if err != nil {
					log.Fatalf("Error: %s.\n", err.Error())
				}
				outputCh <- pcm

			}
		}
	}()
	return outputCh

}

//quickCheck
//typestate
//john hues erlang
func setupDevice(audioDeviceIndex int, frameLength int) (pvrecorder.PvRecorder, func()) {
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

	return recorder, delete

}
