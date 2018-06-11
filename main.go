package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	timer := NewTimer()
	timer.OnTick = func(remaining int64) {
		fmt.Printf("%d\r", remaining)
	}
	timer.OnFinish = func() {
		fmt.Println("Exiting...")
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanRunes)

		for {
			for scanner.Scan() {
				fmt.Printf("Pressed %s\n", scanner.Text())
			}
		}
	}()

	timer.Start(20)
}
