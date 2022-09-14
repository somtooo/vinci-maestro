package client

import (
	"log"

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

func ReadAudioData(audioDeviceIndex int, frameLength int) (<-chan []int16, func()) {
	recorder, delete := setupDevice(audioDeviceIndex, frameLength)
	defer delete()
	done := make(chan struct{})
	outputCh := make(chan []int16)
	cancel := func() {
		close(done)
	}
	go func() {
		for {
			select {
			case <-done:
				log.Println("Stopping...")
				close(outputCh)
				return
			default:
				pcm, err := recorder.Read()
				if err != nil {
					log.Fatalf("Error: %s.\n", err.Error())
				}
				outputCh <- pcm
			}
		}
	}()
	return outputCh, cancel

}

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
		log.Fatalf("Error: %s.\n", err.Error())
	}

	if err := recorder.Start(); err != nil {
		log.Fatalf("Error: %s.\n", err.Error())
	}

	log.Printf("Using device: %s", recorder.GetSelectedDevice())

	return recorder, delete

}
