package main

import (
	"fmt"
	"testing"
)

func TestNewTimer(t *testing.T) {
	timer := NewTimer()
	if timer.pause != false || timer.Finished != nil {
		t.Error("Testing not properly initialised")
	}
}

func TestTimer_Toggle(t *testing.T) {
	timer := NewTimer()
	if timer.pause {
		t.Error("Invalid default on pause")
	}

	timer.Toggle()
	if !timer.pause {
		t.Error("Timer not paused")
	}

	timer.Toggle()
	if timer.pause {
		t.Error("Timer still paused")
	}
}

func ExampleTimer_OnTick() {
	timer := NewTimer()
	timer.OnTick = func(c int64, t int64) {
		fmt.Printf("%d", c)
	}
	timer.Start(3)
	// Output: 123
}

func ExampleTimer_OnTick_pause() {
	timer := NewTimer()
	timer.pause = true
	timer.OnTick = func(c int64, t int64) {
		fmt.Printf("%d", c)
	}
	go timer.Start(3)
	timer.stop <- true
	// Output:
}

func ExampleTimer_OnFinish() {
	timer := NewTimer()
	timer.OnTick = func(c int64, t int64) {
		fmt.Printf("%d", c)
	}
	timer.OnFinish = func() {
		fmt.Print("stop")
	}
	timer.Start(3)
	// Output: 123stop
}

func ExampleTimer_Stop() {
	ch := make(chan bool, 1)
	timer := NewTimer()
	timer.Finished = ch
	timer.OnFinish = func() {
		fmt.Print("stop")
	}
	go timer.Start(10)
	timer.Stop()
	<-timer.Finished
	// Output: stop
}
