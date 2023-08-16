package clock

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)

	// hand length is 90
	// since svg origin is top left
	// our origin is 150,150
	p = Point{p.X*90 + 150, -p.Y*90 + 150}
	return p
}

func secondsInRadians(t time.Time) float64 {
	// 1s = 2pi/60 radians
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
