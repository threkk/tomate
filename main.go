package main

import (
	"fmt"
)

func main() {
	timer := NewTimer()
	// echo $COLUMNS
	rb := &RedBlackBar{Size: 40}
	timer.OnTick = func(current int64, total int64) {
		bar := rb.Frame(current, total)
		fmt.Printf("%s\r", bar)
	}

	timer.Start(10)
}
