package timer

import (
	"time"
)

// Timer - A Timer holds the countdown and the events that will trigger.
type Timer struct {
	ticker   *time.Ticker
	stop     chan bool
	pause    bool
	Finished chan bool                        // If present, it will notify when it finishes.
	OnTick   func(current int64, total int64) // Executed on every tick.
	OnFinish func()                           // Executed when the timer expires.
}

// NewTimer - Mocks up a new SyncTimer object.
func NewTimer() *Timer {
	timer := &Timer{
		stop:     make(chan bool, 1),
		pause:    false,
		Finished: nil,
		OnTick:   func(current int64, total int64) {},
		OnFinish: func() {},
	}
	return timer
}

// Start - Starts a countdown starting at the amount of seconds passed as
// parameter. For every second it passes, if it is not paused it will increase
// the counter and trigger the OnTick function. When the countdown finishes,
// it will trigger the OnFinish function.
func (timer *Timer) Start(seconds int64) {
	current := int64(0)
	total := seconds
	timer.ticker = time.NewTicker(time.Second)

	defer timer.ticker.Stop()

	for {
		select {
		case <-timer.ticker.C:
			if !timer.pause {
				current++
				timer.OnTick(current, total)

				if current == total {
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

// Toggle - If the timer is paused, it will resume it. If it is running, it will
// stop it.
func (timer *Timer) Toggle() {
	timer.pause = !timer.pause
}

// Stop -  Stops the timer
func (timer *Timer) Stop() {
	timer.stop <- true
}
