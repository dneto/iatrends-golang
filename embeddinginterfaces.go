package main

import "fmt"

type Device interface {
	Info() string
}

type AudioDevice interface {
	Device
	PlayAudio()
}

type Radio struct {
	Model string
}

func (r *Radio) PlayAudio() {
	fmt.Println("Playing music ♬ ♬")
}

func (r *Radio) Info() string {
	return r.Model
}

func main() {
	var radio AudioDevice = &Radio{"MyRadio"}
	fmt.Println(radio.Info())
	radio.PlayAudio()
}
