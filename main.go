package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const version = "1.0.0"

var signals chan os.Signal
var size int
var duration time.Duration
var message string
var isRepeat bool
var isVersion bool

func usage() {
	fmt.Printf("tomato - Simple pomodoro 🍅 (v%s)\n", version)
	fmt.Printf("\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("  tomate [-repeat] [-duration <duration>] [-message <message>]")
	fmt.Printf("  tomate -h | -help\n")
	fmt.Printf("  tomato -v | -version\n")
	fmt.Printf("\n")
	fmt.Printf("Options:\n")

	flag.PrintDefaults()
}

// lol
func leftpad(str string, pad int) string {
	gen := fmt.Sprintf("%%%ds", pad/2)
	return fmt.Sprintf(gen, str)
}

func init() {
	size = int(getColumns())

	flag.DurationVar(&duration, "duration", 0, "Duration of the timer. 25 minutes by default.")
	flag.StringVar(&message, "message", "", "Message to display once the timer finishes.")
	flag.BoolVar(&isRepeat, "repeat", false, "Restart the pomodoro once it finishes.")
	flag.BoolVar(&isVersion, "version", false, "Show the version number")

	flag.Usage = usage

	signals = make(chan os.Signal)
}

func main() {
	flag.Parse()

	if isVersion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if duration == 0 {
		usage()
		os.Exit(0)
	}

	// Uses the Homebrew bar ui.
	ui := &HomeBrewBar{Size: size}

	timer := NewTimer()
	timer.OnTick = func(current int64, total int64) {
		bar := ui.Frame(current, total)
		fmt.Printf("\r%s", bar)
	}

	if message != "" {
		timer.OnFinish = func() {
			fmt.Printf("\n\n%s\n\n", leftpad(Bold(message), size))
		}
	}

	// Listen to signals to exit.
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		timer.Stop()
		os.Exit(0)
	}()

	timer.Start(int64(duration.Seconds()))
	for isRepeat {
		timer.Start(int64(duration.Seconds()))
	}
}
