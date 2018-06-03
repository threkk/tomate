package main

import (
	"time"
)

// Timer - A Timer holds the countdown and the events that will trigger.
type Timer struct {
	ticker    *time.Ticker
	stop      chan bool
	pause     bool
	Remaining int64                 // Remaining time in the timer.
	Finished  chan bool             // Notifies when it finishes. Useful with routines.
	OnStart   func()                // Executed at the begin of timer.
	OnTick    func(remaining int64) // Executed on every tick.
	OnFinish  func()                // Executed when the timer expires.
}

// NewTimer - Mocks up a new Timer object.
func NewTimer(finished ...chan bool) *Timer {
	var f chan bool
	f = nil
	if len(finished) == 1 {
		f = finished[0]
	}
	timer := &Timer{
		stop:      make(chan bool, 1),
		pause:     false,
		Remaining: 0,
		Finished:  f,
		OnStart:   func() {},
		OnTick:    func(remaining int64) {},
		OnFinish:  func() {},
	}
	return timer
}

// Start - Starts a countdown starting at the amount of seconds passed as
// parameter. It will trigger the OnStart function and for every second it passes,
// if it is not paused it will decrease the counter and trigger the OnTick
// function. When the countdown finishes, it will trigger the OnFinish function.
func (timer *Timer) Start(seconds int64) {
	timer.Remaining = seconds
	timer.OnStart()
	timer.ticker = time.NewTicker(time.Second)

	defer timer.ticker.Stop()
	defer close(timer.stop)

	for {
		select {
		case <-timer.ticker.C:
			if !timer.pause {
				timer.Remaining--
				timer.OnTick(timer.Remaining)

				if timer.Remaining == 0 {
					timer.stop <- true
				}
			}
		case <-timer.stop:
			timer.OnFinish()
			if timer.Finished != nil {
				timer.Finished <- true
			}
			return
		}
	}
}

// Pause - Pauses the timer. It will set the flag to true and it will stop the
// countdown and the tick event.
func (timer *Timer) Pause() {
	timer.pause = true
}

// Resume - Resumes the timer.
func (timer *Timer) Resume() {
	timer.pause = false
}

// Stop -  Stops the timer
func (timer *Timer) Stop() {
	timer.stop <- true
}
