package main

import (
	"os"
	"time"

	clock "github.com/juaoose/lgwt/math"
)

func main() {
	t := time.Now()
	clock.SVGWriter(os.Stdout, t)
}
