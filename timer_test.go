package main

import (
	"fmt"
	"testing"
)

func TestNewTimer_nochan(t *testing.T) {
	timer := NewTimer()
	if timer.Remaining != 0 || timer.pause != false || timer.Finished != nil {
		t.Error("Testing not properly initialised")
	}

	ch := make(chan bool, 1)
	timer = NewTimer(ch, ch)
	if timer.Remaining != 0 || timer.pause != false || timer.Finished != nil {
		t.Error("Testing not properly initialised")
	}
	close(ch)
}

func TestNewTimer_chan(t *testing.T) {
	ch := make(chan bool, 1)
	timer := NewTimer(ch)
	if timer.Remaining != 0 || timer.pause != false || timer.Finished == nil {
		t.Error("Testing not properly initialised")
	}
	close(ch)
}

func TestTimer_Pause(t *testing.T) {
	timer := NewTimer()
	if timer.pause {
		t.Error("Invalid default on pause")
	}

	timer.Pause()
	if !timer.pause {
		t.Error("Timer not paused")
	}

	timer.Pause()
	if !timer.pause {
		t.Error("Timer still not paused")
	}
}

func TestTimer_Resume(t *testing.T) {
	timer := NewTimer()
	if timer.pause {
		t.Error("Invalid default on pause")
	}

	timer.Resume()
	if timer.pause {
		t.Error("Timer paused")
	}

	timer.Resume()
	if timer.pause {
		t.Error("Timer still paused")
	}
}

func ExampleTimer_OnTick() {
	timer := NewTimer()
	timer.OnTick = func(r int64) {
		fmt.Printf("%d", r)
	}
	timer.Start(3)
	// Output: 210
}

func ExampleTimer_OnTick_pause() {
	timer := NewTimer()
	timer.pause = true
	timer.OnTick = func(r int64) {
		fmt.Printf("%d", r)
	}
	go timer.Start(3)
	timer.stop <- true
	// Output:
}

func ExampleTimer_OnStart() {
	timer := NewTimer()
	timer.OnStart = func() {
		fmt.Print("start")
	}
	timer.OnTick = func(r int64) {
		fmt.Printf("%d", r)
	}
	timer.Start(3)
	// Output: start210
}

func ExampleTimer_OnFinish() {
	timer := NewTimer()
	timer.OnTick = func(r int64) {
		fmt.Printf("%d", r)
	}
	timer.OnFinish = func() {
		fmt.Print("stop")
	}
	timer.Start(3)
	// Output: 210stop
}

func ExampleTimer_Stop() {
	ch := make(chan bool, 1)
	timer := NewTimer(ch)
	timer.OnStart = func() {
		fmt.Print("start")
	}
	timer.OnFinish = func() {
		fmt.Print("stop")
	}
	go timer.Start(10)
	timer.Stop()
	<-timer.Finished
	// Output: startstop
}
