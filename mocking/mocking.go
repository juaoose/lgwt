package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalMessage = "Go!"

func Countdown(writer io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(writer, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(writer, finalMessage)
}

func main() {
	Countdown(os.Stdout)
}
