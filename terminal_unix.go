// +build !windows

package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

// Extracted from https://stackoverflow.com/questions/16569433/get-terminal-size-in-go/16576712#16576712
const sttyRowCols = `(?m)\d+ (\d+)`

func getColumns() uint {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		return 80
	}

	re := regexp.MustCompile(sttyRowCols)
	cols := re.FindSubmatch(out)

	if cols == nil {
		return 80
	}

	c, err := strconv.ParseUint(string(cols[1]), 10, 32)
	if err != nil {
		log.Println(err)
		return 80
	}

	return uint(c)
}
