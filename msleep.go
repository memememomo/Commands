package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var sec float64
	_, err := fmt.Fscanf(os.Stdin, "%f\n", &sec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(0)
	}
	milli := int(sec * 1000)
	time.Sleep(time.Duration(milli) * time.Millisecond)
}
