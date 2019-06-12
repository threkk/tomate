package main

import (
	"flag"
	"fmt"
	"github.com/threkk/tomate/internal/pkg/ui"
	"github.com/threkk/tomate/pkg/terminal"
	"github.com/threkk/tomate/pkg/timer"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var version string

var stop chan os.Signal
var pause chan os.Signal
var size int
var duration time.Duration
var message string
var isQuiet bool
var isRepeat bool
var isVersion bool

func usage() {
	fmt.Printf("tomate - Simple pomodoro 🍅 (v%s)\n", version)
	fmt.Printf("\n")
	fmt.Printf("Stop it with Ctrl+C. Pause/resume with Ctrl+Z\n")
	fmt.Printf("\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("  tomate [-repeat] [-quiet] [-duration <duration>] [-message <message>]\n")
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
	size = int(terminal.GetColumns())

	if version == "" {
		version = "dev"
	}

	flag.DurationVar(&duration, "duration", 25*time.Minute, "Duration of the timer. 25 minutes by default.")
	flag.StringVar(&message, "message", "", "Message to display once the timer finishes.")
	flag.BoolVar(&isQuiet, "quiet", false, "Quite mode (hides the UI).")
	flag.BoolVar(&isRepeat, "repeat", false, "Restart the pomodoro once it finishes.")
	flag.BoolVar(&isVersion, "version", false, "Show the version number")

	flag.Usage = usage

	stop = make(chan os.Signal)
	pause = make(chan os.Signal)
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
	bar := &ui.HomeBrewBar{Size: size}

	t := timer.NewTimer()
	if !isQuiet {
		t.OnTick = func(current int64, total int64) {
			frame := bar.Frame(current, total)
			fmt.Printf("\r%s", frame)
		}
	}

	if message != "" {
		t.OnFinish = func() {
			fmt.Printf("\n\n%s\n\n", leftpad(ui.Bold(message), size))
		}
	}

	// Listen to signals
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(pause, syscall.SIGTSTP, syscall.SIGCONT)
	go func() {
		for {
			select {
			case <-stop:
				t.Stop()
				os.Exit(0)
			case <-pause:
				fmt.Printf("\r")
				t.Toggle()
				break
			}
		}
	}()

	t.Start(int64(duration.Seconds()))
	for isRepeat {
		t.Start(int64(duration.Seconds()))
	}
}
