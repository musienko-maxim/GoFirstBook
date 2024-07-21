package main

import "github.com/headfirstgo/gadget"

type Player__ interface {
	Play(string)
	Stop()
}

func TryOut__(player Player__) {
	player.Play("Test Track")
	player.Stop()
	recorder := player.(gadget.TapeRecorder)
	recorder.Record()
}

func main() {
	var player Player__ = gadget.TapePlayer{}
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		println("Player__ was not a TapeRecorder")
	}

}
