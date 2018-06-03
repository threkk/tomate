package main

import (
	"fmt"
)

func main() {
	timer := NewTimer()
	timer.OnTick = func(remaining int64) {
		fmt.Printf("%d\r", remaining)
	}
	timer.OnFinish = func() {
		fmt.Println("Exiting...")
	}
	timer.Start(3)
}
